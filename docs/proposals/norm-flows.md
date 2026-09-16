# Norm Flows

> **Proposal.** Not normative: an end state under discussion, binding nothing until ratified.

> **Tag:** `norm-flows` — remaining work to complete this document: the query named in
> `docs/index.md`.

Every change to a tree is a flow resolving an item, and changing a rule is no exception: the
amendment that changes it, the upgrade that adopts a home's release, the reconciliation that
catches a project up. This document names the item types that carry that work and the steps each
one runs — what a step does, what it decides, and which step follows it.

> **These flows carry a change to a normative document and nothing else**: a specification in a
> `docs/` root, which [location](../normative.md#location) makes binding, or a proposal written
> as the specification it would become. A defect anywhere else — a README, a research note, a
> comment, a file nobody is measured against — is an ordinary item on the ordinary route.

What the weight here buys is that a rule changes where it is defined and reaches everyone
measured against it, and a file nobody is measured against needs none of it. A wrong sentence in
a README is fixed by the change that fixes it.

Every repository is the home of its own specifications, so these flows run wherever there are
norms: the corpus in this repository, a project's own `docs/` root in that one. The release side
engages only where another repository references the document — a home nobody cites amends its
own norms and owes no one a release.

What each act *is* belongs to the documents that own it: [release-cycle](release-cycle.md) the
corpus's loop, [references](references.md) what a home and a dependent owe each other,
[normative](../normative.md#reconciliation) what any item must carry. Nothing here restates them.
A step says when the flow reads them and what it does with what it finds.

## The types

> **An item carries a type, and the type decides its steps.** The type is a required field on
> every item, not a convention laid over one: an orchestrator that cannot carry it cannot run
> these flows, and how it stores the field — a field of its own, or a label, as a GitHub
> repository does — is its own business.

> **Every type here is `norm: <type>`.** The prefix scopes the set to work on the norms, so it
> collides with nothing an orchestrator already calls a bug or a task, and one prefix lists all
> of it.

The type is orthogonal to the document's tag: the tag says what the item is about, the type says
how it is worked. An item about one document carries both; an item about the corpus as a whole
carries the type alone. An item with no type is not one of these: it is ordinary work, on the
ordinary route.

| Type | Side | Filed by | Finalized when |
|---|---|---|---|
| `norm: request` | home | anyone, from any repository | an amendment has decided it and its flow has recorded that |
| `norm: amendment` | home | the maintainer, or the first request that finds none open | the pass has landed and every request it read is decided |
| `norm: ratification` | home | anyone, from any repository | the proposal has moved into the root and the move has landed |
| `norm: release` | home | the amendment or ratification that landed a change | the tag, its notes, and the items it files exist |
| `norm: upgrade` | dependent | a release, or the dependent itself | the change bumping the declaration has landed |
| `norm: reconciliation` | any repository | a person, a ratification, or a flow that found drift | the gap items it found are filed |

What a pass cannot decide is a `norm: request` the pass files against its own corpus, carrying
the decision, its candidates, and the evidence the items it read had gathered. It needs no type
of its own, because it takes the same steps as any other request: it waits for a pass, a pass
decides it, and its flow finalizes it. Filing it is what makes a document plus its open items the
end state ([end state](../normative.md#end-state)) — an undecided question left as prose is read
once and built around, while an item is listed by the tag query and somebody owns pushing it.

A **gap item** carries the document's tag and no type. The implementation short of a rule is
ordinary work on the ordinary route, and it closes when the gap is gone
([reconciliation](../normative.md#reconciliation)) — the one thing the types above never do is
close it early.

> **Every change to a project comes from a flow step, or from a person.** A project's state is
> its tree and its items alike — a file written, an item filed, labelled, commented on, or closed
> — and nothing else may change either. A tool that finds work to be done is a tool a step runs,
> never an actor with a schedule of its own; what runs unwatched may only report.

> **A step may file an item in another repository, and may do nothing else there.** Not a
> commit, not a branch, not a review — and making the labels that item needs is part of filing
> it. The item is what starts that repository's own flow, so every change to a tree is the
> product of a flow resolving an item in that same tree.

Dissemination has to cross a repository boundary somewhere, and an item is the narrowest thing
that can cross it: it carries a request and no authority, it is visible to the project before
anything happens there, and the flow that answers it runs under that project's own gates, review
and maintainer. A step that reached across and committed would be doing the project's work
without the project, which is the same defect as a scheduled actor, one repository further out.

A person files an item whenever they have one to file, and that is the act everything else
begins from. After it, every change is attributable: it belongs to a step, of a flow, resolving
a named item, and it is reviewable as that step's product. The automation holds no authority a
contributor lacks. A change nobody's flow produced is a change no review, gate, or transcript
accounts for — and an item that appeared from a schedule is one nobody asked for, arriving in a
tracker where every other item means somebody wants something.

## Steps

> **A flow is a graph of steps, and every step names the step that follows it.** Which step
> follows may depend on what the step found, and a step may name one already run. A step that
> names none is **terminal**: reaching it finalizes the item, and the flow is complete.

> **A flow runs until the item is finalized.** A step that cannot complete — waiting on an
> answer, on a signal, on another item's outcome, on an actor with a capability the present one
> lacks, on infrastructure that is down — **parks**, and the flow resumes at that step when what
> it waits for arrives. Parked is in progress. Nothing else ends a flow.

> **A flow may continue into another's steps.** Where two types owe the same work, the steps are
> written once and the second flow names where it joins. The item keeps its type and finalizes at
> whatever terminal step it reaches.

> **A finalized item is never reopened.** What a finalization missed arrives as a new item, and
> the flow that finalized the old one stays the record of what was decided and why.

> **Every step runs from a person's account, and a step may require a capability of it.** A
> contributor does not review and land their own change; the maintainer does. A step whose
> capability the present actor lacks parks until someone holding it takes the item up, and the
> flow may return to steps the first actor has to complete. It is one flow, on one item,
> whoever is running which step.

Seven kinds, and the kind says what a step may touch:

| Kind | Does | May write |
|---|---|---|
| `read` | gathers what the flow needs and decides nothing | nothing |
| `check` | answers one question about the tree or the items, and its answer picks the next step | nothing |
| `edit` | changes this tree | this tree |
| `land` | runs the gate, commits, and pushes — the only step that does | this tree, the mainline, tags |
| `file` | creates, labels, comments on, or closes items, here or in another repository | items |
| `ask` | puts one typed question to a role and parks until the answer arrives | nothing |
| `converse` | asks and answers in a loop, until the step or the person calls the topic settled | nothing |

> **An `ask` names a role, not a person**: the **maintainer** of the repository the item is in,
> or the **filer** of the item. The roles resolve per repository, so the same flow runs in a
> project whose maintainer it has never met.

`converse` is `ask` without the single round trip: the same park for a person's answer, except
that the answer is not the end of it — the two go back and forth until the topic is settled, and
only then does the flow move on. What that costs is a person's attention for a stretch rather
than once, so it wants a surface built for it, and it is the one kind no flow runs unattended
today. Where a flow reaches one now, a person runs the step and the flow resumes at the step it
names.

## request

A request that the corpus be amended, whatever the reason: a rule that is wrong or unclear, a
rule that is missing, a mechanism proposed, a project's wording offered as the norm, a
dictionary entry. The type says only that something should change; where the filer knows which
document, the tag says which, and the body says why. It carries what
[reconciliation](../normative.md#reconciliation) requires — the rule quoted, what fails under it,
the replacement text or the request for one.

> **A request need not name a document.** A rule nothing states yet has no tag to carry, and
> where the rule belongs — an existing document, or one created for it — is a decision, not
> something the filer owes.

Asking the filer to place it first is asking them to know the corpus before they may complain
about it, which is how a gap goes unreported. What the request must bring is what fails and what
should be true instead; deciding whose subject that is belongs to the pass, which reads every
document anyway ([one home](../normative.md#one-home)).

> **A `norm: request` is triaged, never decided on its own.** Its flow places the item and then
> parks: what the rule should say is decided by the amendment pass, reading it beside every other
> open request, and the flow resumes to record that decision and finalize the item.

Two requests against one document, decided separately, produce two amendments that may contradict
each other, and the third reader finds the contradiction ([intake](release-cycle.md#intake)).
That is why this flow decides nothing of its own: everything before the park makes sure the item
is found, placed, and read with the others.

| # | Step | Kind | Next |
|---|---|---|---|
| 1 | **Place it** — where a document is its home, give it that tag and name the section that decides its concern or the open request that holds it, so the pass reads it beside what is already settled or already asked; where none is, leave it untagged for the pass to place | `file` | 2 |
| 2 | **Find the pass** — is a `norm: amendment` open at this repository, and where more than one is, the one filed first | `check` | 4 where one is, else 3 |
| 3 | **Open one** — file a `norm: amendment` with no subject of its own: it is where the requests open when it runs are read together | `file` | 4 |
| 4 | **Is it urgent** — does the item state that the rule is actively producing wrong work | `check` | 5 where it does, else 6 |
| 5 | **Say so on the pass** — name this item on the amendment as a reason to run early | `file` | 6 |
| 6 | **Block on the pass** — record that this item is blocked by that amendment: what tells a reader when it is read, and what the pass reads to find its subject | `file` | 7 |
| 7 | **Read what it decided** — the amendment's finalization clears the block; take this item's outcome from its record and state it here: the section that now carries the decision, the reason it was rejected, or the request that now holds the question it raised | `file` | terminal |

> **The block is the wait.** Step 6 records it and step 7 cannot run until it clears, which is
> what parking is ([steps](#steps)) — no step polls, and none exists to do the waiting. Urgency
> is said at step 5, before the block, because after it nothing here runs.

> **A carried question is not something to wait for.** Where the pass could not decide this
> item's question, step 7 names the request that now holds it and finalizes anyway. Blocking on
> that one instead would hold this item open behind a question that may take two passes to
> answer, and every item behind it after that.

> **One amendment is open at a time**, and a request blocks on it. A pass reads every request
> blocked on it, so a second pass would be reading the first one's subject while the first is
> deciding it — two amendments that may contradict each other, which is the defect the pass
> exists to prevent, one level up.

> **Where more than one is open, a request blocks on the one filed first.** The rule exists so
> that nothing chooses.

A request does not decide whether a pass runs; it attaches itself to the one that will read it,
and opens that pass only when there is none. So the pass's subject is never declared in advance:
it is whatever is blocked on it when it runs, which is how a request filed by someone who has
never read the corpus still gets read beside everything else about the same document. Urgency is
a note on the pass, not a pass of its own — what it asks for is that this one runs sooner.

A second amendment is how the one-at-a-time rule is relaxed when it has to be: a pass that must
land in two parts, or a change that cannot wait for one already under review, is opened as its
own amendment, and the requests that belong to it are moved onto it deliberately, by whoever
split them. What the rule refuses is a second pass nobody decided to open — which is why a
request never opens one while another is there, and why the first-filed one is where it lands
when the situation has been split on purpose. The order is the one thing about two open
amendments that is not a judgment call, so it is the one thing a flow may act on alone.

## amendment

The pass, run over every open item at once. It amends the corpus and stops there: cutting the
release is a [release](#release) of its own, which step 12 files.

> **An amendment and a release are two items.** The pass ends when its change has landed, and a
> release begins when someone decides the corpus should move; a pass that amends nothing ends
> the same way and files nothing.

One item spanning both would have two endings — landed, or landed and released — and every
reader of a half-finished one would have to work out which. They also sit on different sides of a
decision: the pass is a judgment about what the rules should say, while the release is a judgment
about when every dependent should be asked to move, and the second is not owed the moment the
first lands. Separate items let a release carry two passes, or wait for a third.

| # | Step | Kind | Next |
|---|---|---|---|
| 1 | **Triage** — take every request blocked on this item; give each unlabelled one its document's tag; map any citation by number to the section's slug. A request no document covers stays untagged until step 3 decides where its rule belongs | `file` | 2 |
| 2 | **Read the corpus whole** — every specification and every proposal, before the items against them; the contradictions between documents and the gaps no item names are pass items too | `read` | 3 |
| 3 | **Resolve** — every request to one outcome, with the reason for each, including which document gains a rule that has none, and the questions the pass will carry rather than decide ([the amendment pass](release-cycle.md#the-amendment-pass)) | `converse` | 4 once the table is approved |
| 4 | **Amend** — the documents, in this tree | `edit` | 5 |
| 5 | **Review** — the diff and the table naming the section that now carries each row | `converse` | 6 once approved |
| 6 | **Check coherence** — contradictions, duplicate definitions, gaps the amendments opened, references that no longer resolve, header and index consistency, and whatever the gate checks; fix what has one right answer | `check` | 7 where anything is left undecided, else 8 |
| 7 | **Ask the batch** — every open question at once, each with its evidence and a recommendation | `ask` | 4 |
| 8 | **Land** — the gate green, one commit | `land` | 9 |
| 9 | **Record the outcomes** — on this item, one row per request read: the section that now carries it, the reason it was rejected, or the question it was carried into; file a `norm: request` for each question the pass could not decide, carrying the decision, its candidates, and the evidence; merge and relabel per the table | `file` | 10 |
| 10 | **Reconcile this tree** — file a `norm: reconciliation` here, naming the documents this pass amended, for what this repository itself now owes them | `file` | 11 |
| 11 | **Did anything change** — did the pass amend a document | `check` | 12 where it did, else 13 |
| 12 | **Ask for the release** — file a `norm: release` naming the landed commit | `file` | 13 |
| 13 | **Finish** — record the commit, and the requests the pass left open | `file` | terminal |

> **Every request blocked on this item at step 1 has a row at step 9**, incorporated, rejected,
> or carried into the request that now holds its question. Each of them reads its own row when
> this item finalizes, and finalizes itself; the pass states each outcome once, in one place.

A pass that amends nothing is a legitimate result: it reaches step 13 through step 11, records
why on each row, and cuts no release.

## ratification

A proposal becomes a specification: it binds from the moment it lands in the root, so what this
flow does before the move is establish that it should ([lifecycle](../normative.md#lifecycle)).

> **Anyone may ask for a ratification; only the maintainer may approve one.** The party proposing
> that a text should bind is not the party that decides it binds, and the step that decides names
> the capability it needs.

| # | Step | Kind | Next |
|---|---|---|---|
| 1 | **Read the proposal and its items** — the document, and every question still open under its tag | `read` | 2 |
| 2 | **Is the text settled** — would any open question's answer change what this document says ([end state](../normative.md#end-state)) | `check` | 3 where none would, else 4 |
| 3 | **Approve the binding** — the maintainer decides that this text should be what every reader is measured against | `ask` | 5 |
| 4 | **Name what is unsettled** — each question whose answer would rewrite the document, and park until they are decided | `ask` | 2 |
| 5 | **Move it** — `git mv` into the root, delete the proposal line, move its index entry, and re-root the relative links the move changed; the body is not touched | `edit` | 6 |
| 6 | **Land** — the gate green, one commit | `land` | 7 |
| 7 | **Reconcile this tree** — file one `norm: reconciliation` naming the document, which binds now and has never been measured against this tree | `file` | 8 |
| 8 | **Ask for the release** — file a `norm: release` naming the landed commit, so a dependent can stand on the new specification | `file` | 9 |
| 9 | **Finish** — record the commit and what the document now binds | `file` | terminal |

> **A ratification that rewrites is not a ratification.** Step 5 moves the text and nothing else;
> a document that needs new words first gets them from an amendment, as its own reviewed change,
> and the move follows. Otherwise what was approved at step 3 is not what lands.

## release

| # | Step | Kind | Next |
|---|---|---|---|
| 1 | **Check the precondition** — the pass has landed, and this repository's own `norm: reconciliation` is closed | `check` | 3 where both hold, else 2 |
| 2 | **Say what is missing** — name the unmet precondition to the maintainer | `ask` | 1 |
| 3 | **Write the notes** — per amended document, what changed and what a dependent now looks for in its own tree, and every heading renamed and every document moved or retired ([the two sides](references.md#the-two-sides)) | `edit` | 4 |
| 4 | **Approve the notes** — they are what every dependent's upgrade and reconciliation runs on | `ask` | 5 |
| 5 | **Tag** — annotated, at the landed commit, with the release entry carrying the notes | `land` | 6 |
| 6 | **File the upgrades** — a `norm: upgrade` in every trailing dependent, carrying the notes; each project's own flow adopts the release and then walks its tree against it ([upgrade](#upgrade)) | `file` | 7 |
| 7 | **Finish** — close this item, naming the tag and the items it filed | `file` | terminal |

## upgrade

One home, one release, one reviewed change ([references](references.md#references)).

| # | Step | Kind | Next |
|---|---|---|---|
| 1 | **Read the notes** — the home's, for the delta between the declared release and the new one | `read` | 2 |
| 2 | **Bump the declaration** — the home's row, to the new release | `edit` | 3 |
| 3 | **Rewrite the references** — every one to that home, at the new release; and where the home is the corpus, the vendored members as that release has them ([updating the copy](distribution.md#updating-the-copy)) | `edit` | 4 |
| 4 | **Does every reference carry** — each one resolving at the new release | `check` | 6 where they all do, else 5 |
| 5 | **Ask what became of it** — name each reference that stopped, with what the notes say happened to its target, and the two ways out: repoint it at what now owns the fact, or delete it together with the claim it supports | `ask` | 3 |
| 6 | **Land** — the gate green, the change opened against the mainline, riding the project's own gates and review | `land` | [reconciliation](#reconciliation) 2 |

The bump lands first and the tree catches up second, on the same item: from step 6 the flow
continues at the reconciliation's walk, with the notes read at step 1 as its subject, and
finalizes there. Declaring the new release and meeting it are different work — the first is
mechanical and the second is judgment — but they are not different items, because nobody needs a
second filing to be told that a tree which just adopted a rule should be measured against it.

> **An upgrade carries nothing else.** Not a gap the new rules opened, not a repair noticed on
> the way, and not the open items the new rules left stale: those are items of their own ([keep a
> change to its subject](../engineering-guide.md#keep-a-change-to-its-subject)), and the
> [reconciliation](#reconciliation) is what walks them. A tree that holds the new rules and has
> not met them is the normal state, and the upgrade carries all of its references or none of
> them.

## reconciliation

> **A `norm: reconciliation` re-establishes one invariant: every gap between a normative document
> and this tree is covered by an open item carrying that document's tag**
> ([reconciliation](../normative.md#reconciliation)). It is filed whenever that invariant is in
> doubt, and its flow is what walks the two against each other.

The gap moves from either side, and the item is the same whichever side moved. **The rules
changed** — an amendment or a ratification moved what the tree is measured against. Where that
reaches a dependent as a release, no item is filed for it: the [upgrade](#upgrade) that adopts
the release continues into these steps once its bump has landed, and the notes are the delta.
**The tree or its world changed** — the project grew something a rule already forbade, or a
platform it does not control moved beneath it, and no rule moved at all. **The capture was
incomplete** — nothing moved, and the open items simply never covered what was already true. The
first arrives with notes and a scope; the other two arrive with a suspicion and take their
documents whole. That is why more than one party reaches these steps: an upgrade continues into
them carrying the notes, a person files the item on a hunch or on a rhythm of their own, a
ratification files it for a document that binds from today, and a flow that trips over drift it
must not fix inline files it rather than carrying a passenger.

| # | Step | Kind | Next |
|---|---|---|---|
| 1 | **Take the subject** — the documents the item names where a person or another flow filed it, and otherwise every specification this tree is held to. An [upgrade](#upgrade) joins at step 2 with the release notes it has already read | `read` | 2 |
| 2 | **Walk it against the tree** — each document in the subject against the code, the tools, and this project's own documents, collecting every gap | `check` | 3 |
| 3 | **Walk the open items against it** — an item whose gap the amendment closed is closed, saying which rule went; one quoting a rule that moved or was reworded is repointed at what now carries it; one whose gap still stands is left alone | `file` | 4 |
| 4 | **Which side is wrong** — for each gap, the tree short of the rule, or the rule short of what is intended ([reconciliation](../normative.md#reconciliation)) | `check` | 5 for the tree's; 6 for the rule's |
| 5 | **File the gaps** — as items sized to be resolved: gaps one change would close travel together and carry every tag they answer to, and gaps that would make one change touch unrelated work are separate items | `file` | 7 |
| 6 | **File at the home** — a `norm: request` carrying the rule quoted, what fails under it, and the replacement text or the request for one | `file` | 7 |
| 7 | **Finish** — close this item, naming what it filed and what it closed; the gap items are ordinary work from here. An item that arrived here as an upgrade finalizes as one | `file` | terminal |

> **The open items are half the walk.** A rule that changed leaves items quoting text that no
> longer exists and items whose gap it closed outright, and nothing else in the system ever
> revisits them. An item nobody can check is worse than no item: it reads as remaining work, it
> is cited in reviews, and the first person to act on it does the wrong thing carefully.

> **One reconciliation is open at a time.** A release that finds one already open adds its notes
> to it rather than filing a second — two walks of one tree would each file items for gaps the
> other is filing too, and neither would know.

> **A gap item is sized by the change that would close it, never by the document that named
> it.** One item may close gaps from several documents and carries all their tags; one document's
> gaps may be several items. What is fixed is coverage: no gap without an item
> ([reconciliation](../normative.md#reconciliation)).

Every item costs a whole resolution — a plan, a review, a full run of the gates — so an item
closing one sentence of one rule spends what an item closing a section spends. Splitting by
document multiplies that by the corpus and imposes an order nobody chose: two items that must
edit the same file either wait for each other or collide. Grouping by the change pays the
overhead once. The bound in the other direction is the ordinary one — an item is one subject
([keep a change to its subject](../engineering-guide.md#keep-a-change-to-its-subject)) — so the
item is as large as one subject allows and no larger, which is a judgment the walk is in the best
position to make and the reason this step is not mechanical.

> **A gap the project judges to be a defect in the rule is not worked in the project.** Step 5 is
> the loop closing: the request lands at the home, and the home's next amendment reads it beside
> every other ([reconciliation in the project](release-cycle.md#reconciliation-in-the-project)).

## Open questions

- What a `converse` step is, exactly, once a flow can run one. The shape is an `ask` in a loop
  — round after round until the step or the person calls the topic settled — which makes the
  open part the edges rather than the semantics: what the step is handed to open with, what it
  returns when the loop ends, whether the artifact is the transcript or something drawn from it,
  who may declare the topic settled, and what happens to a loop nobody closes.
