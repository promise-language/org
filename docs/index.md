# Documentation Index

The map of `docs/`. The rules this tree is written under — which locations bind, the header
every document opens with, where status lives and where it never does, one fact one home, the
lifecycle, and what is checked mechanically — are [normative.md](normative.md)'s, and this file
does not restate them. It carries only what that shared document cannot: which project this is,
and its query.

This repository is the **home of the organization-wide corpus**. Every specification in its root
binds every managed project and is distributed to each as `docs/org/`; the root here *is* the
corpus, so there is no `docs/org/`, and a rule changes here and nowhere else. How a change
arrives here and reaches the fleet is [proposals/release-cycle.md](proposals/release-cycle.md).

**The status query.** Every specification's tag line points here for the query that lists its
remaining work — in this repository, the remaining work on the *definition*, together with this
tree's own compliance gaps; in a project's copy, that project's compliance work. The query is:

> `gh issue list --label <tag> --state open --limit 200`

## Specifications

- [normative.md](normative.md) — What makes a document in a managed project binding, and the one
  docs structure every project holds.
- [engineering-guide.md](engineering-guide.md) — How code in this organization is written, in
  any language.
- [engineering-guide-promise.md](engineering-guide-promise.md) — The engineering guide applied
  to Promise source.
- [engineering-guide-go.md](engineering-guide-go.md) — The engineering guide applied to Go
  source.
- [cli-guide.md](cli-guide.md) — How every command-line tool behaves at its invocation surface.

## Proposals — not binding

- [proposals/anchoring.md](proposals/anchoring.md) — Aspects whose modification requires a
  person's approval, and the path an approval travels.
- [proposals/distribution.md](proposals/distribution.md) — How a ratified document reaches every
  managed project, and what keeps the copies honest.
- [proposals/doc-sync.md](proposals/doc-sync.md) — The org tools and org CI that drive the
  distribution loop, and what may run unattended.
- [proposals/release-cycle.md](proposals/release-cycle.md) — How the corpus changes and how a
  change reaches the fleet: intake, the amendment pass, the release, and the reconciliation that
  follows.
