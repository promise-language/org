# Distribution

**Proposal. Not normative.** How a document ratified in this repository reaches every managed
project, and what keeps the copies honest.

A specification ratified here binds every project, and each project carries a copy in its own
tree — a rule that lives in another repo is not in an agent's context at the moment it has to be
followed. This document is about the machinery behind that sentence: where the copy comes from,
how it changes, and how a wrong copy is caught.

## The vendored set

The corpus is more than the guides, and not all of it sits under `docs/org/` — some members are
effective only at the conventional path other tooling reads:

| Member | Lands at | Form |
|---|---|---|
| The ratified documents | `docs/org/` | Byte-identical |
| `LICENSE`, `LICENSE-APACHE`, `LICENSE-MIT` | the repository root | Byte-identical — the pointer is worded without a repo name, and the MIT holder line is the one org-wide holder |
| The CLA workflow | `.github/workflows/cla.yml` | Byte-identical, including its skip-on-private guard |
| The shared `CONTRIBUTING.md` sections — CLA, licensing of contributions, commit identity | inside the project's own `CONTRIBUTING.md` | The one member that is not a whole file; see open questions |

One rule covers them all: **a fleet whose copies of a legal text or a policy check drift per
repository has several policies where it means to have one.** The stamp names every vendored
file wherever it lands, the edit guard refuses every byte-identical path, and the integration
gate verifies each against the claimed release.

## The copies are ordinary committed files

> **A project's `docs/org/` is tracked content, committed like any other file. It reaches every
> clone, worktree, and resolution by `git pull` alone — no provisioning step touches it.**

This is what makes the copies trustworthy in the one dimension that matters most: they are
always consistent **with the tree they sit in**. A resolution materialized at a commit sees
exactly the org docs that commit had; a bisect sees the rules in force at each step; a fresh
clone is complete. Nothing updates under a working tree, because nothing but a commit can change
what a working tree holds.

Alongside the documents, `docs/org/` carries a version stamp naming the org tag the copies came
from and listing every vendored file wherever it lands — the members outside `docs/org/`
included. The stamp is a claim, not a proof — what makes it honest is the check below.

> **Every document in the vendored set opens with its home line** — directly under the title,
> naming this repository as the original and the issue against `org` as the way to change it.

A reader of a project's copy is the reader that line exists for, and it must reach them through
the copies being byte-identical — so it is written **in the source**, worded to be true wherever
the file sits, never injected per copy by the sync tool. An injected banner would be content
nobody reviewed here, and the one thing the gate could no longer compare exactly.

## Updating the copy is an act, and CI is the actor

> **The vendored copy changes only by a commit that rides the project's normal gates.** A
> periodic CI process in this repository drives those commits for the whole fleet.

On a schedule — and on every docs release; a tag push is a release, nothing else is — the process
walks every managed project and, where the project's stamp trails the latest tag, produces two
things:

- **The sync change**: the tagged snapshot copied into `docs/org/` with its stamp, opened as an
  ordinary change that rides the project's own gates and review. Mechanical, byte-exact, and
  idempotent — a project already current produces nothing.
- **The reconciliation items**: the project has to *catch up with* the amended rules, not merely
  hold them. This is the reconciliation pass the normative-docs convention already requires after
  any amendment, run per project: walk the delta between the old and new rules against the
  project, and file an item for every gap found — code the new rule now forbids, tools the new
  contract now binds, project docs the change now contradicts — each carrying the amended
  document's tag. The CI files the pass itself as an item per project; expanding it into the
  concrete gaps is judgment work, and the project's ordinary resolution flow is what performs it.

The two outputs are deliberately separate: the sync change is never blocked on the
reconciliation, because a tree holding the new rules with open gap items is the convention's
normal state, while a tree holding old rules is wrong in a way nothing tracks. **The open sync
and reconciliation items are the fleet's convergence status** — which projects are behind, and by
how much, is a label query, not a spreadsheet.

The process may use the network freely; it runs where the network is legitimate. What stays
hermetic is each project's commit gate.

## The check: divergence is refused, staleness is an item

Two positions, each doing what it is good at:

- **A guard, at the edit.** An agent proposing to modify any vendored path — `docs/org/`, the
  legal files, the CLA workflow — is refused before it happens, and the refusal carries the
  whole recovery: *this content comes from `org`; to change it, file an issue against `org`*.
  The blocked edit is not lost work — its substance becomes the issue's body. Cheap and
  immediate — and it fails open, which is why it is not the only check.
- **A gate, at integration.** Where the network is already legitimate, the gate compares every
  vendored file against the org tag the stamp claims. A copy that does not match its claimed
  tag — however the edit happened — is **divergence**, and the change is refused.

**Staleness is never a failure.** A copy that faithfully matches an *older* tag is behind, not
wrong: being behind is the open sync item's job, and blocking every commit in every project the
moment a tag lands would make a docs bump a fleet-wide outage. The commit gate does not check
org content at all — it has no network and needs none, because the edit guard covers the agent
path and the integration gate covers every path.

## What this is not

- **Not provisioned.** Provisioning sets up what is deliberately untracked — tools, hooks,
  arena plumbing. The docs are content, and content moves by commits. A provisioning step that
  rewrote tracked files would dirty trees it does not own; one that delivered docs outside the
  tree would put the rules somewhere agents do not read.
- **Not git submodules.** A submodule's content is not in the tree until someone initializes it —
  absent from a fresh clone, from most tooling's view of the repo, and from the agent context the
  whole mechanism exists to fill — and it makes every clone and worktree operation
  two-step. The copies are ordinary tracked files.
- **Not subtree merges.** A subtree is editable in place, which reintroduces drift, and its merge
  history is noise in every log.
- **Not fetched at commit time.** The commit gate stays hermetic. The network lives where it
  already legitimately is: in the sync tool at resolution time, and in the integration gate.

## Open questions

- The stamp's form — one file naming the tag, or the tag plus per-file hashes so the integration
  gate can name exactly which file diverged without a full diff.
- The labels the sync and reconciliation items carry. The machinery itself — the tools, the
  workflows, and their triggers — is drafted in [doc-sync](doc-sync.md); both live
  in this repository.
- The reconciliation pass's shape: one pass item per project per release, or one per amended
  document.
- Whether the edit guard's refusal of the vendored paths ships in the shared guard
  configuration for every project, or each project's own guard seam.
- How the shared `CONTRIBUTING.md` sections are checked, given they share a file with content
  the project owns — a marked region the gate compares, or a `CONTRIBUTING.md` split so the
  shared part is its own vendored file the project's file includes by link.
