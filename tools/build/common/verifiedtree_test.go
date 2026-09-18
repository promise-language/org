package common

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/promise-language/forge/primitives"
)

// The contract, tested from the recording end: the id verify records must
// equal the id the guard computes over the real index after `git add -A`,
// because that agreement is the whole of what the two ends share. The guard
// is a workspace tool and is not in this repository, so the comparison is
// spelled out here as `git write-tree` — exactly what the guard runs.

func git(t *testing.T, dir string, args ...string) string {
	t.Helper()
	cmd := exec.Command("git", args...)
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("git %s: %v\n%s", strings.Join(args, " "), err, out)
	}
	return strings.TrimSpace(string(out))
}

func writeFile(t *testing.T, path, body string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("mkdir %s: %v", filepath.Dir(path), err)
	}
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}

// fixtureIgnores is what every project's committed .gitignore carries for this
// contract: .workspace/ holds the record (tool-contract §3), and .home/ holds
// the temporary index that computes it. Stated once, because a test that
// rewrites .gitignore and drops an entry is testing a checkout no project has.
const fixtureIgnores = ".workspace/\n.home/\n"

// verifyRepoForTest is a fresh checkout with an identity and the .gitignore
// every project is required to carry.
func verifyRepoForTest(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	git(t, dir, "init", "-q", "-b", "main")
	git(t, dir, "config", "user.email", "t@example.com")
	git(t, dir, "config", "user.name", "T")
	writeFile(t, filepath.Join(dir, ".gitignore"), fixtureIgnores)
	return dir
}

func recordedTree(t *testing.T, dir string) string {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(dir, ".workspace", "verified-tree"))
	if err != nil {
		t.Fatalf("read the record: %v", err)
	}
	return strings.TrimSpace(string(data))
}

func TestRecordMatchesRealStage(t *testing.T) {
	dir := verifyRepoForTest(t)
	writeFile(t, filepath.Join(dir, "a.txt"), "a\n")
	writeFile(t, filepath.Join(dir, "sub", "b.txt"), "b\n")
	git(t, dir, "add", "-A")
	git(t, dir, "commit", "-q", "-m", "base")
	writeFile(t, filepath.Join(dir, "a.txt"), "a2\n")

	if err := recordVerifiedTree(dir); err != nil {
		t.Fatalf("recordVerifiedTree: %v", err)
	}
	git(t, dir, "add", "-A")
	staged := git(t, dir, "write-tree")
	if got := recordedTree(t, dir); got != staged {
		t.Errorf("recorded %s, but git add -A stages %s — the two ends disagree", got, staged)
	}
}

func TestRecordIncludesUntracked(t *testing.T) {
	// A step whose whole output is new files must produce a committable match,
	// which is why `git stash create` (tracked modifications only) is not it.
	dir := verifyRepoForTest(t)
	writeFile(t, filepath.Join(dir, "a.txt"), "a\n")
	git(t, dir, "add", "-A")
	git(t, dir, "commit", "-q", "-m", "base")
	writeFile(t, filepath.Join(dir, "new.txt"), "new\n")

	if err := recordVerifiedTree(dir); err != nil {
		t.Fatalf("recordVerifiedTree: %v", err)
	}
	names := git(t, dir, "ls-tree", "-r", "--name-only", recordedTree(t, dir))
	if !strings.Contains(names, "new.txt") {
		t.Errorf("untracked non-ignored file missing from recorded tree: %q", names)
	}
}

func TestRecordRespectsIgnoreRules(t *testing.T) {
	// An ignored file stays out; a tracked-but-ignored file stays in — the
	// reason the temp index is seeded rather than left empty.
	dir := verifyRepoForTest(t)
	writeFile(t, filepath.Join(dir, ".gitignore"), fixtureIgnores+"ignored.txt\npinned.txt\n")
	writeFile(t, filepath.Join(dir, "pinned.txt"), "pinned\n")
	git(t, dir, "add", "-A")
	git(t, dir, "add", "-f", "pinned.txt")
	git(t, dir, "commit", "-q", "-m", "base")
	writeFile(t, filepath.Join(dir, "ignored.txt"), "ignored\n")

	if err := recordVerifiedTree(dir); err != nil {
		t.Fatalf("recordVerifiedTree: %v", err)
	}
	names := git(t, dir, "ls-tree", "-r", "--name-only", recordedTree(t, dir))
	if strings.Contains(names, "ignored.txt") {
		t.Errorf("ignored file should not be in the recorded tree: %q", names)
	}
	if !strings.Contains(names, "pinned.txt") {
		t.Errorf("tracked-but-ignored file should be in the recorded tree: %q", names)
	}
}

