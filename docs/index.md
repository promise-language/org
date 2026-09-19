# Documentation Index

The map of `docs/`. The rules this tree is written under — which locations bind, the header
every document opens with, where status lives and where it never does, one fact one home, the
lifecycle, and what is checked mechanically — are [normative.md](normative.md)'s, and this file
does not restate them. It carries only what that shared document cannot: which project this is,
and its query.

This repository is the **home of the organization-wide corpus**. Every specification in its root
binds every managed project, which stands on it by reference at the release it declares
([references.md](references.md)); a rule changes here and nowhere else. How a change arrives
here and reaches the fleet is [release-cycle.md](release-cycle.md).

**The status query.** Every tagged document's header points here for the query that lists its
remaining work — in this repository, the remaining work on the *definition*, together with this
tree's own compliance gaps; in a project that stands on it, that project's compliance work. The
query is:

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
- [identity.md](identity.md) — What names a host, a guest, an arena, a tool and a process, and
  how each identity is created.
- [logging.md](logging.md) — How the development tools and the orchestration system log: the
  line, where it is written, its bounds, and the path to one store for the fleet.
- [release-cycle.md](release-cycle.md) — How the corpus changes and how a change reaches the
  fleet: intake, the amendment pass, the release, and the reconciliation that follows.
- [norm-flows.md](norm-flows.md) — The item types that carry a change to a normative document, and
  the steps each one runs.
- [references.md](references.md) — How a document in one repository cites one in another, and the
  declaration a project holds instead of a copy.
- [consolidation.md](consolidation.md) — The pass that restructures a repository's open items into
  one item per area of the product, current against the norms.

## Proposals

- [proposals/anchoring.md](proposals/anchoring.md) — Aspects whose modification requires a
  person's approval, and the path an approval travels.
- [proposals/distribution.md](proposals/distribution.md) — The files this repository places in
  every managed project's tree, and what keeps those copies honest.
- [proposals/org.md](proposals/org.md) — `bin/org`: what carries a release to every project that
  stands on this one, and what may run unattended.
