package common

import (
	"fmt"
	"os"
)

// Two output modes, one rule (docs/cli-guide.md, Output modes): human-readable
// when stdout is a terminal, JSON when it is not, and -json / -human force it.
//
// The mode is decided by stdout ALONE — never stderr, never an environment
// variable. An environment variable is a mode a caller did not type and cannot
// see in the command line it is reading.
//
// This lives here rather than in forge/primitives because primitives does not
// hold it yet. When it does, this file is a deletion and an import.

// OutputMode is how a result is rendered.
type OutputMode int

const (
	// OutputHuman is the mode for a person at a terminal.
	OutputHuman OutputMode = iota
	// OutputJSON is the mode for anything reading the result.
	OutputJSON
)

// OutputFlags is what -json and -human were given as, before the default is
// applied. Both false is "not asked", which is the ordinary case.
type OutputFlags struct {
	JSON  bool
	Human bool
}

// TakeOutputFlags strips -json and -human from args and returns the rest.
//
// Stripping rather than parsing with the flag package: these tools take their
// arguments in any position, and a FlagSet stops at the first non-flag, so
// `run --list -json` and `run -json --list` would not mean the same thing.
// NormalizeArgs has already made --json and -json one name.
func TakeOutputFlags(args []string) ([]string, OutputFlags) {
	var of OutputFlags
	rest := make([]string, 0, len(args))
	for _, a := range args {
		switch a {
		case "-json":
			of.JSON = true
		case "-human":
			of.Human = true
		default:
			rest = append(rest, a)
		}
	}
	return rest, of
}

// Mode resolves the flags against the default. Passing both is a usage error
// (docs/cli-guide.md, Fail closed), named before anything is done.
//
// The default asks stdout whether it is a character device: a terminal is, and
// a pipe or a redirect is not.
func (of OutputFlags) Mode() (OutputMode, error) {
	if of.JSON && of.Human {
		return OutputHuman, fmt.Errorf("-json and -human are mutually exclusive: pass one, or neither to let stdout decide")
	}
	switch {
	case of.JSON:
		return OutputJSON, nil
	case of.Human:
		return OutputHuman, nil
	}
	fi, err := os.Stdout.Stat()
	if err != nil {
		// A stdout that cannot be described is not a terminal anyone is reading,
		// and JSON is the answer that survives being piped somewhere unexamined.
		return OutputJSON, nil
	}
	if fi.Mode()&os.ModeCharDevice != 0 {
		return OutputHuman, nil
	}
	return OutputJSON, nil
}