func TestRecordTrackedSetFollowsIndexNotHEAD(t *testing.T) {
	// The tracked set `git add -A` starts from is the real index's, not
	// HEAD's, and the two differ exactly where ignore rules bite: an ignored
	// file force-added but not yet committed must be in the blessed tree, and
	// one just `git rm --cached`ed must be out. Seeded any other way, the
	// record is a tree no `git add -A` can stage — a permanent guard refusal
	// whose named recovery, re-running verify, reproduces it.
	dir := verifyRepoForTest(t)
	writeFile(t, filepath.Join(dir, ".gitignore"), fixtureIgnores+"added.txt\ndropped.txt\n")
	writeFile(t, filepath.Join(dir, "dropped.txt"), "dropped\n")
	git(t, dir, "add", "-A")
	git(t, dir, "add", "-f", "dropped.txt")
	git(t, dir, "commit", "-q", "-m", "base")
	writeFile(t, filepath.Join(dir, "added.txt"), "added\n")
	git(t, dir, "add", "-f", "added.txt")
	git(t, dir, "rm", "-q", "--cached", "dropped.txt")

	if err := recordVerifiedTree(dir); err != nil {
		t.Fatalf("recordVerifiedTree: %v", err)
	}
	git(t, dir, "add", "-A")
	staged := git(t, dir, "write-tree")
	if got := recordedTree(t, dir); got != staged {
		t.Errorf("recorded %s, but git add -A stages %s — the two ends disagree", got, staged)
	}
	names := git(t, dir, "ls-tree", "-r", "--name-only", recordedTree(t, dir))
	if !strings.Contains(names, "added.txt") {
		t.Errorf("force-added file missing from recorded tree: %q", names)
	}
	if strings.Contains(names, "dropped.txt") {
		t.Errorf("rm --cached'ed file should not be in the recorded tree: %q", names)
	}
}

func TestRecordLeavesIndexAlone(t *testing.T) {
	dir := verifyRepoForTest(t)
	writeFile(t, filepath.Join(dir, "a.txt"), "a\n")
	git(t, dir, "add", "-A")
	git(t, dir, "commit", "-q", "-m", "base")
	writeFile(t, filepath.Join(dir, "staged.txt"), "staged\n")
	git(t, dir, "add", "staged.txt")
	writeFile(t, filepath.Join(dir, "unstaged.txt"), "unstaged\n")

	before := git(t, dir, "diff", "--cached", "--name-only")
	if err := recordVerifiedTree(dir); err != nil {
		t.Fatalf("recordVerifiedTree: %v", err)
	}
	after := git(t, dir, "diff", "--cached", "--name-only")
	if before != after {
		t.Errorf("recording disturbed the real index: before %q, after %q", before, after)
	}
}

// NOTHING IS WRITTEN OUTSIDE THE CHECKOUT. identity.md, Where records are
// kept: "A project or workspace tool running in an arena writes nothing
// outside its checkout." The system temporary directory is the one path every
// arena on the machine shares, so it is the one this recording used to write
// to — every bin/verify in every arena, with neither able to tell.
//
// The check is that recording does not need that directory to exist, not that
// it leaves it tidy: the temp index was always removed at the end of the call,
// so a run that had written there is indistinguishable afterwards from one
// that had not. Pointed at a path that is not there, a recording that reaches
// for it fails and one that never does passes. A missing directory rather than
// an unwritable one, because chmod is a permission quirk of the machine the
// tests run on and is none at all for root. os.TempDir reads TMPDIR everywhere
// but Windows, where it reads TMP then TEMP, so all three are moved.
func TestRecordDoesNotReachForTheSystemTempDirectory(t *testing.T) {
	// The checkout is taken before the environment moves: t.TempDir reads the
	// same variables, and it must not be sent to the missing path too.
	dir := verifyRepoForTest(t)
	missing := filepath.Join(t.TempDir(), "no-such-dir")
	t.Setenv("TMPDIR", missing)
	t.Setenv("TMP", missing)
	t.Setenv("TEMP", missing)
	writeFile(t, filepath.Join(dir, "a.txt"), "a\n")

	if err := recordVerifiedTree(dir); err != nil {
		t.Fatalf("recording reached outside the checkout for scratch space: %v", err)
	}
	if got := recordedTree(t, dir); len(got) < 40 {
		t.Errorf("record should hold one tree id, got %q", got)
	}
	if primitives.Exists(missing) {
		t.Errorf("recording created %s — it wrote outside the checkout", missing)
	}
}

