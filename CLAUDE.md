## Dev tooling

Two sets of tools land in `bin/`, which is gitignored, and each name has exactly
one builder (workspace `docs/tool-contract.md`):

- **Project tools** — `make`, `verify`, `gate`, `run`, `setup` — are compiled from
  the in-repo Go module `tools/build/` by `./make`. Always built locally, never
  committed.
- **Workspace tools** — `tool-guard`, `precommit-guard`, `issue`, `workspace` — are
  installed by `workspace setup` from the workspace release. This repository
  builds no twin of any of them.

**Fresh clone — bootstrap once:**

```bash
./make            # Windows: .\make.cmd
```

`./make` compiles every project tool into `bin/` and wires git's `core.hooksPath`
to `.githooks/`. It is idempotent and finishes in well under a second once built.
`workspace setup` (from the workspace release) installs the workspace tools into
the same `bin/`; without it the pre-commit hook fails closed.

**Before every commit — run the gate:**

```bash
bin/verify        # format → vet → build → test → record, then a pass/FAIL summary
```

A green "OK to Commit" line means it is safe to commit; a red FAIL means it is
not. `verify` records the tree it blessed at `.workspace/verified-tree`, and the
pre-commit hook (`bin/precommit-guard`) refuses a commit of any other tree — so
every commit is one `verify` has seen. The hook also blocks staged binaries,
enforces GitHub noreply commit identities, and keeps `docs/index.md` and
`tools/gates/baselines.json` consistent.

**Measure one thing:**

```bash
bin/gate --list           # the gates this project answers
bin/run <gate>            # measure one gate and judge it, term by term
bin/run integration       # everything that must hold before a change may land
```

`bin/gate` measures and modifies nothing; `bin/run` holds the terms. The
thresholds are a distinct artefact in `tools/gates/thresholds.json`, and a change
may not author its own. `bin/workspace doctor` reports whether this checkout
satisfies the tool contract, and it must be green.

**Edit a tool, then rebuild.** If any tool prints
`tools source has changed — run: ./make`, re-run `./make` to rebuild `bin/`.
That is the whole loop: edit → `./make` → use. Add a tool by dropping a new dir
under `tools/build/cmd/` and re-running `./make` — no registration step.
