package main

import (
	"encoding/json"
	"strings"
	"testing"

	"org/tools/build/common"
)

// The workspace asks `bin/run --list -json`, so the flag must be recognised
// alongside the output flags in either order. A bare `list` is not a second
// spelling of the flag: it is a gate name this project does not have.
func TestWantsList(t *testing.T) {
	for _, args := range [][]string{
		{"-list"},
		{"-list", "-json"},
		{"-json", "-list"},
		{"-list", "-human"},
	} {
		if !wantsList(args) {
			t.Errorf("wantsList(%q) = false, want true", args)
		}
	}
	for _, args := range [][]string{
		{},
		{"list"},
		{"fit"},
		{"fit", "-verdict"},
	} {
		if wantsList(args) {
			t.Errorf("wantsList(%q) = true, want false", args)
		}
	}
}

// Every problem with the invocation is named, and a clean one resolves a mode.
func TestListProblems(t *testing.T) {
	if mode, problems := listProblems([]string{"-list", "-json"}); len(problems) != 0 || mode != common.OutputJSON {
		t.Errorf("-list -json: mode %v, problems %v; want JSON and none", mode, problems)
	}
	if mode, problems := listProblems([]string{"-human", "-list"}); len(problems) != 0 || mode != common.OutputHuman {
		t.Errorf("-human -list: mode %v, problems %v; want human and none", mode, problems)
	}
	_, problems := listProblems([]string{"-list", "fit", "-json", "-human"})
	if len(problems) != 2 {
		t.Fatalf("problems = %v, want both the stray argument and the contradictory modes", problems)
	}
	if !strings.Contains(problems[0], "fit") {
		t.Errorf("problems[0] = %q, want it to name the stray argument", problems[0])
	}
}

// Human mode says which name is which; JSON mode is the object the workspace
// reads. Both render the same two lists.
func TestRenderListBothModes(t *testing.T) {
	answer := listAnswer{Commands: []string{"gate", "run"}, Gates: []string{"fit", "integration"}}

	var human strings.Builder
	if err := renderList(&human, answer, common.OutputHuman); err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"command  gate", "command  run", "gate     fit", "gate     integration"} {
		if !strings.Contains(human.String(), want) {
			t.Errorf("human output %q is missing %q", human.String(), want)
		}
	}

	var jsonOut strings.Builder
	if err := renderList(&jsonOut, answer, common.OutputJSON); err != nil {
		t.Fatal(err)
	}
	var got listAnswer
	if err := json.Unmarshal([]byte(jsonOut.String()), &got); err != nil {
		t.Fatalf("JSON mode did not emit one decodable object: %v", err)
	}
	if strings.Join(got.Commands, ",") != "gate,run" || strings.Join(got.Gates, ",") != "fit,integration" {
		t.Errorf("decoded %+v, want %+v", got, answer)
	}
}

// The discovery answer must name the two gates a flow cannot run without.
func TestProjectAnswersTheRequiredGates(t *testing.T) {
	for _, required := range []string{"fit", "integration"} {
		if !common.KnownGate(required) {
			t.Errorf("this project does not answer %q; gates: %v", required, common.GateNames())
		}
	}
}
