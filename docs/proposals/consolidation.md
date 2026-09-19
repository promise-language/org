# Consolidation

> **Proposal.** Not normative: an end state under discussion, binding nothing until ratified.

> **Tag:** `consolidation` — remaining work to complete this document: the query named in
> `docs/index.md`.

A pass over every open item in one repository. It resolves nothing; it makes the backlog
resolvable.

> **The invariant it re-establishes: the open items are a complete, current, non-overlapping
> statement of the distance between a tree and what its norms call for — one item per area of
> the product, each one current enough to act on today.**

Consolidation is the visible half and the half the name carries: many small items become one
item per area. Two more come with it and neither is optional — every surviving item is brought
up to the current direction, and every gap the norms name that no item covers is filed. A pass
that merged without those two hands back a tidier backlog saying the same wrong things.

Three forces break the invariant, and nothing else in the system repairs them. **Filing is
continuous and per-incident** — a person or a flow step files what it hit, at the size it hit
it, so the backlog grows small, overlapping, and shaped by the order things were noticed rather
than by the surfaces they belong to. **Items age past the direction** — the rule an item quotes
was amended, the subsystem it names was rewritten, the fix it proposes is now wrong; an item
nobody can check still reads as remaining work, and the first person to act on it does the wrong
thing carefully. **Nothing files the gaps nobody tripped over** — a rule a tree has never met is
invisible until someone walks the two against each other.

