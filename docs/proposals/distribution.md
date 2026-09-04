# Distribution

**Proposal. Not normative.** How a document ratified in this repository reaches every managed
project, and what keeps the copies honest.

A specification ratified here binds every project, and each project carries a copy in its own
tree — a rule that lives in another repo is not in an agent's context at the moment it has to be
followed. This document is about the machinery behind that sentence: where the copy comes from,
how it changes, and how a wrong copy is caught.

## Three parties

| Party | Lives | Changes when |
|---|---|---|
| **The source** | this repository's `docs/` root | a rule changes — an ordinary reviewed amendment, here |
| **The reference** | with the provisioned tools, **outside the tracked tree** | the tools are provisioned — the workspace release pins one org version and carries its snapshot |
| **The copy** | tracked, in the project's `docs/org/` | a deliberate commit in that project, and nothing else |

## Setup never dirties the tree

> **Provisioning updates tools and the reference only. It never modifies a tracked file.**

A provisioning run that leaves the tree dirty has made a change nobody committed: it bypasses the
commit gate, surprises whoever owns the working tree, and turns "run setup" into an action with
two unrelated meanings. Setup delivers the new reference alongside the binaries; the tracked copy
stays exactly as the last commit left it.

## Updating the copy is an act

> **The vendored copy changes only by a commit that rides the project's normal gates.**

A sync tool writes the pinned snapshot into `docs/org/` and stops — it stages nothing and commits
nothing; the commit that follows is reviewed and gated like any other change. When the pin moves,
an item is filed in each managed project naming the new version, and the ordinary resolution flow
drives the sync commit — so the update is scheduled, attributed, and recorded, never a side
effect of something else.

## Kept in sync: the loop

Sync is a pipeline of deliberate acts, each one the trigger for the next, and its order is what
makes every step hermetic:

1. **An amendment lands here**, reviewed like any change to a specification.
2. **This repository cuts a docs release** — a tag; a tag push is a release, nothing else is.
3. **The workspace release bumps its org pin** and carries the tagged snapshot and its manifest —
   the reference.
4. **Provisioning delivers the new reference** to every project alongside the tools. The tree
   stays clean; nothing has changed in any project yet.
5. **A sync item is filed in each managed project**, naming the tag — filed by the automation
   that bumped the pin, so an item never precedes the reference it needs.
6. **The project's ordinary resolution flow works the item**: the sync tool copies `docs/org/`
   from the local reference — no network — and the commit rides the project's own gates.
7. **The open sync items are the fleet's convergence status.** Which projects are behind is a
   label query, not a spreadsheet — the same rule as every other gap.

## The check: divergence is refused, staleness is an item

The shared precommit set compares the tracked `docs/org/` against the reference that arrived with
the tools:

- **Divergence** — a copy matching **no** known org version — refuses the commit, naming the file
  and pointing here: the copy's home is this repository, and the edit belongs upstream.
- **Staleness** — a copy matching an **older** org version — is never a commit failure. Being
  behind is not a defect in the change being committed; it is the open sync item's job. Blocking
  every commit in every project the moment the pin moves would make a docs bump a fleet-wide
  outage.

> **The reference comes from outside the tree it judges.** It arrives with the provisioned
> binaries, so the party whose change is being checked — the tree, which an agent can edit — never
> holds the thing it is checked against. A manifest committed beside the copies could be edited in
> the same commit as the copies; the reference cannot.

## What this is not

- **Not git submodules.** A submodule's content is not in the tree until someone initializes it —
  absent from a fresh clone, from most tooling's view of the repo, and from the agent context the
  whole mechanism exists to fill — and it makes every clone and worktree operation
  two-step. The copies are ordinary tracked files.
- **Not subtree merges.** A subtree is editable in place, which reintroduces drift, and its merge
  history is noise in every log.
- **Not fetched at commit time.** A check that reaches for the network makes verify non-hermetic
  and turns an outage into a commit block. The reference is already local, delivered with the
  tools.

## Open questions

- The reference's transport — embedded in the check's binary, or a data file beside it.
- Whether the reference carries every historical manifest (they are small, and it makes
  "divergence = matches nothing ever" exact) or a bounded window.
- Whether the sync tool belongs to the workspace toolset or to this repository's own tooling.
- The label the sync items carry.