func TestRecordScratchLivesInsideTheCheckout(t *testing.T) {
	// The other half of the same rule: the temp index is in the checkout's own
	// scratch directory, and gone again once the record is written.
	dir := verifyRepoForTest(t)
	writeFile(t, filepath.Join(dir, "a.txt"), "a\n")

	if err := recordVerifiedTree(dir); err != nil {
		t.Fatalf("recordVerifiedTree: %v", err)
	}
	scratch := filepath.Join(dir, filepath.FromSlash(scratchRel))
	entries, err := os.ReadDir(scratch)
	if err != nil {
		t.Fatalf("the temp index belongs under %s in the checkout: %v", scratchRel, err)
	}
	for _, e := range entries {
		if strings.HasPrefix(e.Name(), "verified-tree-") {
			t.Errorf("temp index left behind: %s", filepath.Join(scratch, e.Name()))
		}
	}
}

// EVERY ARENA BUT A FRESH CLONE ALREADY HAS THE SCRATCH DIRECTORY, WITH OTHER
// TOOLS' SCRATCH IN IT. Every fixture here starts without .home/, so they only
// ever exercise creating it; the checkout a person or an agent runs bin/verify
// in has been through a run of something and .home/tmp/ is full. Three things
// follow, and none of them held before the temp index moved inside the
// checkout: creating the directory tolerates one already there, what is
// already in it stays out of the blessed tree now that `git add -A` walks
// past it, and recording removes its own temp index rather than the directory
// — clearing the scratch of every other tool in the arena would be a worse
// failure than the one this change fixed.
func TestRecordWhenTheScratchDirectoryAlreadyExists(t *testing.T) {
	dir := verifyRepoForTest(t)
	writeFile(t, filepath.Join(dir, "a.txt"), "a\n")
	elsewhere := filepath.Join(dir, filepath.FromSlash(scratchRel), "someone-elses.txt")
	writeFile(t, elsewhere, "another tool's scratch\n")

	if err := recordVerifiedTree(dir); err != nil {
		t.Fatalf("recording into a scratch directory that already exists: %v", err)
	}
	git(t, dir, "add", "-A")
	staged := git(t, dir, "write-tree")
	if got := recordedTree(t, dir); got != staged {
		t.Errorf("recorded %s, but git add -A stages %s — the two ends disagree", got, staged)
	}
	if names := git(t, dir, "ls-tree", "-r", "--name-only", recordedTree(t, dir)); strings.Contains(names, "someone-elses.txt") {
		t.Errorf("another tool's scratch is in the blessed tree: %q", names)
	}
	if !primitives.Exists(elsewhere) {
		t.Error("recording cleared the whole scratch directory, not just its own temp index")
	}
}

