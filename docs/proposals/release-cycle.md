# Release Cycle

> **Proposal.** Not normative: an end state under discussion, binding nothing until ratified.

> **Home:** [promise-language/org](https://github.com/promise-language/org) — this document is
> distributed into each managed project as `docs/org/`. A copy is never edited in place: to
> change it, file an issue against `org`.

How the corpus changes, and how a change reaches the fleet: the loop this repository runs, from
an issue filed against a rule to the reconciliation items every project receives.
[Distribution](distribution.md) owns the copies and what keeps them honest; [doc-sync](doc-sync.md)
owns the tools and CI; this document owns the process they serve — who acts, on what, in what
order — so that a cycle run by hand and a cycle run by machine are the same cycle.

## The loop

Five steps in one direction, each with one actor and one product:

| Step | Actor | Produces |
|---|---|---|
| **Intake** — a defect or gap in a rule is filed | anyone, from any project | an issue here, carrying the document's tag |
| **The amendment pass** — a document's open items are resolved together | the corpus maintainer | one reviewed change amending the documents, closing the items it incorporates |
| **The release** — the amended corpus is tagged | the corpus maintainer | a tag, and release notes stating the delta per document |
| **Dissemination** — every project receives the release | org CI and `bin/fleet` | a sync change and a reconciliation pass item per project |
| **Reconciliation** — a project catches up with the amended rules | the project | gap items under each document's tag, and the work that closes them |

Nothing in the loop pushes to a project. A project receives proposals — a change riding its
gates, an item in its queue — and its own process does the rest.

## Intake

An issue about a document is filed in this repository, carries the document's tag, and carries what
[reconciliation](../normative.md#reconciliation) requires of any item: the rule quoted, what fails
under it, and the replacement text or the request for one. A reader who meets the document in a
project's tree and files there has filed in the natural wrong place; the issue is transferred here,
not worked where it landed.

> **An issue is not worked alone.** It waits for the amendment pass, where it is read beside every
> other open item on the same document.

Two requests against one document, read separately, produce two amendments that may contradict
each other, and the third reader finds the contradiction. Read together, they produce one. The
wait is also what makes an amendment worth a release: a stream of one-line bumps costs every
project a sync and a pass apiece, for a delta none of them can act on. An issue that cannot wait
— a rule actively producing wrong work — says so, and the maintainer runs the pass early.

Triage is continuous and mechanical: an unlabelled issue gets its document's tag; a duplicate is
closed against the item it duplicates; an issue against a proposal is worked into the proposal
directly, since a proposal is freely rewritten and has no tag to wait under.

## The amendment pass

> **The pass reads every open item on every document, and the documents themselves, and lands
> one reviewed change.** Items are restructured before they are resolved — fewer, larger items,
> each one decision — rather than answered one increment at a time.

In order:

1. **Read the corpus whole**, for contradictions between documents and for gaps no item names.
   The items are what readers noticed; they are not everything there is.
2. **Resolve each document's items together** into amendments. An item is incorporated, merged
   into another, or closed with the reason. The wording an item proposes is a starting point,
   never a constraint; what it says must survive is.
3. **Amend**, in one change under review, every document the pass touched — the
   [lifecycle](../normative.md#lifecycle)'s amendment, landing before any release that carries it.
4. **Close what the change incorporates**, each closing comment naming the section that now
   carries the decision, so the item's reader can check the claim in a minute.
5. **Reconcile this repository** against the amended corpus
   ([reconciliation](../normative.md#reconciliation)). Its own tree,
   tools, and proposals are held to the rules like any project's, and its gaps are items here
   under each document's tag.

A pass that finds nothing to amend is a legitimate result: it closes or keeps its items with the
reason, and cuts no release.

## The release

> **A tag push is a release, nothing else is** (distribution). The tag is `docs-<year>.<n>`, with
> `n` counting the year's releases from 1, and it names the commit the copies are taken from.

Every release carries **release notes**, published with the tag as the repository's release
entry: for each amended document, what changed and what a project now looks for in its own tree.
The notes are the material the reconciliation pass in every project runs on, and they are written
by the maintainer who made the amendment, at the moment the delta is cheapest to state — a
project reading a raw diff of five documents does the maintainer's work over, once per project.

A release is cut once the pass has landed and this repository's own reconciliation items are
filed. Nothing else is a precondition: a release with open items here is the normal state,
exactly as a project with open gap items is.

## Dissemination

Distribution and doc-sync own the machinery. The process is that the sync is automatic and the
reconciliation is asked for:

- **The sync change** opens in every trailing project with nobody acting — on the release, and
  on the schedule that catches a project the release missed.
- **The reconciliation pass item** is filed in every project when the maintainer triggers it for
  the release, carrying the release notes and the amended documents' tags. It is triggered once
  the sync changes are open, so a project meets the new rules and the request to catch up with
  them together.

A step run by hand follows the specification of the `bin/fleet` command that performs it, exactly
— so a hand-run cycle and an automated one produce identical changes and items, and automating a
step changes who runs it, never what it does.

## Reconciliation in the project

The pass item is the project's own work: read the release notes against the tree; file a gap
item under each document's tag for every gap found — code the amended rule now forbids, a tool
the amended contract now binds, a project document the change now contradicts; close the pass
item when the gaps are filed. The gap items are then ordinary work, and the project's tag queries
are its distance from the corpus.

A gap the project judges to be a defect in the *rule* is not worked in the project. It is filed
here, and the loop begins again.

## Open questions

- The cadence of the pass and the release — a fixed period, a threshold of open items, or on
  demand only.
- Who the corpus maintainer is, and whether the role rotates.
- Whether the pass item is one per project per release, or one per project per amended document
  — shared with distribution.
- The release notes' fixed form, so that `fleet reconcile` carries them into the pass item
  verbatim.
