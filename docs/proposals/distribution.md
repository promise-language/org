# Distribution

> **Proposal.** Not normative: an end state under discussion, binding nothing until ratified.

> **Tag:** `distribution` — remaining work to complete this document: the query named in
> `docs/index.md`.

What this repository places in every managed project's tree, and what keeps those copies honest.
The documents are not among them: a project stands on the corpus by
[reference](references.md), declaring the release it holds to, and holds no copy. What is
distributed is the handful of files another party reads at a fixed path in each repository —
GitHub reading a licence at the root, an action reading a workflow — where a reference cannot
stand in for the file. The process around this machinery — intake, the amendment pass, the
release — is [release-cycle](release-cycle.md)'s; the tools that run it are [org](org.md)'s.

## The vendored set

| Member | Lands at | Form |
|---|---|---|
| `LICENSE`, `LICENSE-APACHE`, `LICENSE-MIT` | the repository root | Byte-identical — the pointer is worded without a repo name, and the MIT holder line is the one org-wide holder |
| The CLA workflow | `.github/workflows/cla.yml` | Byte-identical, including its skip-on-private guard |
| The shared `CONTRIBUTING.md` sections — CLA, licensing of contributions, commit identity | inside the project's own `CONTRIBUTING.md` | The one member that is not a whole file; see open questions |

One rule covers them all: **a fleet whose copies of a legal text or a policy check drift per
repository has several policies where it means to have one.** Every member is a copy, and a copy
is sanctioned only where a machine checks it ([links](../normative.md#links)): each is taken from
the release the project declares, refused at the edit, and verified against that release at
integration. The set is closed by this table, and a file no row names is the project's own.

## The copies are ordinary committed files

> **A vendored member is tracked content, committed like any other file. It reaches every clone,
> worktree, and resolution by `git pull` alone — no provisioning step touches it.**

A licence that is not in the tree is not a licence GitHub can show, and a workflow that is not in
`.github/workflows/` does not run: these members exist to be found at their paths by parties that
read nothing else. Tracked, they are consistent with the tree they sit in — a bisect sees the
policy in force at each step, and a fresh clone is complete. The record of which release they came
from is the project's declaration ([references](references.md#the-declaration)): one line,
project-owned, and the same line the documents resolve through, so there is no second stamp to
keep beside it.

## Updating the copy

> **The vendored set changes only by a commit that rides the project's normal gates.** This
> repository's release asks for that commit, by filing the item, in every project that trails.

On a schedule, and on every release ([release-cycle](release-cycle.md#the-release)), the process
walks every managed project and, where the project's declaration trails the latest release, files
**the upgrade item** ([norm-flows](norm-flows.md#upgrade)). The project's own flow resolves it into
one change: the declaration bumped to the release, every reference to this repository rewritten
at it ([references](references.md#references)), and every vendored member written as that release
has it. The filing is idempotent per project and release — a project already current gets
nothing, and one with the item open gets nothing more.

Catching up with the amended rules, rather than merely declaring them, is the second half of that
same item ([upgrade](norm-flows.md#upgrade)): the bump lands, and the walk that follows files what
the tree owes. It never blocks the bump, because the two are sequential steps of one resolution —
a tree declaring the new release with open gap items is the convention's normal state, while a
tree declaring an old one is behind in a way only the open upgrade item tracks. **The open upgrade
items and the gap items they file are the fleet's convergence status** — which projects are
behind, and by how much, is a label query, not a spreadsheet.

The process may use the network freely; it runs where the network is legitimate. What stays
hermetic is each project's commit gate.

## The check

Two positions, each doing what it is good at:

- **A guard, at the edit.** An agent proposing to modify a vendored member is refused before it
  happens, and the refusal carries the whole recovery: *this content comes from `org`; to change
  it, file an issue against `org`*. The blocked edit is not lost work — its substance becomes the
  issue's body. Cheap and immediate — and it fails open, which is why it is not the only check.
- **A gate, at integration.** Where the network is already legitimate, the gate compares every
  vendored member against the release the project declares. A member that does not match —
  however the edit happened — is **divergence**, and the change is refused.

**Staleness is never a failure.** A member that faithfully matches an *older* declared release is
behind, not wrong: being behind is the open upgrade item's job, and blocking every commit in every
project the moment a tag lands would make a release a fleet-wide outage. The commit gate does not
check members at all — it has no network and needs none, because the edit guard covers the agent
path and the integration gate covers every path.

## What this is not

- **Not the documents.** A document is referenced, never copied
  ([references](references.md#what-a-project-holds)); the members here are the files a reference
  cannot replace, because their reader is not a person following a citation but a platform
  reading a path.
- **Not provisioned.** Provisioning sets up what is deliberately untracked — tools and hooks. A
  licence is content another party reads from the tree, and content moves by commits.
- **Not git submodules.** A submodule's content is not in the tree until someone initializes it —
  absent from a fresh clone and from most tooling's view of the repo — and it makes every clone
  and worktree operation two-step. The members are ordinary tracked files.
- **Not subtree merges.** A subtree is editable in place, which reintroduces drift, and its merge
  history is noise in every log.
- **Not fetched at commit time.** The commit gate stays hermetic. The network lives where it
  already legitimately is: in the sync tool, and in the integration gate.

## Open questions

- The label the upgrade item carries beside its type. The reconciliation items carry the amended
  documents' tags ([reconciliation](../normative.md#reconciliation)), and the pass item carries
  all of them.
- Whether the edit guard's refusal of the vendored paths ships in the shared guard configuration
  for every project, or each project's own guard seam.
- How the shared `CONTRIBUTING.md` sections are checked, given they share a file with content the
  project owns — a marked region the gate compares, or the shared part becoming a document in this
  repository that the project's file references.
