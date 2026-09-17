package common

// What this project builds, answered by looking rather than by reading a list
// someone maintained.
//
// The tool set is one tool per directory under tools/build/cmd, except make
// itself, which runs from source and is never compiled into bin/. The listing
// IS the registry — there is no list anywhere to keep in step with it, so adding
// a tool is adding a directory and retiring one is deleting it (retiring `guard`
// and `precommit` was exactly that).
//
// The meta-builder and `run --list` both read it through here: two readers
// each walking the directory would be two answers to one question, free to
// disagree about what bin/ holds.

import (
	"os"
	"path/filepath"
	"sort"
)

// CommandNames returns every command this project builds into bin/, sorted.
func CommandNames(repoRoot string) ([]string, error) {
	return CommandsIn(filepath.Join(repoRoot, "tools", "build", "cmd"))
}

// CommandsIn lists the tools under one cmd directory, sorted.
//
// A file under cmd/ is not a tool: `go build ./cmd/<name>` wants a package.
// An unreadable directory is an error, never an empty set: no tools with no
// error would take make's up-to-date short circuit and leave bin/ however it
// was found, and tell `run --list` callers this project builds nothing.
func CommandsIn(cmdDir string) ([]string, error) {
	entries, err := os.ReadDir(cmdDir)
	if err != nil {
		return nil, err
	}
	names := make([]string, 0, len(entries))
	for _, e := range entries {
		if e.IsDir() && e.Name() != "make" {
			names = append(names, e.Name())
		}
	}
	sort.Strings(names)
	return names, nil
}
