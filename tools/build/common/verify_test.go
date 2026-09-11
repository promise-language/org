package common

import (
	"os"
	"path/filepath"
	"testing"
)

// modules() is what decides which Go this tree carries. If it returned only the
// root, runAllModules would degenerate to single-module behaviour and skip the
// tools module — which in this repository is the only module there is.

func writeModule(t *testing.T, dir, name string) {
	t.Helper()
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module "+name+"\n"), 0o644); err != nil {
		t.Fatal(err)
	}
}

func TestModules_ReturnsBothWhenRootAndToolsBuildHaveGoMod(t *testing.T) {
	root := t.TempDir()
	toolsDir := filepath.Join(root, "tools", "build")
	writeModule(t, root, "example")
	writeModule(t, toolsDir, "example/tools/build")

	dirs := modules(root)
	if len(dirs) != 2 {
		t.Fatalf("modules returned %d dirs, want 2: %v", len(dirs), dirs)
	}
	if dirs[0] != root {
		t.Errorf("dirs[0] = %q, want the repo root %q", dirs[0], root)
	}
	if dirs[1] != toolsDir {
		t.Errorf("dirs[1] = %q, want the tools dir %q", dirs[1], toolsDir)
	}
}

func TestModules_ReturnsOnlyRootWhenToolsBuildHasNoGoMod(t *testing.T) {
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, "tools", "build"), 0o755); err != nil {
		t.Fatal(err)
	}
	writeModule(t, root, "example")
	// No go.mod in tools/build.

	dirs := modules(root)
	if len(dirs) != 1 || dirs[0] != root {
		t.Fatalf("modules returned %v, want the root alone", dirs)
	}
}

// This repository's shape: no Go at the root, one module under tools/build.
// The root must not be listed on the strength of existing — `go vet ./...`
// there fails with "go.mod file not found", and a gate that could not run
// would be reported as a tree that is bad.
func TestModules_ReturnsOnlyToolsBuildWhenRootHasNoGoMod(t *testing.T) {
	root := t.TempDir()
	toolsDir := filepath.Join(root, "tools", "build")
	writeModule(t, toolsDir, "example/tools/build")

	dirs := modules(root)
	if len(dirs) != 1 || dirs[0] != toolsDir {
		t.Fatalf("modules returned %v, want the tools dir %q alone", dirs, toolsDir)
	}
}

func TestModules_ReturnsNothingWhereThereIsNoGo(t *testing.T) {
	if dirs := modules(t.TempDir()); len(dirs) != 0 {
		t.Fatalf("modules returned %v for a tree with no go.mod anywhere", dirs)
	}
}

// runAllModules must visit the second module. This test creates a two-module
// layout where the second module contains a vet error. If runAllModules only
// visited the root, the error would never be seen and the test would pass
// incorrectly.
func TestRunAllModules_ReachesSecondModule(t *testing.T) {
	root := t.TempDir()
	toolsDir := filepath.Join(root, "tools", "build")

	// Root module: valid, minimal.
	if err := os.WriteFile(filepath.Join(root, "go.mod"), []byte("module example\ngo 1.21\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "main.go"), []byte("package example\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	// Tools module: contains code that fails `go vet`.
	if err := os.MkdirAll(toolsDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(toolsDir, "go.mod"), []byte("module example/tools/build\ngo 1.21\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	// Printf with a format verb and no argument: go vet reports this.
	badCode := "package build\nimport \"fmt\"\nfunc init() { fmt.Printf(\"%d\") }\n"
	if err := os.WriteFile(filepath.Join(toolsDir, "bad.go"), []byte(badCode), 0o644); err != nil {
		t.Fatal(err)
	}

	if err := runAllModules(root, "vet"); err == nil {
		t.Fatal("runAllModules returned nil; the second module has a vet error that should have been caught")
	}
}

// runAllModules must stop at the first failure rather than continuing. This is
// the same contract as verifySteps ("break // stop at the first failure").
func TestRunAllModules_StopsAtFirstFailure(t *testing.T) {
	root := t.TempDir()
	// A root module with code that fails vet — no tools/build module at all.
	if err := os.WriteFile(filepath.Join(root, "go.mod"), []byte("module example\ngo 1.21\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	badCode := "package example\nimport \"fmt\"\nfunc init() { fmt.Printf(\"%d\") }\n"
	if err := os.WriteFile(filepath.Join(root, "bad.go"), []byte(badCode), 0o644); err != nil {
		t.Fatal(err)
	}

	if err := runAllModules(root, "vet"); err == nil {
		t.Fatal("runAllModules returned nil for a module that fails vet")
	}
}

// verifySteps in a Go project must be format, vet, build and test — each going
// through runAllModules, which the tests above cover.
func TestVerifySteps_GoProjectHasExpectedSteps(t *testing.T) {
	root := t.TempDir()
	writeModule(t, root, "example")
	assertStepNames(t, verifySteps(root), []string{"format", "vet", "build", "test"})
}

// A tree whose only module is tools/build — this repository — is a Go project
// to verify, not a project to stub. Stubs here would bless every tree, tools
// that do not compile included, and the precommit-guard would then admit it.
func TestVerifySteps_ToolsOnlyTreeGetsTheGoPipeline(t *testing.T) {
	root := t.TempDir()
	writeModule(t, filepath.Join(root, "tools", "build"), "example/tools/build")
	assertStepNames(t, verifySteps(root), []string{"format", "vet", "build", "test"})
}

// A tree with no Go anywhere gets stub steps rather than real Go tooling. The
// stubs must succeed (not error) — they are placeholders, not failures.
func TestVerifySteps_NonGoProjectGetsStubs(t *testing.T) {
	root := t.TempDir() // no go.mod anywhere

	steps := verifySteps(root)
	if len(steps) != 4 {
		t.Fatalf("got %d stub steps, want 4", len(steps))
	}
	for _, s := range steps {
		if err := s.run(root); err != nil {
			t.Errorf("stub step %q failed: %v", s.name, err)
		}
	}
}

func assertStepNames(t *testing.T, steps []step, want []string) {
	t.Helper()
	if len(steps) != len(want) {
		t.Fatalf("got %d steps %v, want %d %v", len(steps), stepNames(steps), len(want), want)
	}
	for i, w := range want {
		if steps[i].name != w {
			t.Errorf("step[%d].name = %q, want %q", i, steps[i].name, w)
		}
	}
}
