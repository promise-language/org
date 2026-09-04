# Documentation Index

This is the map of `docs/`. It is the one file in the root that is not a specification —
everything else there is.

## How to read this tree

| Location | What a file there is | Binding? |
|----------|----------------------|----------|
| `docs/` root | A **specification**: what should be — the intended end state. It never records current state, progress, or phasing. | **Yes.** Work that contradicts a root doc must stop and be resolved — amend the doc, adjust the item, or reject it — not shipped as a quiet deviation. |
| `docs/proposals/` | An end state that has **not been ratified** — a draft, an RFC, a direction still under discussion. | No. Ratifying one means `git mv` into the root and giving it a tag. |
| `docs/archive/` | An end state that has been **superseded or delivered** — kept for history. | No. |
| `docs/research/` | Background analysis feeding a decision — an assessment, not a design. | No. |

This repository is the home of the organization-wide corpus: a specification ratified here binds
**every managed project**, and each project carries a provisioned, hash-checked copy so the rules
are in an agent's context at the moment they have to be followed. A copy is never edited in place
— a rule changes here.

**Where progress lives.** A root doc has no status section. Each one declares a **tag** on the
line under its title, and the gap between the end state and today is the set of open issues
carrying that label:

> **Tag:** `example` — remaining work to complete this document: `gh issue list --label example --state open --limit 200`

One tag per root document, spelled exactly as the file's basename minus `.md`. The enumeration is
the directory — `ls docs/*.md` — and this file deliberately does not copy it. Files under
`proposals/`, `archive/`, and `research/` take no tag. A project-local gap against a shared
document is filed in that project under the same label; a gap in the document itself is filed
here.

**Ratifying a proposal is four steps**: `gh label create <tag>`; `git mv
docs/proposals/<tag>.md docs/<tag>.md`; add the tag line under the title; move its entry in this
file into Specifications. The last three are one commit.

**Conventions in a root doc.** A rule stated as a blockquote is an invariant, and the prose under
it is why. An **Open questions** section is for undecided *design* only — work that is merely not
done yet is an issue carrying the document's label.

---

## Specifications

**None yet.** Nothing has been ratified: the corpus below is drafted and under discussion, and a
document promoted before that discussion ends would bind eight projects to a shape nobody agreed
to.

## Proposals — not binding

- [proposals/engineering-guide.md](proposals/engineering-guide.md) — How code in this
  organization is written, in any language.
- [proposals/engineering-guide-promise.md](proposals/engineering-guide-promise.md) — The
  engineering guide applied to Promise source.
- [proposals/engineering-guide-go.md](proposals/engineering-guide-go.md) — The engineering guide
  applied to Go source.
- [proposals/cli-guide.md](proposals/cli-guide.md) — How every command-line tool behaves at its
  invocation surface.
- [proposals/anchoring.md](proposals/anchoring.md) — Aspects whose modification requires a
  person's approval, and the path an approval travels.
- [proposals/distribution.md](proposals/distribution.md) — How a ratified document reaches every
  managed project, and what keeps the copies honest.