// A SCRATCH DIRECTORY THE CHECKOUT DOES NOT IGNORE IS REFUSED. `git add -A`
// walks the worktree, so an unignored temp index stages itself: the blessed
// tree gains an entry for a file that is deleted moments later, and no later
// `git add -A` can ever stage that tree again. That is a permanent, silent
// refusal of every commit whose named recovery — run bin/verify — reproduces
// it, so it is named at the moment it can still be fixed.
func TestRecordRefusesWhenScratchIsNotIgnored(t *testing.T) {
	dir := verifyRepoForTest(t)
	writeFile(t, filepath.Join(dir, ".gitignore"), ".workspace/\n")
	writeFile(t, filepath.Join(dir, "a.txt"), "a\n")

	err := recordVerifiedTree(dir)
	if err == nil {
		t.Fatal("recording should refuse a checkout that does not ignore the scratch directory")
	}
	if !strings.Contains(err.Error(), scratchRel) {
		t.Errorf("err = %v, want it to name %s", err, scratchRel)
	}
	if primitives.Exists(filepath.Join(dir, filepath.FromSlash(primitives.VerifiedTreeRecord))) {
		t.Error("a refused recording must bless nothing")
	}
	// And the refusal came before anything was created, so there is nothing to
	// clean up and nothing unignored left in the worktree.
	if primitives.Exists(filepath.Join(dir, ".home")) {
		t.Error("the refusal created the scratch directory it was refusing to use")
	}
}

// A DIRECTORY-ONLY RULE IGNORES THE SCRATCH DIRECTORY AND IS NOT A REFUSAL.
// `.home/tmp/` matches a directory and nothing else, and git decides what is a
// directory by looking: in a fresh clone nothing has created .home/tmp yet, so
// asked about that path alone git answers "not ignored" and the first
// bin/verify of the clone fails, naming a .gitignore change the checkout does
// not need. Asked about a file inside it — the temp index, which is what
// `git add -A` would actually stage — the leading components are directories
// by construction and the rule matches.
func TestRecordAcceptsADirectoryOnlyIgnoreRule(t *testing.T) {
	dir := verifyRepoForTest(t)
	writeFile(t, filepath.Join(dir, ".gitignore"), ".workspace/\n.home/tmp/\n")
	writeFile(t, filepath.Join(dir, "a.txt"), "a\n")
	if primitives.Exists(filepath.Join(dir, ".home")) {
		t.Fatal("the fixture already holds the scratch directory: this is the fresh-clone case")
	}

	if err := recordVerifiedTree(dir); err != nil {
		t.Fatalf("a checkout that does ignore %s was refused: %v", scratchRel, err)
	}
	// And the tree it blessed is still the one `git add -A` stages, so the
	// temp index the rule ignores stayed out of it.
	git(t, dir, "add", "-A")
	staged := git(t, dir, "write-tree")
	if got := recordedTree(t, dir); got != staged {
		t.Errorf("recorded %s, but git add -A stages %s — the two ends disagree", got, staged)
	}
}

func TestRecordFailsWhenScratchCannotBeCreated(t *testing.T) {
	dir := verifyRepoForTest(t)
	writeFile(t, filepath.Join(dir, "a.txt"), "a\n")
	// A regular file exactly where the scratch directory belongs: MkdirAll
	// refuses it, and it is neither "absent" nor a permission quirk of the
	// machine the tests run on.
	writeFile(t, filepath.Join(dir, filepath.FromSlash(scratchRel)), "x\n")

	err := recordVerifiedTree(dir)
	if err == nil {
		t.Fatal("recording passed although the temp index had nowhere to go")
	}
	// The full path, which is what the creation failure names and the
	// not-ignored refusal does not — so this cannot pass on the other branch.
	scratch := filepath.Join(dir, filepath.FromSlash(scratchRel))
	if !strings.Contains(err.Error(), scratch) {
		t.Errorf("err = %v, want it to name %s", err, scratch)
	}
	if primitives.Exists(filepath.Join(dir, filepath.FromSlash(primitives.VerifiedTreeRecord))) {
		t.Error("a tree nothing could be staged into must not be blessed")
	}
}

