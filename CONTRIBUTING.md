# Contributing to org

**org** is part of the **Promise Lang** project, hosted in the `promise-language` organization
and maintained under Promise Lang LLC.

This repository holds the organization-wide normative documents. A document ratified here is
vendored into **every managed project's public tree** — so although this repository may be
private, everything in it is written, licensed, and CLA-covered as if it were public, because
its contents float into repositories that are.

## Contributor License Agreement (CLA) required

Before any pull request can be merged, you must sign the **Promise Lang Contributor License
Agreement**. When you open your first pull request, the CLA Assistant bot will post a link to
sign. You only need to sign once — it covers all future contributions across the project.

- **Individual contributors** sign the Individual CLA.
- **Contributors acting on behalf of an employer** also have their employer sign the
  Corporate CLA.

You retain copyright in your contribution; the CLA grants Promise Lang LLC the rights it
needs to administer, distribute, and sublicense it as part of the project.

## Licensing of contributions

Unless you state otherwise, any contribution you intentionally submit for inclusion is
dual-licensed under the [Apache License 2.0](LICENSE-APACHE) and the [MIT
License](LICENSE-MIT), with no additional terms or conditions. Contributions must **not**
introduce material under a copyleft license (GPL, LGPL, AGPL, EUPL, or similar), text of
uncertain provenance, or excerpts whose license does not permit this dual licensing.

## How to contribute

1. Open an issue first. A change to a document here changes what "correct" means in every
   managed project, so the discussion belongs ahead of the diff.
2. Read [docs/index.md](docs/index.md) — the corpus's own conventions govern changes to it:
   location decides binding, specifications carry no status, one fact has one home.
3. A gap between a document and reality is an issue carrying the document's tag — here for
   defects in the rule, in the affected project for defects in its compliance.
4. Open a pull request and sign the CLA when prompted.

## Commit identity

Commits must carry a `@users.noreply.github.com` author and committer email. The repo's
pre-commit hook enforces this locally. Activate the tooling in a fresh clone with `./make`,
which also wires the hooks.
