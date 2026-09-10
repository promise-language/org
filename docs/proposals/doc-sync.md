# Doc Sync

> **Proposal.** Not normative: an end state under discussion, binding nothing until ratified.

> **Home:** [promise-language/org](https://github.com/promise-language/org) — this document is
> distributed into each managed project as `docs/org/`. A copy is never edited in place: to
> change it, file an issue against `org`.

The machinery behind [distribution](distribution.md)'s copies and
[release-cycle](release-cycle.md)'s loop — what runs it, from where, and what may run unattended.
Everything described here is **org tools and org CI**: the tools live in this repository's
`tools/` and compile into `bin/` like every other project's tooling; the workflows live in this
repository's `.github/workflows/`. A managed project implements none of it — projects only
receive changes and items.

These tools are also the first consumers of the [CLI guide](../cli-guide.md): they follow it from
their first commit, so the org's own machinery is the guide's proving ground rather than its
first exception.

## The tool

One binary, `bin/fleet`, with the closed command set:

| Command | Does | Mutates |
|---|---|---|
| `fleet status` | For every managed project: the stamp it holds, the latest release, and the open sync and reconciliation items. The convergence label query, rendered. | Nothing |
| `fleet sync` | For every project whose stamp trails the latest release: branch, write the vendored set and its stamp, open the change against the project's mainline, riding its gates and review. Idempotent — a current project, or one with the sync change already open, produces nothing. | PRs |
| `fleet reconcile` | For a named release: create the amended documents' labels where a project lacks them, then file the reconciliation pass item in each managed project, carrying the release notes and the amended documents' tags — the material the pass runs on. Idempotent per (project, release). | Labels, issues |

The managed-project list is a committed file in this repository — a closed set, extended by an
ordinary reviewed change, so "managed" is a fact in one place rather than a convention.

## Triggers: mechanical runs unattended, judgment is asked for

> **The scheduled path does only mechanical work.** `status` and `sync` run on a schedule and on
> every release: byte-exact, idempotent, no judgment anywhere in them.

> **Reconciliation is triggered by a person.** Filing the pass item is mechanical, but its
> arrival starts judgment work in every project at once — reading a rule delta against each tree
> and deciding what it now owes. A person decides when the fleet takes that on, triggers the pass
> for a release, and reviews what it files.

The asymmetry is deliberate and mirrors how authority widens elsewhere in the organization: start
with the narrow grant, widen when the narrow one has earned it. If experience shows the pass
files accurate items without supervision, arming it on release is one deliberate change here —
the reverse migration, un-trusting an automation that was wrong at scale, has no such single
step.

## What it must never do

- **Push to a project's mainline.** Every change it produces rides the project's own gates and
  review as an ordinary proposal; the automation holds no authority a contributor lacks.
- **File a duplicate.** Idempotency is keyed on (project, release) for items and on the open
  change for syncs; a rerun after a partial failure completes the remainder and repeats nothing.
- **Exceed its subject.** It writes the vendored set — each member at the path the stamp names —
  and files items about it; it never touches any other path in a managed project.

## Open questions

- The identity it acts as — a GitHub App or a fine-grained token — and the name its changes and
  items carry.
- Whether the reconciliation pass drafts its items with an agent's help under the human trigger,
  or the trigger hands a person a prepared delta and the filing stays manual at first.
- The fleet manifest's shape: one list, or per-project subsets of the vendored set (every repo
  takes the legal files; whether every repo takes every guide is a real question).
- What `fleet status` owes the engagement feed once one exists: convergence as an article, not a
  terminal report.
