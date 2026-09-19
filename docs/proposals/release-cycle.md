# Release Cycle

> **Proposal.** Not normative: an end state under discussion, binding nothing until ratified.

> **Tag:** `release-cycle` — remaining work to complete this document: the query named in
> `docs/index.md`.

How the corpus changes, and how a change reaches the fleet: the loop this repository runs, from
an issue filed against a rule to the upgrade every project receives and the reconciliation it
runs. [References](references.md) owns how a project stands on the corpus; this document owns the
process it serves — who acts, on what, in what order — so that a cycle run by hand and a cycle run
by machine are the same cycle.

## The loop

Five steps in one direction, each with one actor and one product:

| Step | Actor | Produces |
|---|---|---|
| **Intake** — a defect or gap in a rule is filed | anyone, from any project | an issue here, carrying the document's tag |
| **The amendment pass** — every open item is resolved beside every other | the corpus maintainer | one reviewed change amending the documents, and an item for each question it could not decide |
| **The release** — the amended corpus is tagged | the corpus maintainer | a tag, and release notes stating the delta per document |
| **Dissemination** — every project receives the release | the release's own flow, through `bin/org` | an upgrade item per trailing project |
| **Reconciliation** — a project catches up with the amended rules | the project | gap items under each document's tag, and the work that closes them |

Nothing in the loop pushes to a project. A project receives proposals — a change riding its
gates, an item in its queue — and its own process does the rest.

## Intake

An issue about a document is filed in this repository, carries the document's tag where the
filer knows which document it is, and carries what
[reconciliation](../normative.md#reconciliation) requires of any item: the rule quoted, what
fails under it, and the replacement text or the request for one. An issue about a rule no
document states yet carries no tag, and the pass places it. A reader who meets the document in a
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
closed against the item it duplicates; an issue against a proposal carries that proposal's tag
like any other ([header](../normative.md#header)), and is worked into the proposal directly,
since a proposal is freely rewritten and binds nothing while it waits.

An issue on a question an open item already holds ([the amendment pass](#the-amendment-pass))
says so and names it, so the pass reads the two together rather than as two subjects.

## The amendment pass

> **The pass reads every open item on every document, and the documents themselves, and lands
> one reviewed change.** Items are restructured before they are resolved — fewer, larger items,
> each one decision — rather than answered one increment at a time. The pass amends
> specifications; it ratifies nothing.

In order:

1. **Read the corpus whole**, for contradictions between documents and for gaps no item names.
   The items are what readers noticed; they are not everything there is. A contradiction or a
   gap found here is a pass item like any filed one.
2. **Resolve every item** to exactly one outcome — **incorporate**, into a named section, or
   into a document created for it where no existing one is its home;
   **merge**, into the item it is one decision with; **rework**, where the gap is real and the
   proposed wording is not the fix; **close**, as against the corpus's direction, a duplicate, or
   already carried by a named section; or **carry**, into an item the pass files for the question,
   with the reason the pass cannot decide it. The wording an item proposes is a starting point,
   never a constraint; what it says must survive is. An item whose only home would be a proposal
   is not resolved into one and called closed: either the deciding rule is carried into a
   specification, or the proposal is ratified — which is its own item and its own flow
   ([ratification](norm-flows.md#ratification)), never something an amendment does on the way
   past.
3. **Amend**, in one change under review, every document the pass touched — the
   [lifecycle](../normative.md#lifecycle)'s amendment, landing before any release that carries it.
   Every amendment states the rule and what it decides; prose that decides nothing is left out.
4. **Close what the change resolves**: each incorporated item with a comment naming the section
   that now carries the decision, each rejected item with its reason, each carried item against
   the item that now holds its question — so the item's reader can check the claim in a minute.
5. **Reconcile this repository** against the amended corpus
   ([reconciliation](../normative.md#reconciliation)). Its own tree, tools, and proposals are held
   to the rules like any project's, and its gaps are items here, one per document.

> **What the pass cannot decide it files**, as an item under the document's tag carrying the
> decision, its candidates, and the evidence — never as prose saying the matter is open
> ([end state](../normative.md#end-state)).

Every item open when a pass begins ends closed — incorporated, rejected, or carried into the item
that now holds its question — and what the pass carried is open in its place, so the tag query is
the whole of what the document still owes: the rules not yet met, and the rules not yet decided.
A carried item earns its place by being a decision rather than a note, and it is where the next
request on that concern says it belongs, so two people asking the same question meet in one
place.

A pass that finds nothing to amend is a legitimate result: it closes or keeps its items with the
reason, and cuts no release.

## The release

> **A tag push is a release, nothing else is.** The tag is `v<major>.<minor>.<patch>`, annotated,
> and it names the commit a project stands on when it declares the release.

> **The number says what adopting it costs.** **Major**: a citation may break — a heading renamed,
> a document moved, renamed or retired. **Minor**: rules changed or were added and no citation
> breaks. **Patch**: no rule changed, so nothing in any dependent can have become wrong.

A dependent reads the number before it reads the notes, and the three answer the question it
actually has: whether this release will stop its upgrade, give its tree work, or cost nothing.
Major is exactly the case where the upgrade's carry check fails and a person has to say what a
reference should now cite ([upgrade](norm-flows.md#upgrade)); minor is the case where the bump
lands unattended and the walk that follows finds the work; patch is the case where both are
mechanical. A number that counted releases instead would tell a dependent when this one was cut,
which is the one thing the tag list already says.

> **Every release carries release notes, published with the tag as the repository's release
> entry, in one form: per amended document, what changed, and what a project now looks for in
> its own tree.** The notes are what `org sync` carries into each upgrade item, verbatim, and
> they carry what any home owes its dependents ([the two sides](references.md#the-two-sides)):
> every heading renamed and every document moved or retired.

The notes are the material every dependent's upgrade runs on, and they are written
by the maintainer who made the amendment, at the moment the delta is cheapest to state — a
project reading a raw diff of five documents does the maintainer's work over, once per project.

A release is cut once the pass has landed and this repository's own reconciliation items are
filed. Nothing else is a precondition: a release with open items here is the normal state,
exactly as a project with open gap items is.

## Dissemination

The process is one request per trailing project, filed by the release itself with nobody else
acting:

- **The upgrade item** carries the release notes. The project's own flow resolves it into the
  change that bumps the declaration and rewrites the references, lands that, and then walks the
  tree against the notes and files what it owes
  ([upgrade](norm-flows.md#upgrade)). A project the release missed is one `org status` reports as
  trailing, and a person files its item.
- **Adopting and meeting are not two items.** The bump lands before the walk begins, so the tree
  that holds the new rules with open gap items — the convention's normal state — is reached in
  one resolution rather than two, and nothing waits on a second filing to be told that a tree
  which has just adopted a rule should be measured against it.

A step run by hand follows the specification of the `bin/org` command that performs it, exactly
— so a hand-run cycle and an automated one produce identical changes and items, and automating a
step changes who runs it, never what it does.

## Reconciliation in the project

The walk is the project's own work, on the upgrade item that carried the notes: read the release
notes against the tree; file the gaps found — code the amended rule now forbids, a tool the
amended contract now binds, a project document the change now contradicts — as items sized by
area rather than one per document ([reconciliation](../normative.md#reconciliation)), each
carrying the tags it answers to; the upgrade item finalizes when the items are filed. The gap
items are then ordinary work, and the project's tag queries are its distance from the corpus.

A gap the project judges to be a defect in the *rule* is not worked in the project. It is filed
here, and the loop begins again.
