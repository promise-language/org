# org — the organization's shared conventions

The normative documents that apply to **every** promise-language project and belong to **none** of
them: the engineering guides, the CLI guide, the documentation conventions, and cross-project
mechanisms such as anchoring. What is specific to one project lives in that project; what is
specific to the BASE model lives in `base`; what is written here binds the whole fleet.

Start at [docs/index.md](docs/index.md), the map of the corpus;
[docs/normative.md](docs/normative.md) is the rules it is written under.

A project stands on a release of this corpus by reference: it declares the release in its own
`docs/index.md`, and an agent reads each rule by following the reference at that release
([docs/references.md](docs/references.md)). A rule changes here and nowhere
else — by an issue, resolved in the next amendment pass and released to the fleet
([docs/release-cycle.md](docs/release-cycle.md)).

## Dev tooling

```
./make            # compiles the project tools into bin/
bin/verify        # the commit gate: format → vet → build → test → record
bin/run <gate>    # measure one gate and judge it (bin/gate --list names them)
```

The build workflow is described in `CLAUDE.md`; the tooling layout follows
[forge](https://github.com/promise-language/forge)'s blueprint, and the tool set
is the one the workspace's tool contract fixes — `workspace doctor` checks it.
