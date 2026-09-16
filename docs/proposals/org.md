# Org

> **Proposal.** Not normative: an end state under discussion, binding nothing until ratified.

> **Tag:** `org` — remaining work to complete this document: the query named in
> `docs/index.md`.

`bin/org`: what carries a release to every project that stands on this one, and what may run
unattended. It syncs no document — a document is referenced at the release a project declares
([references](references.md)), never copied into it — so what this machinery moves is items, and
the handful of files [distribution](distribution.md) still places in each tree.

The tool lives in this repository's `tools/` and compiles into `bin/` like every other project's
tooling. A managed project implements none of it, and receives no change from it: what arrives
in a project is an item, and the project's own flow makes every change that follows
([norm-flows](norm-flows.md#the-types)).

> **Every command that files runs as a step of a release's flow**
> ([release](norm-flows.md#release)), from the account of whoever is resolving that item, under
> the same rule as any other step. `status` changes nothing, so anything may run it, as often as
> it likes.

Whatever runs the report on a schedule, and wherever it publishes what it found, is the
orchestrator's business and not this tool's: a report is read by a person who then decides, and
a reader who never comes costs nothing. What may not happen is the other half — something that
files on a schedule is an actor nobody named, and this repository is the worst place for one,
since a single mistaken run arrives in every project at once.

This tool is also the first consumer of the [CLI guide](../cli-guide.md): it follows it from its
first commit, so the org's own machinery is the guide's proving ground rather than its first
exception.

## The tool

One binary, `bin/org`, with the closed command set. Every command that files anything is run
by a step of a release's flow ([release](norm-flows.md#release)); nothing here is an actor of its
own:

| Command | Does | Mutates |
|---|---|---|
| `org status` | For every managed project: the release it declares, the latest release, and its open `norm: upgrade` and `norm: reconciliation` items — and, where an upgrade is waiting on a person, each reference that stopped it. For every other declared dependency in the fleet, the release the dependent declares beside the latest its home has cut, so a trailing dependent is visible without anything being filed in it ([the two sides](references.md#the-two-sides)). The convergence label query, rendered. | Nothing |
| `org sync` | For every project whose declaration trails the latest release: make the tags and item types the project lacks, then file the `norm: upgrade` item carrying the release notes. The project's own flow bumps the declaration, rewrites every reference to this repository, writes the vendored members, and then walks its tree against the notes ([norm-flows](norm-flows.md#upgrade)). Idempotent per (project, release) — a current project, or one with the item already open, gets nothing. | Labels, issues |

The managed-project list is a committed file in this repository — a closed set, extended by an
ordinary reviewed change, so "managed" is a fact in one place rather than a convention.

## Triggers

> **Only `status` runs on a schedule, because it changes nothing.** `sync` and `reconcile` file
> items, so each is a step of a release's flow and runs when that step runs
> ([every change](norm-flows.md#the-types)).

A schedule that filed items would be an actor nobody named, and the fleet is where that costs
most: one scheduled mistake arrives in every project at once. Scheduled reporting has the
opposite property — `status` run hourly and read by nobody costs nothing, and it is what shows a
project the release missed, so that a person can file its item.

> **What is filed is a request, and every judgment it needs happens in the project that receives
> it.** The item asks a project to adopt the release and then measure itself against it; when
> that happens, how large the resulting items are, and what a moved reference should now cite are
> all decided there, by the people who live with the answers ([upgrade](norm-flows.md#upgrade)).

Nothing here decides for a project, and nothing here waits on one either. That is what lets the
filing be mechanical without being presumptuous: a release asks every trailing project the same
question on the same day, and each answers on its own clock, in its own order, under its own
gates and review.

## What it must never do

- **Write to a project's tree.** It files items, and the project's own flow produces every
  change ([norm-flows](norm-flows.md#the-types)), riding the project's gates and review; the
  automation holds no authority a contributor lacks.
- **File a duplicate.** Idempotency is keyed on (project, release) for every item; a rerun after
  a partial failure completes the remainder and repeats nothing.
- **Exceed its subject.** Its items are about the declaration, the references to this
  repository, and the vendored members, and about nothing else in a managed project.

## Open questions

- Whether `sync` and `reconcile` stay commands of this binary or become a library the
  orchestrator calls at the step. They are run by a step and by nothing else, so the case for a
  command is that a person resolving the release runs it and can see what it would do; the case
  against is a surface that exists for one caller.
- What identity the scheduled report runs as, and where it publishes. The filing commands need
  no answer: they run from the account of whoever is resolving the item, and reach another
  repository with that person's own rights.
- Whether the reconciliation pass drafts its items with an agent's help under the human trigger,
  or the trigger hands a person a prepared delta and the filing stays manual at first.
- What `org status` owes the engagement feed once one exists: convergence as an article, not a
  terminal report.
- Whether the vendored members ([distribution](distribution.md)) ever need items of their own,
  `org: <type>` beside the `norm: <type>` set. Today they ride the `norm: upgrade`, because they
  come from the release the declaration names and land in the change that bumps it. A member that
  had to move on its own — a licence changed between releases — would be the case that mints the
  namespace, and it does not exist yet.
