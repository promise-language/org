module org/tools/build

go 1.26

// The shared helpers every managed project's tooling needs — hash, staleness,
// platform, exec, args, hook wiring, the verified-tree record — live once in
// forge/primitives rather than as a copy here (forge docs/primitives.md, One
// implementation). The dependency is pinned at an exact version and verified by
// go.sum, so an upstream change reaches this project when this line is raised
// and never before (The dependency is pinned).
require github.com/promise-language/forge v0.0.0-20260916195433-2511337e312b
