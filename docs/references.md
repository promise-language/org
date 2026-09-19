# References

> **Tag:** `references` — remaining work to complete this document: the query named in
> `docs/index.md`.

> **Home:** [promise-language/org](https://github.com/promise-language/org) — this document
> changes here and nowhere else. To change it, file an issue against `org`.

How a document in one repository cites a document in another, and what a project holds instead
of a copy: a declaration of the repositories it stands on and the release of each. One reference
system for every document in the organization — the corpus this repository holds is one
dependency in it, declared as any other — and no document exists in two trees.

A fact has one home ([one home](normative.md#one-home)), and a copy of a document, however it
is checked, is that fact in a second tree: two versions in every project, the project's own and
the corpus's beside it, and a reader deciding which is the real one. A reference names the home
and the release, and puts nothing in a second tree. What an agent must have in its context it
reads by following the reference, at the release the declaration names.

## Identity

> **A document is `<repository>/<name>`, and a section is `<repository>/<name>#<slug>`.** The
> repository is its GitHub name in the organization; the name is the document's basename minus
> `.md` — its tag, where it has one ([header](normative.md#header)); the slug is the section's
> ([sections](normative.md#sections)).

`org/normative#reconciliation`, `base/gate-contract#the-exec-line`, `promise/language-design`.
The name is unique across the repository's `docs/` — root, `proposals/`, `archive/`,
`research/` — so the identity carries no path and survives the moves the
[lifecycle](normative.md#lifecycle) makes: ratifying or retiring a document changes its
location, its header, and its index entry, and its identity not at all.

Inside a tracked Markdown document a reference is written as a link ([references](#references)),
which carries the identity's parts in a form a reader can follow. The identity itself is the
whole citation everywhere a link cannot go — an item, a comment, a step prompt, a verdict, and
the argument the resolver takes.

## The declaration

> **A project declares, in `docs/index.md`, every repository whose documents it references, and
> one release of each.** The declaration is the one place "which rules does this project stand
> on" is written, and adopting a newer release is an edit to it.

```markdown
## Dependencies

| Repository | Release |
|---|---|
| `org` | `v1.0.0` |
| `base` | `v0.4.0` |
```

The index is the one file in the root that is not a specification and already the home of the
per-project facts the shared documents cannot carry ([location](normative.md#location)); a
second file would be a sixth location. A release is a tag the repository pushed —
[release-cycle](release-cycle.md#the-release) for this one — never a branch, whose target moves
with no change in the citing tree, and never a bare commit, which is a release nobody announced
and whose notes nobody wrote. A dependent that needs what the home has not released asks the home
to release it. One release per repository, never a combination of its documents that never
existed together.

This repository is one repository among them: it declares the projects whose documents it cites,
at releases they cut, and is declared by every project at the release it holds them to.

## References

> **A reference to another repository's document is a link to that document at the release the
> declaration names**, and its text is the heading or the phrase the sentence needs
> ([sections](normative.md#sections)). The URL carries the repository, the release, the
> document's path at that release, and the section's slug:

```markdown
[the exec line](https://github.com/promise-language/base/blob/v0.4.0/docs/gate-contract.md#the-exec-line)
```

The reader is why the link is written out. A citation only a tool can follow is a citation most
readers will not follow, and a reference nobody opens is a reference nobody checks. The release
in the URL is what holds the target still — a link into a branch moves with no change in the
citing tree — and it is the release the declaration names, so what a sentence cites and what the
project stands on cannot disagree. Inside a repository a reference stays a relative link, because
there the target is in the same tree at the same commit.

> **Upgrading a declaration rewrites, in the same change, every reference to that repository.**
> The release appears in each of them, so the row and the links move together or the change is
> refused.

The rewrite is mechanical — the release segment, and the path with it where the document moved
between releases — so a reviewer reads the declaration row and skims the rest, and the release a
reader clicks into is the one the reviewer approved. Carrying the release in each reference is
what makes the pin checkable without the network: the check reads the URL and the declaration and
compares them, where a bare name would have to fetch something to learn anything at all.

> **An upgrade carries every reference or it does not happen.** Where each one resolves at the
> new release, the rewrite needs no judgment and whatever performs it may. Where one does not —
> a section renamed, a document retired, a fact moved to another document — the upgrade stops
> and names each reference that stopped it, and a person decides.

A reference that cannot be carried forward is the [mechanical
checks](normative.md#mechanical-checks)' case exactly: it is repaired by repointing it at
what now owns the fact, or by deleting it together with the claim it supports, and the two are
not the same decision. Something choosing between them to finish a bump would be answering a
question about the citing document while doing arithmetic on a URL. So the upgrade is all of the
references or none of them: a tree half-bumped states two releases at once, and the pin stops
meaning anything.

> **A specification references only specifications.** A proposal may reference a proposal.

A binding document standing on an unratified one lets authority in through a door the
[location](normative.md#location) table keeps shut: what binds would depend on what does not. A
proposal is under discussion, and so may be what it cites.

## The two sides

> **Every reference has a home and a dependent.** The **home** is the repository that owns the
> document and releases it; the **dependent** is the repository that declares the home and cites
> the document. A repository is a home to some and a dependent of others, and the two roles carry
> different obligations.

> **A home cuts releases, and each release's notes name what a dependent must act on**: every
> document amended and what changed in it, every heading renamed, and every document moved,
> retired, or renamed. A repository that releases nothing is a repository nothing may cite.

A home renames a heading and repairs every reference in its own tree, because
[sections](normative.md#sections) requires it; the references it cannot repair are the ones
in other repositories, and it does not know where they are. The notes are the whole mechanism
that stands in for that: a dependent's upgrade stops on exactly the changes named above, and the
notes are where it finds what the reference should say instead. A rename a release does not
announce is a dependent's stopped upgrade with nothing to read.

> **A dependent upgrades as one reviewed change, when it has a reason to.** Nothing files it for
> the dependent, and trailing a release is not a defect: the release it declares still exists and
> still says what it said.

> **The corpus is the one exception: what binds is pushed, what is merely cited is pulled.** The
> corpus's home files the upgrade in every trailing project, and the upgrade goes on to reconcile
> the project with the amended rules ([dissemination](release-cycle.md#dissemination)).

One item, two acts, and neither stands in for the other. The bump moves the declaration and the
references and nothing else; a project that lands it holds the new rules and has not yet met them.
The walk that follows is the judgment the bump cannot do: read the notes against this tree, and
file what now fails. The item carries the notes and the amended documents' tags, never a list of
gaps — the gaps are found in the project, by the project, and each becomes its own item under the
tag it belongs to. That is also why the bump never waits on the walk: a tree holding the new rules
with open gap items is the normal state, and a tree holding the old ones is behind in a way nothing
tracks.

A project trailing the corpus is behind on the rules it is measured against, and its own tag
queries are meant to state its distance from them — so convergence there is the fleet's business.
A project trailing another project is standing on a contract it chose at a release that has not
moved. Filing work in a repository because an unrelated repository tagged would spend a review in
every dependent on every release, which is how a fleet learns to merge upgrades without reading
them.

## Visibility

> **A reference never widens visibility.** A document in a public repository references only
> public repositories, and no reference leaves the organization.

A public document citing a private one fails as a reference for nearly every reader while looking
like a working one, and announces a path it cannot show.

## Releases order the graph

> **Repositories may cite each other.** A declaration names a release that already exists, so the
> versioned graph — each document at its release, citing others at theirs — is acyclic by
> construction, and no rule about the repository graph is kept.

forge cites workspace's tool contract and workspace cites forge's tooling; each does so at a
release the other already cut, and each upgrades its declaration when it chooses, never both in
one step. What a graph rule would guard against — one fact defined in two places, each in terms
of the other — is [one home](normative.md#one-home)'s rule already: the model in the document
that owns the concept, the contract in the reference for the surface, each linking to the other
once. A reader following a chain may arrive at an older release of a document than the one they
left, because a document's references resolve through its own repository's declaration at the
release being read. That is what a pin means.

> **A release never obliges the home's dependents to act, so no release triggers another.** Two
> repositories that cite each other release independently, and neither one's tag starts work in
> the other.

The only upgrades anything files are the corpus's, and nothing files an upgrade in the corpus's
home ([the two sides](#the-two-sides)) — so however the citations run, what is triggered runs
one way, from the corpus outward, and stops. Were a release to oblige each dependent to upgrade,
a pair citing each other would each be answering the other's last release forever, and a fleet
would spend itself carrying version numbers around a ring.

## What a project holds

> **A tree holds no copy of another repository's document, tracked or not.** What it holds is
> the declaration. A reference is followed by the **resolver** — a workspace tool that fetches
> the named document at the declared release and returns the section — and the same resolver
> serves an agent, a person, and the integration gate. It is both a workspace command-line tool,
> which prints the document or the section to the terminal for a person, and an MCP server over
> the same binary, which is how an agent's harness reaches it: one implementation, two surfaces.

The rule a project is measured against has to be in the agent's context at the moment it is
followed. A tracked copy put it there at the cost of a second home. A provisioned copy would put
it there at the cost of a copying step, a set to police, and a reader still choosing between two
files on disk. Following the reference costs one call, and an agent that is running has the
network already — it needs it to think. So a step prompt cites `org/engineering-guide`, the
agent's first act is to read it, and what it reads is the release its tree declares. To learn what
a repository holds, the resolver returns its `docs/index.md` at the declared release: the map,
one line per document with its binding status ([location](normative.md#location)). A
repository not yet declared gets its row first, and finding its latest release is the resolver's
other network act. Whatever the resolver keeps between calls is its own affair and no rule's
subject: nothing lists it, nothing edits it, and nothing is checked against it.

## The checks

Two positions, each doing what it is good at:

- **At commit, hermetic.** The declaration parses; every cross-repository link names a declared
  repository and carries exactly the release declared for it, with no branch and no bare commit;
  every relative link resolves ([mechanical checks](normative.md#mechanical-checks)). Nothing
  is fetched, so nothing here says a target exists — that is the next position's.
- **At integration, where the network is legitimate.** Every reference resolves, at the release it
  names, to a document that exists there and a section that exists in it; each declared release
  exists at the repository it names; the visibility rule holds. An unresolvable reference is a
  failure, not a skip: nothing to check and nothing was checked must never look alike.
