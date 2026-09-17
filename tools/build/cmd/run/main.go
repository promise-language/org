// Command run asks one gate for a measurement and reaches a verdict on it.
//
// This is the by-hand path, and it takes the same route a runner takes rather
// than a parallel one: it executes bin/gate as a process and reads what came
// back. Running a single gate is not a lesser case — it is faster than
// everything that blocks a change from landing, and it is what someone
// iterating on one failure actually wants.
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/promise-language/forge/primitives"
	"org/tools/build/common"
)

// Injected by the meta-builder via -ldflags at build time; empty otherwise.
var (
	repoRoot   = ""
	sourceHash = ""
)

func usage() string {
	var sb strings.Builder
	sb.WriteString("run — measure one gate and judge what it measured.\n\n")
	sb.WriteString("Usage:\n  run <gate> [-help]\n  run <gate> --verdict < envelope\n  run --list [-json | -human]\n\n")
	sb.WriteString("Runs bin/gate <gate> --envelope, then prints each measurement beside the\n")
	sb.WriteString("term it was judged on. Exit 0 means every capped measurement is within its\n")
	sb.WriteString("cap; non-zero means one is not, or that nothing could be measured.\n\n")
	sb.WriteString("With --verdict it judges an envelope it is GIVEN, on stdin, and runs no\n")
	sb.WriteString("gate: it prints one JSON verdict on stdout and nothing else. That is the\n")
	sb.WriteString("mode the SDK asks — the SDK spawns the gate, because a judge that ran its\n")
	sb.WriteString("own measurement would be the runner, and the runner comes from outside the\n")
	sb.WriteString("tree.\n\n")
	sb.WriteString("Gates:\n")
	for _, n := range common.GateNames() {
		fmt.Fprintf(&sb, "  %-12s %s\n", n, common.GateSummary(n))
	}
	capped := common.CappedMetrics(repoRoot)
	if len(capped) > 0 {
		fmt.Fprintf(&sb, "\nJudged against a cap: %s\n", strings.Join(capped, ", "))
	} else {
		fmt.Fprintf(&sb, "\nThresholds defined in %s\n", common.ManifestFile)
	}
	sb.WriteString("Anything else is reported and not judged.\n\n")
	sb.WriteString("With --list it names what this project builds and what it answers:\n")
	sb.WriteString("`commands` are the binaries ./make puts in bin/, `gates` are the names\n")
	sb.WriteString("bin/gate --list declares. That is the discovery query — it is how anything\n")
	sb.WriteString("outside the tree learns both without holding a copy that can go stale.\n")
	sb.WriteString("Output is human-readable at a terminal and JSON when stdout is not one;\n")
	sb.WriteString("-json and -human force the mode, and passing both is a usage error.\n")
	return sb.String()
}

// listAnswer is what --list prints: the two kinds of name a caller outside this
// tree can ask this project for — a command it can execute, and a gate it can
// measure. They stay separate lists, because a check about binaries compares
// commands and a gate is not one; one object carries both because a caller that
// must not confuse the two needs the whole vocabulary at once.
type listAnswer struct {
	Commands []string `json:"commands"`
	Gates    []string `json:"gates"`
}

// wantsList reports whether the argv names the discovery query, so main can
// route to it before ParseRunArgs refuses `-list` as an unknown flag. The flag
// is the only spelling: a bare `list` is a gate name, and this project has no
// gate of that name.
func wantsList(args []string) bool {
	for _, a := range args {
		if a == "-list" {
			return true
		}
	}
	return false
}

// renderList writes the answer in the mode stdout asked for (docs/cli-guide.md,
// Output modes). The human form tags each name with its kind, so a reader can
// tell what this project builds from what it answers.
func renderList(w io.Writer, answer listAnswer, mode common.OutputMode) error {
	if mode == common.OutputJSON {
		enc := json.NewEncoder(w)
		enc.SetIndent("", "  ")
		return enc.Encode(answer)
	}
	for _, c := range answer.Commands {
		if _, err := fmt.Fprintf(w, "command  %s\n", c); err != nil {
			return err
		}
	}
	for _, g := range answer.Gates {
		if _, err := fmt.Fprintf(w, "gate     %s\n", g); err != nil {
			return err
		}
	}
	return nil
}

// listProblems returns every usage problem with a --list invocation, and the
// output mode when there are none. Fail closed: all of them are reported
// before anything is done.
func listProblems(args []string) (common.OutputMode, []string) {
	rest, outFlags := common.TakeOutputFlags(args)
	var bad []string
	for _, a := range rest {
		if a != "-list" {
			bad = append(bad, a)
		}
	}
	var problems []string
	if len(bad) > 0 {
		problems = append(problems, fmt.Sprintf("unexpected argument(s) alongside --list: %s", strings.Join(bad, ", ")))
	}
	mode, err := outFlags.Mode()
	if err != nil {
		problems = append(problems, err.Error())
	}
	return mode, problems
}

func main() {
	args := primitives.NormalizeArgs(os.Args[1:])
	if primitives.HasHelpFlag(args) {
		fmt.Print(usage())
		os.Exit(0)
	}
	primitives.CheckStale(repoRoot, sourceHash)

	// The discovery query. It comes after CheckStale for the reason the verdict
	// mode does: a binary whose logic has moved since it was compiled must not
	// answer with names it may no longer implement. It comes before
	// ParseRunArgs, which refuses an unknown flag, and --list is not a gate.
	if wantsList(args) {
		mode, problems := listProblems(args)
		if len(problems) > 0 {
			for _, p := range problems {
				fmt.Fprintf(os.Stderr, "run: %s\n", p)
			}
			fmt.Fprintf(os.Stderr, "run: run `%s -help` for usage\n", os.Args[0])
			os.Exit(2)
		}
		commands, err := common.CommandNames(repoRoot)
		if err != nil {
			fmt.Fprintf(os.Stderr, "run: cannot say what this project builds: %v\n", err)
			os.Exit(1)
		}
		answer := listAnswer{Commands: commands, Gates: common.GateNames()}
		if err := renderList(os.Stdout, answer, mode); err != nil {
			fmt.Fprintf(os.Stderr, "run: %v\n", err)
			os.Exit(1)
		}
		return
	}

	name, verdict, err := common.ParseRunArgs(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "run: %v; run `%s -help` for usage\n", err, os.Args[0])
		os.Exit(2)
	}

	// The judging mode. Nothing is spawned: the envelope arrives on stdin from
	// whoever ran the gate, and stdout carries one verdict and nothing else.
	// CheckStale has already run above, so stale tooling exits before it can
	// print a verdict rather than answering with terms nobody currently holds.
	if verdict {
		if err := common.JudgeStdin(repoRoot, name, os.Stdin, os.Stdout); err != nil {
			fmt.Fprintf(os.Stderr, "run: %v\n", err)
			os.Exit(1)
		}
		return
	}

	if err := common.RunOneGate(repoRoot, common.GateBinary(repoRoot), name); err != nil {
		fmt.Fprintf(os.Stderr, "run: %v\n", err)
		os.Exit(1)
	}
}