// A RECORDING THAT FAILS PART WAY LEAVES NO TEMP INDEX BEHIND. The removal is
// deferred, so every other test reaches it on the way out of a call that
// worked, and moving it to the end of the happy path would fail none of them.
// The failures are where it earns its keep, and they changed with this fix:
// a temp index leaked by a failed run used to accumulate in the system
// temporary directory, which the operating system eventually clears, and now
// accumulates inside the checkout, in a directory `git add -A` walks, where
// nothing clears it.
func TestRecordRemovesTheTempIndexWhenRecordingFails(t *testing.T) {
	dir := verifyRepoForTest(t)
	writeFile(t, filepath.Join(dir, "a.txt"), "a\n")
	// A regular file where the record's own directory belongs. That failure
	// lands after the temp index has been created and the tree computed, which
	// is the only stretch of the call where the deferred removal is observable.
	writeFile(t, filepath.Join(dir, ".workspace"), "x\n")

	err := recordVerifiedTree(dir)
	if err == nil {
		t.Fatal("recording passed although the record had nowhere to go")
	}
	// Pinned to that failure and not an earlier one, or the temp index this
	// asserts about would never have been created.
	if !strings.Contains(err.Error(), ".workspace") {
		t.Fatalf("err = %v, want the failure that writes the record", err)
	}
	scratch := filepath.Join(dir, filepath.FromSlash(scratchRel))
	entries, readErr := os.ReadDir(scratch)
	if readErr != nil {
		t.Fatalf("read %s: %v", scratch, readErr)
	}
	for _, e := range entries {
		if strings.HasPrefix(e.Name(), "verified-tree-") {
			t.Errorf("a failed recording left its temp index in the checkout: %s", filepath.Join(scratch, e.Name()))
		}
	}
}

// THE FORMAT IS HALF THE CONTRACT. The guard is a workspace tool that cannot be
// imported here, so nothing but agreement on the bytes connects the two ends:
// one tree id, newline terminated, and nothing else in the file. Every other
// test in this file reads the record through a trim, so a record written
// without its newline — or with a second line of commentary — passes all of
// them and is refused by a reader this repository cannot run.
func TestRecordIsOneTreeIdNewlineTerminated(t *testing.T) {
	dir := verifyRepoForTest(t)
	writeFile(t, filepath.Join(dir, "a.txt"), "a\n")

	if err := recordVerifiedTree(dir); err != nil {
		t.Fatalf("recordVerifiedTree: %v", err)
	}
	raw, err := os.ReadFile(filepath.Join(dir, filepath.FromSlash(primitives.VerifiedTreeRecord)))
	if err != nil {
		t.Fatalf("read the record: %v", err)
	}
	body := string(raw)
	if !strings.HasSuffix(body, "\n") {
		t.Errorf("record = %q, want it newline terminated", body)
	}
	id := strings.TrimSuffix(body, "\n")
	if strings.ContainsAny(id, "\n ") || id == "" {
		t.Errorf("record = %q, want exactly one bare tree id and nothing else", body)
	}
	// And it must name a tree git can resolve, not merely look like an id.
	if got := git(t, dir, "cat-file", "-t", id); got != "tree" {
		t.Errorf("recorded id is a %q, want a tree", got)
	}
}

func TestClearVerifiedTree(t *testing.T) {
	dir := t.TempDir()
	record := filepath.Join(dir, ".workspace", "verified-tree")
	writeFile(t, record, "abc\n")
	if err := clearVerifiedTree(dir); err != nil {
		t.Fatalf("clearing an existing record: %v", err)
	}
	if primitives.Exists(record) {
		t.Error("record should be gone after clear")
	}
	if err := clearVerifiedTree(dir); err != nil {
		t.Fatalf("clearing an absent record should not be an error: %v", err)
	}
}

func TestRecordOutsideGitCheckout(t *testing.T) {
	dir := t.TempDir()
	if err := recordVerifiedTree(dir); err != nil {
		t.Fatalf("outside a checkout recording should be a no-op, not an error: %v", err)
	}
	if primitives.Exists(filepath.Join(dir, ".workspace", "verified-tree")) {
		t.Error("no record should be written outside a git checkout")
	}
	// A no-op that leaves a directory behind is not a no-op: the return comes
	// before the scratch directory is created, not after.
	if primitives.Exists(filepath.Join(dir, ".home")) {
		t.Error("a no-op run created the scratch directory")
	}
}

func TestVerifyPipelineEndsWithRecord(t *testing.T) {
	stub := verifyPipeline(t.TempDir())
	if len(stub) == 0 || stub[len(stub)-1].name != "record" {
		t.Errorf("stub pipeline should end with record: %v", stepNames(stub))
	}
	goDir := t.TempDir()
	writeFile(t, filepath.Join(goDir, "go.mod"), "module example.test\n")
	goSteps := verifyPipeline(goDir)
	if len(goSteps) == 0 || goSteps[len(goSteps)-1].name != "record" {
		t.Errorf("go pipeline should end with record: %v", stepNames(goSteps))
	}
}