> **A consolidation is an item of type `item: consolidation`, and not a norm flow.** It carries no
> change to a normative document, so it is none of the `norm:` [types](norm-flows.md#the-types).
> The `item:` prefix scopes the types whose subject is the items themselves, as `norm:` scopes
> those whose subject is a rule, and this document defines `item: consolidation`: the merge, the
> refresh, and the coverage walk, which a person filing one gets together.

## What it is given

- **The repository** — the checkout the pass runs in. Its `git origin` decides the item store.
- **The item store** — [the item store](#the-item-store) names what differs between them.
- **In progress** — what marks an item as being worked right now, in that store.
- **The norms** — the specifications in the `docs/` root ([location](../normative.md#location)),
  and the corpus at the release the project declares ([the declaration](references.md#the-declaration)).
  `docs/proposals/` is direction, not a measure.
- **The areas** — the product's surfaces. The pass derives them and the maintainer confirms them
  at step 13; a wrong list of areas is a wrong pass, so it is the first thing shown.
- **Whole or incremental** — whether the pass honours the marks or ignores them
  ([idempotency](#idempotency)). The first pass in a repository is necessarily whole.
- **A focus** — an item or an area to start from. A focus orders the reading; it does not narrow
  the scope.

## What it changes

> **A consolidation writes items, and nothing else.** No commit, no branch, no fix — not even an
> obvious one-liner found on the way.

A defect the pass discovers is an item like any other ([keep a change to its
subject](../engineering-guide.md#keep-a-change-to-its-subject)): a pass that repaired code on its
way past would be unreviewable as either thing, and its repair would collide permanently with
the same repair landing on its own. The pass may read anything, run the project's own gates and
tests, and run the product, to check what an item claims. It leaves the tree as it found it.

## Scope

Every open item that is not in progress and carries no `norm:` type is **in scope**: it may be
amended, retitled, relabelled, merged, closed, or left alone.

Items in progress, `norm:`-typed items, closed items, and items in other repositories are
**read and never touched**. They are evidence, and they are constraints:

- An **in-progress** item may be about to close a gap another item names. The overlap is noted on
  the other item, and nothing merges into an item being worked.
- A **`norm:` item** says a rule is about to change. The amendment pass owns it, and a
  consolidation that rewrote one would be a second pass reading the amendment's subject while the
  amendment decides it ([request](norm-flows.md#request)). Where an in-scope item depends on a
  rule an open `norm:` item is questioning, the item says so and the question stays where it is.
- A **closed** item is how the pass checks whether a claim was already answered. It is cited by
  number and evidence, never reopened.

## Idempotency

> **A consolidation run twice with nothing moved in between changes nothing.** The second run
> reads the marks, finds no arrival and no trigger, and reports that.

Every item the pass decides on comes out carrying a **mark**: the consolidation that last placed
it. The mark is one value, and what it points at — the mainline commit, the corpus release, and
the time the pass ran — is recorded once on the pass's own item rather than copied onto every
item it touched ([one home](../normative.md#one-home)). Between them they answer the only
question a later pass has about an item: which world was this judged against, and has that world
moved.

An item carries no mark until a pass has placed it, so the mark's presence is the partition:

- **A new arrival** carries none. It gets the whole treatment — checked against the norms and the
  tree, placed in an area, then merged, amended, closed, or kept, like any item in a first pass.
- **A placed item** carries one. Its placement was decided and approved, and a later pass leaves
  it alone unless something moved.

> **A placed item is revisited on a trigger, never on a fresh opinion.** The triggers: a rule it
> quotes was amended since its mark, its area gained an arrival, a claim it makes was answered by
> work that landed, or a person edited it. Absent one, its outcome is `keep`.

That rule is what makes the pass converge rather than oscillate. Two areas a pass deliberately
kept separate carry no record of that decision anywhere else, so a second pass holding an opinion
and no trigger merges them, and a third splits them — churn costing the maintainer an approval
each time and leaving the backlog no better than the first pass left it. The mark is that
missing record: it says the placement was decided, and it puts the burden on the next pass to
name what changed.

**A pass's base is its predecessor's mark.** An incremental run walks the delta — the documents
amended since the base, the tree's changes since it, the items filed since it — rather than the
corpus and the backlog whole, and that is what makes a re-run cost a fraction of a first run.

> **A pass may be run whole, ignoring every mark**, and the first pass in a repository is
> necessarily one. That is how an incomplete walk is repaired, because an incremental pass
> inherits its predecessor's coverage — including the gaps its predecessor missed.

## Steps

The kinds are [norm-flows](norm-flows.md#steps)'s, so this converts to a flow without rewriting.

| # | Step | Kind | Next |
|---|---|---|---|
| 1 | **Read the norms** — every specification in the `docs/` root and the corpus at the declared release, whole, before any item; then the proposals, as direction | `read` | 2 |
| 2 | **Read the direction** — the README's Status section, the project's agent instructions, and the mainline's log since the base, or since the oldest in-scope item on a pass run whole. What changed underneath the backlog is what makes items stale | `read` | 3 |
| 3 | **Read every item** — in scope and out, each with its number, title, body, labels, dates, assignee, linked changes, blocks, the mark it carries, and every reference it makes to another item | `read` | 4 |
| 4 | **Partition** — the arrivals from the placed items, and of the placed, those a trigger reaches ([idempotency](#idempotency)). A pass run whole takes every item as an arrival. Where nothing is in play, the flow goes to 16 | `check` | 16 where nothing is in play, else 5 |
| 5 | **Build the claim ledger** — one line per distinct claim the items in play make: what is wrong or missing, and which items assert it. It is built before the pass has an opinion about the answer, and step 11 ticks it off | `read` | 6 |
| 6 | **Take the areas** — the areas the base pass established, plus what the arrivals need; a pass run whole derives them from the tree, the documents, and the ledger. Every claim lands in exactly one ([areas](#areas)) | `check` | 7 |
| 7 | **Walk the tree** — check every claim in play that can be checked: build it, run the gates, run the product. A claim about the tree is a fact rather than a judgment, and staleness is a claim about the tree | `check` | 8 |
| 8 | **Walk the norms against the tree** — every gap, including those no item names ([gaps nobody filed](#gaps-nobody-filed)); the documents amended since the base, or every one on a pass run whole | `check` | 9 |
| 9 | **Classify** — every item in play to exactly one outcome ([outcomes](#outcomes)) | `check` | 10 |
| 10 | **Draft** — the restructure table, and the full text of every item that survives it ([the proposal](#the-proposal)) | `edit` | 11 |
| 11 | **Review the draft** — against the ledger and against the norms ([the review](#the-review)) | `check` | 12 where anything is undecided or contradictory, else 13 |
| 12 | **Ask the batch** — every open question at once, each with its evidence and a recommendation | `ask` | 10 |
| 13 | **Approve** — the maintainer reads the areas, the table, and the item texts, and says what to change | `converse` | 14 once approved |
| 14 | **Apply** — in the order that keeps the store consistent at every point ([applying](#applying)) | `file` | 15 |
| 15 | **Verify the store** — re-read every open item and diff it against the approved table | `check` | 14 on any mismatch, else 16 |
| 16 | **Record and report** — write what this pass measured against on its own item — the mainline commit and the corpus release read at step 1, and the time — which is what the next pass takes as its base; then what was closed, merged, filed, asked, and left alone, and the count before and after | `file` | terminal |

> **Nothing is written to the store before step 14.** Steps 1 through 13 produce one document and
> one conversation. A pass that had already closed forty items when the maintainer disagreed with
> its areas cannot be taken back.

## Areas

> **One item per area of the product, not one item per change.** `flow: the list output` is an
> item; `fix the column alignment in flow list` is a line inside it. The unit is the surface a
> person names without opening the code — a command and its output, a subsystem, a backend, a
> document's subject, a platform.

An item per change is what a continuously filed backlog already holds, and it is what makes
resolution expensive. The per-item cost [reconciliation](../normative.md#reconciliation) names —
a plan, a review, a full run of the gates — is paid once per item whether it closes a line or a
surface, so twelve items against one surface pay it twelve times, in an order nobody chose, with
two of them colliding in the same file. An area item pays it once, and the person resolving it
holds the whole surface while they work — which is when the second and third defects turn out to
be the same defect.

The bound in the other direction is coherence, not size. **An area is as large as one person can
hold while working it, and no larger.** Where the gaps under one name would make a single change
touch unrelated work — the parser and the installer under `compiler` — that is two areas wearing
one word, and they are two items. Where a defect spans areas, it belongs to the area that would
fix it, and the other item names it as a reference.

> **Inside an item, each defect is a checklist line carrying its own evidence.** That is what
> keeps an area item from becoming a mood: one line per thing that must become true, each with
> the rule quoted or the behaviour observed, each individually checkable, and the item closes
> when they are all ticked.

A line is where a merged item's words survive, and it is how a resolution knows when it is
finished. An item that says only that an area should be better cannot be closed by anything
checkable, which is the defect the checklist exists to prevent at the larger size.

**A title names the area.** `<area>: <what is wrong across it>`, in a form that stays true after
the item is amended and after half its lines are ticked. A title naming one symptom is narrowed
by that symptom, and the next person files a second item for the second symptom of one defect.

## Outcomes

Every item in play takes exactly one, and each carries a reason a reader can check in a minute
([reconciliation](../normative.md#reconciliation)). A placed item no trigger reaches is not in
play and takes `keep` without being read for an outcome at all ([idempotency](#idempotency)).

| Outcome | When | Carries |
|---|---|---|
| **keep** | current, checkable, and already the area's item | nothing changes |
| **amend** | it is the area's item, but the item does not read true — stale quotes, a fix now wrong, an incident title, absorbed lines to add | the rewritten item |
| **merge** | its claims belong to another item's area ([merging](#merging)) | the target, with its claims as lines, and a closing comment naming it |
| **close as fixed** | the tree already meets it | the evidence — a commit, a passing gate, a transcript of the run |
| **close as never true** | the claim does not reproduce and never did | the evidence of the check |
| **close as superseded** | the direction moved past it, or the rule it asks for was decided otherwise | the section or the decision, quoted |
| **close as duplicate** | the same claim, already carried | the number that carries it |
| **refile at the home** | the item is right and the rule is wrong, or the document lives elsewhere | a `norm: request` at the document's home, quoting the rule, what fails under it, and the replacement text or the request for one |
| **new** | an area with gaps and no item, or a gap no item names | a full item |

> **Stale is not an outcome.** An item that looks old is checked, and then it is closed with
> evidence or amended to say what is true now. Closing on a hunch is how a real gap leaves the
> backlog without ever being answered, and an item nobody can check is worse than no item.

> **A gap closes from the side that is wrong** ([reconciliation](../normative.md#reconciliation)).
> The tree short of a rule is an item here. The rule short of what is intended is a `norm:
> request` at the document's home — never work this repository does, and never a rule this pass
> relaxes to fit what was built.

> **A consolidation files requests, never proposals.** A [request](norm-flows.md#request) is the
> defined channel and costs the corpus one row in its next amendment pass. Writing a proposal is
> design work, and design work dressed as triage is a direction nobody chose.

## Merging

> **Merge into the earliest item in the area, always.** The lowest number, the oldest filing
> date.

It holds the history, and it is what other items, commits, and comments already point at.
Choosing by which body reads better makes the choice a judgment nobody can check.

The area's item carries **everything its parts carried**: every label and tag, every block, every
cross-reference, every piece of evidence, and the original words wherever they still decide
something. It takes the **highest priority and urgency** of its parts and names which line earned
it. It ends with the numbers it absorbed, so every old number still leads somewhere.

**A dependency is a reason to merge.** Where one item blocks another, both are in scope, and one
change would close both, they are one item — the dependency is usually the tell that the two are
one surface seen twice.

> **A blocked item never merges into an unblocked one, unless the block is the same work.**

Carrying the block is not enough; it *spreads* the block, and work that could proceed today now
waits behind something else. Where the block crosses out of scope — an in-progress item, a
`norm:` item, an item in another repository — the two stay separate and the block is recorded,
not resolved away. An area whose lines are half blocked is two items: what can move, and what
waits.

## Gaps nobody filed

Step 8 is half the value of the pass and the half nobody asks for. Each normative document is
walked against the tree and what the walk finds is filed, as lines on the area's item or as a new
one. The backlog is what readers noticed; it is not everything there is.

Items that predate the norms, or sit outside them entirely — a bug, a change request, a want —
are **ordinary work and they stay**, as lines on their area's item like any other. Two things are
checked about each: whether a rule now decides it, in which case the line gains that document's
tag and quotes the rule; and whether a rule now contradicts it, in which case it is a question
for the maintainer, never a silent close.

## The proposal

Written to a file and shown. Three parts:

**The areas** — the list, with one line each saying what it covers. It is checked first, because
everything below is shaped by it.

**The table** — every in-scope item, exactly once, in number order: item, title, area, outcome,
one-line reason, evidence. Then every new item. The outcome of the whole backlog is readable
without opening the store.

**The texts** — the full title and body of every item that survives, as it will read after the
pass: what the area is, the checklist of what must become true with the rule quoted or the
behaviour observed on each line, the blocks, the tags, and the numbers it absorbed.

## The review

Before the proposal is shown, it is reviewed against the draft as adversarially as a reviewer
reads someone else's.

1. **Nothing lost.** Every line of the claim ledger is ticked: carried by a named item at a named
   checklist line, or dropped with evidence. A claim dropped without evidence is a defect in the
   pass, not a decision.
2. **No contradiction.** Every surviving item is read against the normative documents. A proposed
   item that contradicts a rule is either a misreading or a defect in the rule, and both go to the
   maintainer. **There should be none**, and a pass about to ship one has not finished this step.
3. **Coverage the other way.** Every gap step 8 found has a line.
4. **The areas hold.** No item whose lines would make one change touch unrelated work, no defect
   on two items, no area with neither an item nor a reason.
5. **The store's own invariants.** Blocks carried, tags carried, no surviving item referencing one
   the pass closes, no merge into an in-progress item, no merge that spreads a block.
6. **The marks.** Every item still open that the pass decided on carries this pass's mark, and no
   placed item was revisited without a named trigger. A revisit the pass cannot justify is an
   opinion, and an opinion is what the mark exists to keep out.
7. **The count.** Fewer, larger items is the point; a count that barely moved is said in the
   report rather than passed over — and on an incremental pass it is the expected result.

## Applying

The order matters, because the store is readable by other people at every moment of it.

1. Create the new items, so closing comments can point at live numbers.
2. Amend the survivors: title, body, labels, priority.
3. Re-point references: any surviving item whose body points at an item the pass closes now points
   at the target.
4. Record the blocks on the targets.
5. Close the absorbed items last, each with a comment naming the item and the line that now
   carries it, or the reason and the evidence.
6. Mark every item still open that this pass decided on, arrivals and placed alike, and mark them
   last. A pass that failed halfway leaves its unmarked items to be taken as arrivals by the next
   one, which is the direction that costs a re-read rather than a missed judgment.
7. No item is deleted, and no closed item is edited.

## The item store

The pass is identical on both; only these phrases differ.

| | GitHub Issues | Tracker |
|---|---|---|
| Read all | `gh issue list --state open` | the tracker's list call |
| In progress | an assignee, an open linked pull request, or the flow's in-progress label — where it cannot be told, the item is treated as in progress | `in_progress`, or leased to an arena |
| Type | a label, as `norm: request` is | the item's own type field |
| Blocks | stated in the body and by reference | the tracker's own field |
| Amend | `gh issue edit` | the tracker's update call |
| Close | `gh issue close` with a comment | update to done or wontfix, with a note |
| Ask | the flow's question sentinel in the final message, carrying the evidence and the recommendation | the tracker's question call |
| Mark | a `Consolidated: <pass>` line in the body's footer, beside the numbers the item absorbed | a field of the item's own, or a note in the one form |
| Edited since | the item's update time against the base's time | the same |

Where a person is running the pass interactively and watching it, the question is put to them
directly and no sentinel is needed.
