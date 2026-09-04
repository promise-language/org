# Anchoring

> **Home:** [promise-language/org](https://github.com/promise-language/org) — this document is
> distributed into each managed project as `docs/org/`. A copy is never edited in place: to
> change it, file an issue against `org`.

**Proposal. Not normative.** The org-level successor to flow's
[`docs/proposals/anchoring.md`](https://github.com/promise-language/flow/blob/main/docs/proposals/anchoring.md)
sketch, widened to the whole fleet. Nothing described here exists yet, anywhere.

Some parts of a project are not ordinary content. Changing them changes what "correct" means for
every change that follows — and an agent resolving an item can change them as easily as it changes
a comment. An **anchored** aspect is one whose modification requires a person to approve it,
specifically and at the time.

The motivating incidents are in flow's sketch: two resolutions amended the normative document they
were being judged against, invisibly, inside large diffs. Neither was self-serving — that is the
point. Nothing distinguishes a resolution that improves a specification from one that relaxes it to
fit what it built, except reading the reasoning. Anchoring makes sure someone reads it.

## What may be anchored

Aspects, not files — the unit is whatever a person would want to be told about, in any managed
project:

| Aspect | Example | What the anchor covers |
|---|---|---|
| A whole document | `docs/resolution.md` | The file's content |
| A section within one | `## Gates` and its body | The section, wherever it moves |
| A code definition | a type, function, interface — Promise or Go | The definition |
| The public surface of a definition | `flow.Backend`'s exported shape | Signature, exported fields, doc — **not** the body |
| A specific value | a cap, a baseline, a timeout chosen deliberately | The value |

An anchor is therefore **not purely text**. A public-surface anchor binds to the parsed
declaration: renaming a local, editing the body, or reformatting does not touch it; changing a
signature, an exported field, or the documented contract does.

## Identity

> **Every anchor has an identity that survives the anchored aspect moving, being renamed, or being
> reformatted.**

An anchor that dies when its section is retitled protects nothing — the first formatting pass
silently unanchors the tree. Identity is two parts:

- **An address**: `(project, kind, stable path)` — a file path for a document, a heading path for a
  section, a symbol path for a definition. The address is how humans and tools name the anchor.
- **A fingerprint**: a hash of the aspect's *normalized* content — for a public-surface anchor,
  the normalized public surface only. The fingerprint is what enforcement compares; the address is
  how it finds what to fingerprint.

## Declaration

The anchor set is declared in the tree it protects, and **the declaration is itself anchored** —
otherwise the first move available to an agent is to unanchor what it wants to change. The
recursion terminates because the approver is a person.

## Enforcement, at three points

**A guard, at the edit.** The agent proposing to modify an anchored aspect is refused before it
happens, with the reason and the way to ask. Cheap and immediate; the agent adapts mid-turn. It
fails open — a route the guard does not cover goes straight past it.

**The precommit, at the commit.** The shared precommit check set fingerprints every anchored
aspect **before and after the staged change**; a difference without a recorded approval refuses
the commit. Local, mechanical, and covers every route to the tree — including a human editor.

**A gate, at integration.** Compares the anchored aspects against **the state the resolution
started from** — not the mainline, which may have legitimately moved — and refuses the change if
any differ without a recorded approval. The authoritative check, whatever happened in between.

None substitutes for the others: the guard is bypassable, the precommit sees one commit at a time,
and the gate costs a whole resolution before it speaks.

## The ask

A request to change an anchor is typed, and carries:

- **what would change** — the diff against the anchored aspect, alone, not embedded in the rest of
  the work;
- **why** — the reasoning, which is the thing actually being reviewed;
- **what it unblocks** — the change that cannot land without it.

Presenting the anchor diff separately is most of the value: "should this document say something
different" is a different question from "is this implementation right", and answering it inside a
fifteen-hundred-line diff means answering it badly.

## Approval flows through the feed

> **A person approves; nothing else does.** The party proposing a change to an anchor is exactly
> the party that must not be able to authorise it.

The request travels one path: the resolving flow emits it, the orchestrator — tracker today,
reactor eventually — routes it, and it lands in the **engagement feed of the person who owns the
decision**. The feed is central to the platform and **per user**: every project talks to one feed
per person, so an approval looks the same whichever project it came from, and a person's pending
decisions are one list rather than eight. **grid hosts the feed** — a web server where each person
signs in with their GitHub account and sees their own feed; the GitHub account *is* the identity,
the same one that holds repo roles and authors commits, so a recorded grant names the same
principal everywhere. The web interface is the first surface; others may follow, but the feed
behind them is one. The answer — approved or rejected, with reasoning — travels the same path back
to the waiting resolution.

A standalone resolution with no orchestrator uses the backend's question mechanism (the same seam
as `answer`): the request is posted against the item, and the resolution waits.

**A grant is recorded with what it permitted**, so a later reader sees that a person decided and
what they decided about. Scope: per-aspect, per-session — approving a change to a document lets
that session keep editing it — and the **final** state of the aspect is presented again before the
change lands, so what a person approved is what ships.

## Where this is specified

Anchoring is one mechanism crossing several owners, so it lands as **one shared guide plus a
normative document per implementing project**, each owning its own surface:

| Document | Owns |
|---|---|
| The org guide (this document's successor) | The concept, the aspect kinds, identity, grant semantics — how any project uses anchoring |
| flow | Emitting the ask, waiting, the guard refusal, the integration gate |
| The shared tool contract | The precommit fingerprint check in the shared check set |
| tracker | Routing requests and answers, recording grants |
| grid | The feed surface: how an ask reaches a person, is presented, answered, and delivered back |

The feed's *semantics* — the article, its actions, ranking, questions and deadlines — are defined
in reactor's `engagement-feed.md`; grid is the feed's first *implementation*, the surface a person
actually signs into. The split is the usual one: reactor owns what the feed means, grid owns one
way of serving it.

## What this is not

**Not immutability** — an anchored aspect changes deliberately, not never; a mechanism that made
normative documents unchangeable would make them wrong instead. **Not a substitute for review** —
it decides what must be looked at, not whether it is good. **Not a file permission system** — the
unit is an aspect a person cares about; anchoring everything produces a mechanism people turn off.

## Open questions

- The normalization rules per aspect kind — what exactly a public-surface fingerprint includes,
  per language.
- Whether an agent may propose additions to the anchor set (adding is safe in a way removing is
  not).
- How a cross-project ask names its approver — the project's owner, the aspect's owner, or the
  item's requester.
- How the vendored org docs interact with anchors: their protection is the hash check against the
  source repo; anchoring their *source* in the org repo is where a person's approval attaches.