func TestRecordStepNotReachedAfterFailure(t *testing.T) {
	// The record step rides the existing break-on-first-failure, so a red step
	// blesses nothing.
	reached := false
	steps := []step{
		{"boom", func(string) error { return os.ErrInvalid }},
		{"record", func(string) error { reached = true; return nil }},
	}
	if err := runVerifySteps(t.TempDir(), steps); err == nil {
		t.Fatal("a failing step should fail the run")
	}
	if reached {
		t.Error("record step must not run after an earlier failure")
	}
}

func TestRunVerifyStubClearsAndRecords(t *testing.T) {
	// End to end through RunVerify on the stub pipeline: a stale record is
	// cleared at the start and a fresh tree id is recorded at the end.
	dir := verifyRepoForTest(t)
	writeFile(t, filepath.Join(dir, "a.txt"), "a\n")
	writeFile(t, filepath.Join(dir, ".workspace", "verified-tree"), "stale-garbage\n")
	if err := RunVerify(dir, nil); err != nil {
		t.Fatalf("RunVerify: %v", err)
	}
	got := recordedTree(t, dir)
	if got == "stale-garbage" {
		t.Fatal("stale record survived the run")
	}
	if len(got) < 40 {
		t.Errorf("record should hold one tree id, got %q", got)
	}
}

func TestRunVerifyRedRunLeavesNothingBlessed(t *testing.T) {
	// The clear at the start of RunVerify is only observable on a red run —
	// a green run overwrites the record at the end anyway, so dropping the
	// clear call is invisible to TestRunVerifyStubClearsAndRecords. A stale
	// blessing surviving a failed verify is the one outcome the check must
	// never produce: an unparseable file cannot pass format, so the run goes
	// red before record — and the pre-seeded blessing has to be gone.
	dir := verifyRepoForTest(t)
	writeFile(t, filepath.Join(dir, "go.mod"), "module example.test\n")
	writeFile(t, filepath.Join(dir, "broken.go"), "package broken\nfunc {\n")
	writeFile(t, filepath.Join(dir, ".workspace", "verified-tree"), "stale-blessing\n")
	if err := RunVerify(dir, nil); err == nil {
		t.Fatal("verify over an unparseable Go file should fail")
	}
	if primitives.Exists(filepath.Join(dir, ".workspace", "verified-tree")) {
		t.Error("a red run must leave nothing blessed — the stale record survived")
	}
}

// A CLEAR THAT FAILS FAILS THE RUN. The clear exists so a run that dies part
// way leaves nothing blessed; if it can fail and be ignored, the case it was
// added for is exactly the case it does not cover — the record it could not
// remove survives the whole run, and a red verify hands the guard a stale
// blessing for a tree nobody checked. Nothing may run past it.
func TestRunVerifyFailsWhenTheStaleRecordCannotBeCleared(t *testing.T) {
	dir := verifyRepoForTest(t)
	writeFile(t, filepath.Join(dir, "a.txt"), "a\n")
	// A non-empty directory where the record belongs: os.Remove refuses it, the
	// one clear failure that is neither "absent" nor a permission quirk of the
	// machine the tests run on.
	writeFile(t, filepath.Join(dir, filepath.FromSlash(primitives.VerifiedTreeRecord), "occupied"), "x\n")

	err := RunVerify(dir, nil)
	if err == nil {
		t.Fatal("RunVerify passed although the stale record could not be cleared")
	}
	if !strings.Contains(err.Error(), primitives.VerifiedTreeRecord) {
		t.Errorf("err = %v, want it to name %s", err, primitives.VerifiedTreeRecord)
	}
	// And it stopped there rather than running the pipeline over it.
	if !primitives.Exists(filepath.Join(dir, filepath.FromSlash(primitives.VerifiedTreeRecord), "occupied")) {
		t.Error("the run went on and disturbed what it could not clear")
	}
}

func stepNames(steps []step) []string {
	var names []string
	for _, s := range steps {
		names = append(names, s.name)
	}
	return names
}
