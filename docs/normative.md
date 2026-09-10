# Normative Documents

> **Tag:** `normative` — remaining work to complete this document: the query named in
> `docs/index.md`.

> **Home:** [promise-language/org](https://github.com/promise-language/org) — this document is
> distributed into each managed project as `docs/org/`. A copy is never edited in place: to
> change it, file an issue against `org`.

What makes a document in a managed project binding, what one must contain, and the rules that
keep two of them from ever disagreeing. Every managed project holds the same documentation
structure, so a reader who has learned one tree has learned them all. This document's subject is
the documents themselves; it is one of them, and every rule below applies to it.

## 1. Location is the whole rule

There is no configuration file, and no marker decides what a file is. The directory a file sits
in determines it, and the header line (§2) only restates what the directory already decided:

| Location | What a file there is | Binding? |
|----------|----------------------|----------|
| `docs/` root | A **specification**: what the project *should* be — the intended end state. | **Yes.** Work that contradicts one stops and is resolved — the document amended, the item adjusted, or the item rejected — never shipped as a quiet deviation. |
| `docs/org/` | The **organization-wide corpus**, vendored from its home repository. | **Yes.** Never edited in place: it changes at its home, and reaches the project by sync. |
| `docs/proposals/` | An end state that has **not been ratified** — a draft, an RFC, a direction under discussion. | No. |
| `docs/archive/` | An end state that has been **superseded or delivered** — kept for history. | No. |
| `docs/research/` | Background analysis feeding a decision — an assessment, not a design. | No. |

Those five rows are the whole vocabulary. A project may lack `research/` or have an empty
`archive/`; it may not invent a sixth location or assign one of these a different meaning. The
corpus's home repository is the one tree with no `docs/org/`: its root *is* the corpus, and it
is held to every rule here exactly as a project is.

**`docs/index.md` is the map and the one file in the root that is not a specification.** Every
tracked file under `docs/` is listed in it, wherever it lives — the section an entry sits under
is where its binding status is written down. `docs/org/` is listed once, as the directory, by
way of the **stamp** it carries: the stamp names the release the copies came from and every
member, so the corpus's map is the corpus's own, and a sync adds and removes members without
editing a file the project owns. The index also names the project's status query (§2), which is
the one per-project fact this shared document cannot carry.

## 2. The header

Every Markdown document under `docs/` — the index excepted — opens with its title and, on the
line beneath it, the header line its location requires. The location decides; the line restates
the decision for the reader who arrived by a link rather than by the directory, and it is
checked against the location (§8) so the two cannot drift. One form per location — the root,
`proposals/`, `archive/`, `research/`, in that order:

```markdown
> **Tag:** `<basename>` — remaining work to complete this document: the query named in `docs/index.md`.
> **Proposal.** Not normative: an end state under discussion, binding nothing until ratified.
> **Archived.** Not normative: superseded or delivered, kept for history.
> **Research.** Not normative: an assessment feeding a decision, not a design.
```

The tag is always the file's basename minus `.md`, so the vocabulary is the directory listing
and nothing duplicates it. The query is the project's own — a `gh issue list --label`
invocation on a GitHub-tracked project — and its exact spelling is stated once, in
`docs/index.md`, which is why the tag line points there instead of carrying it: a query repeated
in every header is a copy in every header. Files under `proposals/`, `archive/`, and `research/`
carry no tag: none of them is something the project owes work against. A document under
`docs/org/` carries the header it has at its home — the tag line — byte for byte, like the rest
of it.

A document that is distributed into other repositories also carries, after its tag line, its
**home line**: the repository it changes in, and how to ask for the change. A copy's reader is
the reader that line exists for (§7); its exact wording belongs to the home repository's
distribution rules.

## 3. A specification states the end state

**A specification describes what the project should be, never how far along it is.** No status
sections, no progress notes, no phasing, no "currently", "not yet", or "implemented", and no
inline markers naming an item — a status section arriving one sentence at a time.

The practical test: **a specification reads identically the day before and the day after the
work that implements it.** A sentence that would have to change when an item closes is status,
and does not belong.

Where status lives instead, and all three are queries or single homes rather than prose in a
specification:

- **Per document**: the open items carrying the document's tag *are* its status section, always
  current.
- **For the project**: the README's Status section is the one sanctioned home for "how far along
  is this" prose. Nothing else records progress.
- **Undecided design**: a root document may carry an **Open questions** section, for questions
  whose answers would change what the specification says. Undecided *time* is not an open
  question — work merely not done yet is an item carrying the document's tag.

**A rule stated as a blockquote is an invariant**, and the prose under it is why. If the rule
and the reasoning ever disagree, the rule is what the implementation must satisfy.

## 4. One fact, one home — supersession is forbidden

**A fact is specified in exactly one document.** Two specifications must never define the same
thing, and no specification may claim authority over another. *Supersedes*, *takes precedence
over*, *overrides*, *this document wins*, *the authoritative version is* — a document that needs
one of these is proof that a fact has two homes, and the remedy is always to delete the
duplicate and cross-reference, never to rank the copies.

Split a shared subject by giving each document a different **kind** of statement: the *model* (what
a thing means, what invariant it preserves) in the document that owns the concept; the *contract*
(its surface, parameters, interactions) in the reference for that surface. Neither restates the
other, and each links to the other once.

**A fact whose home is source code stays there.** Prose says where to look, not what it will
say. And a fact whose home is the org corpus stays there: a project document cites `docs/org/`,
it does not restate it.

## 5. Cross-reference, do not copy

Link to the document that owns a fact. If a passage must be edited whenever its target changes,
it is a copy however it is worded — a paraphrase and a quotation drift identically. A copy is
sanctioned only where a machine checks it: the vendored `docs/org/`, byte-identical, verified
against its stamp, and named as a copy by its own header; and the header line (§2), which
restates what the directory decided and is checked against it.

## 6. Lifecycle

Three transitions, each one reviewed change:

- **Ratification.** A design begins in `docs/proposals/`, unbound, untagged, freely rewritten —
  and **written as the specification it would become**. §3's end-state voice, §4's one home per
  fact, and §5's cross-references all apply to it, because what makes it a proposal is where it
  sits, not how it is written. Ratifying it is one act: create the label, `git mv` into the
  root, replace the proposal line with the tag line, move its index entry — the move *is* the
  decision. The test mirrors §3's: **a proposal reads identically the day before and the day
  after it is ratified**, but for its location, its header line, and its index entry. A proposal
  that would need rewriting first is not ready, and a rewrite folded into the ratification is a
  change nobody can diff against what was proposed.
- **Amendment.** An ordinary reviewed diff, landing **before or with** the change that
  implements it, never after: a specification trailing its implementation has stopped describing
  the end state and started reporting history.
- **Retirement.** `git mv` into `docs/archive/`, replace the tag line with the archive line, and
  move its index entry. Content stays; location — and so authority — changes. The label outlives
  the document: deleting it would strip the closed items that record the work.

  **Completion never retires a specification.** A document whose implementation is finished has
  an empty tag query — its healthiest state, not its end. It stays in the root, where it keeps
  the next change from quietly undoing the work and gives every future reconciliation pass its
  measure; retiring it on delivery would turn "implemented" back into "unspecified", the §7 gap
  it took the work to close. A specification retires only when it stops describing the intended
  end state: a ratified replacement supersedes it, its subject is removed from the project, or
  the direction is abandoned — each a decision about the design, never a report that work
  finished. The one document that *is* delivered is one that was inherently one-shot — a staged
  migration, a bounded sequence — and it archives when it completes; that is the "delivered" in
  §1's archive row, and it is the exception, not the pattern.

## 7. Reconciliation

> **Every gap between a specification and the implementation is covered by an open item carrying
> that document's tag.**

That invariant is what makes the tag query a *complete* status section. Three kinds of gap —
unbuilt, divergent, unspecified — are recorded the same way, and a gap in the *document* is
still a gap: a rule that is wrong, unclear, or missing is covered by an item under the same tag
as a rule that is unmet. After a ratification or amendment, walk the document against the
implementation and file the items that close every gap, as its own change.

**A gap closes from the side that is wrong, never by moving the other side to meet it.** A
compliance gap — the implementation short of the rule — closes by changing the implementation,
and the document stays as written. A definition gap — the rule short of what is intended —
closes by amending the document (§6). Closing the item records that the gap is gone; an item may
not be closed while its gap remains. Relaxing a rule to fit what was built, or building to a rule
known to be wrong, is §1's quiet deviation spelled from the other end.

**An item carries what a reader needs to check it in a minute**: the rule — the section, quoted,
not paraphrased — what fails under it, with evidence the reader can confirm, and what would
close it: the change to the implementation, or the replacement text, or the statement that the
item is asking for one. An item that says only that a section should be clearer cannot be
closed by anything checkable.

> **An issue about a document is filed where the document originates.** A defect in a rule, a
> change request, an amendment proposal — these go to the repository the document's home line
> names, never to a project holding a copy.

Documents live in many repositories; their meaning must not. What a project files locally under
a document's tag is only its **own** gaps against it — the compliance work its tree owes. So a
shared document's tag names two different queries in two places, and both are complete: in the
home repository, the remaining work on the *definition*, together with the home's own compliance
gaps — one tag, because both are remaining work on that document, and a second tag would split
one status into two queries neither of which is whole; in each project, that project's remaining
work toward it. An issue filed in the wrong place is invisible to the query that should have
listed it — the same defect as a fact with two homes. The vendored copies make the wrong place
the *natural* place: a reader meets the document in the project's tree and files where they
stand. Such an issue is transferred to the home repository, not worked where it landed — and the
home line is what tells both the reader and the transferrer where that is.

When the corpus is amended, this same pass runs in every project against the delta: the home
repository's release is what starts it, and each project's tag queries are where it lands.

## 8. What is enforced mechanically

- Every relative link in every tracked Markdown file resolves to a tracked file — and **a
  failing link is repaired by fixing the reference, never by removing the link.** Repoint it at
  the document that now owns the fact, or delete the reference together with the claim it
  supports. A link turned into plain text leaves the reference exactly as wrong as it was and
  makes the check report a pass: a repaired reference and a silenced one must never look alike.
- Every tracked file under `docs/` is listed in `docs/index.md` (§1) — and **a missing index is
  a failure, not a pass**: nothing to report and nothing was checked must never look alike.
- The line under every title is the one its location requires (§2), and a tag names the file
  it sits in.
- `docs/org/` is refused at the edit by the guard and verified against its stamp by the
  integration gate.

Everything else here is upheld by review, and the gaps against this document are items carrying
its tag.
