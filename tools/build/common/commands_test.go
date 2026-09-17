package common

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCommandsIn_ListsDirectoriesExceptMake(t *testing.T) {
	cmdDir := t.TempDir()
	for _, name := range []string{"verify", "make", "gate"} {
		if err := os.Mkdir(filepath.Join(cmdDir, name), 0o755); err != nil {
			t.Fatal(err)
		}
	}
	// A file is not a tool: `go build ./cmd/notes.md` is not a build.
	if err := os.WriteFile(filepath.Join(cmdDir, "notes.md"), nil, 0o644); err != nil {
		t.Fatal(err)
	}

	got, err := CommandsIn(cmdDir)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Join(got, ",") != "gate,verify" {
		t.Errorf("CommandsIn() = %v, want [gate verify] — make runs from source and a file is not a package", got)
	}
}

// A cmd directory that cannot be read is not an empty tool set. Reporting no
// tools with no error would take make's up-to-date short circuit — nothing to
// build, every expected binary present, "Tools up to date" — and leave bin/
// however it was found, which for a fresh clone is empty.
func TestCommandsIn_UnreadableDirectoryIsAnError(t *testing.T) {
	if _, err := CommandsIn(filepath.Join(t.TempDir(), "cmd")); err == nil {
		t.Error("CommandsIn() succeeded on a missing cmd directory; make would report every tool up to date having built none")
	}
}

// CommandNames is CommandsIn over the repository's own cmd directory.
func TestCommandNames_ReadsTheRepositoryCmdDirectory(t *testing.T) {
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, "tools", "build", "cmd", "run"), 0o755); err != nil {
		t.Fatal(err)
	}
	got, err := CommandNames(root)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Join(got, ",") != "run" {
		t.Errorf("CommandNames() = %v, want [run]", got)
	}
}
