# org — the organization's shared conventions

The normative documents that apply to **every** promise-language project and belong to **none** of
them: the engineering guides, the CLI guide, the documentation conventions, and cross-project
mechanisms such as anchoring. What is specific to one project lives in that project; what is
specific to the BASE model lives in `base`; what is written here binds the whole fleet.

Start at [docs/index.md](docs/index.md) — the map of the corpus and the rules it is written under.

A specification ratified here reaches each project as a provisioned, hash-checked copy in that
project's own tree, so the rules are present in an agent's context at the moment they have to be
followed, and a local edit of a copy fails the commit and points back here.

## Dev tooling

```
./make        # compiles the dev tools into bin/
bin/verify    # the commit gate
```

The build workflow is described in `CLAUDE.md`; the tooling layout follows
[forge](https://github.com/promise-language/forge)'s blueprint.
