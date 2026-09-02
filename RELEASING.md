# Releasing

## The short version

Tag `vX.Y.Z` on `main`. There is nothing to bump in the repository — a Go
module's version *is* its git tag.

```bash
scripts/apidiff.sh          # what changed, and what bump it implies
git tag v0.2.0
git push origin v0.2.0
```

## Why not semantic-release

Two reasons, both specific to Go.

**It is npm-shaped.** semantic-release needs Node and a dependency tree in the
release pipeline of a repository that deliberately has none.

**Its major bump would produce a broken release.** In Go, major versions from
v2 change the *module path*: `github.com/basaltic-sh/sdk-go/v2`, in `go.mod`
and in every import line in this repository and in every consumer. A tool that
cuts `v2.0.0` from a `BREAKING CHANGE:` footer produces a tag that nothing can
import, and the failure is silent — `go get` simply keeps resolving v1. Major
version bumps here are a deliberate edit, not an inference.

## What replaces it

The useful half of semantic-release is "decide the bump automatically". Commit
messages are the wrong input for that in this repository: the SDK is generated
from the platform's OpenAPI specifications, so whether a change is breaking is
decided by the *specification diff* — a removed field, a renamed operation — and
whether someone wrote `feat:` or `fix:` is downstream of that, and often wrong.

`scripts/apidiff.sh` answers the question from the API itself. It compares the
exported surface against the last tag and reports the implied bump. CI runs it
on every pull request, so the effect on the public surface is visible in the
review rather than discovered at tag time.

```
Comparing the exported API against v0.1.0.

Incompatible changes:
- ./compute.(*Client).GetInstance: removed
Compatible changes:
- ./compute.(*Client).FetchInstance: added

  The exported API BREAKS.
  Implied bump: MINOR — v0 carries no compatibility promise.
```

It reports; it does not tag. Releasing stays a decision.

## Version policy

**v0.x while the surface settles.** v0 carries no compatibility promise, which
is the honest posture for an SDK whose API has no external consumers yet. A
breaking change is a minor bump — `v0.2.0` — and the release notes say what
broke.

**v1.0.0 is a promise, not a computation.** It says the surface is stable and
that breaking it costs a major version. Cut it when the platform's API is
itself stable, not before, and never from automation.

**v2+ changes the module path.** If it ever comes to that:

1. `go.mod` becomes `module github.com/basaltic-sh/sdk-go/v2`.
2. Every internal import updates.
3. `internal/gen/emit.go` emits the new path in generated files.
4. Only then tag `v2.0.0`.

Prefer almost anything else. An additive change with a deprecation on the old
name costs consumers nothing; a major version costs every consumer an edit.

## What counts as breaking

The SDK is generated, so most of this is decided upstream in the specifications:

| Specification change | Effect |
| --- | --- |
| New operation | Compatible — a new method |
| New optional response field | Compatible — a new struct field |
| New optional request field | Compatible — a new pointer field |
| New required request field | **Breaking** — an added non-pointer field is compatible to compile against but the call now fails at the server |
| Removed or renamed field | **Breaking** |
| Removed or renamed operation | **Breaking** |
| Optional becoming required | **Breaking** — the field's Go type changes from pointer to value |
| New enum value | Compatible — enum types are string-based and open |

The one `apidiff` cannot see is the required-field row: adding a required
field to a request body compiles fine and fails at runtime. Read the
specification diff for those.

## Release steps

1. `make generate SPEC=/path/to/openapi` if the specifications moved, and
   commit the regenerated files.
2. `make test && make vet`.
3. `go test -tags integration .` against a live region, with credentials in
   the environment. The unit tests cannot tell you whether the SDK's idea of
   the API still matches the API.
4. `scripts/apidiff.sh` and read what it says.
5. Tag and push:
   ```bash
   git tag -a v0.2.0 -m "v0.2.0"
   git push origin v0.2.0
   ```
6. Confirm the module is resolvable, which also warms the proxy:
   ```bash
   GOPROXY=proxy.golang.org go list -m github.com/basaltic-sh/sdk-go@v0.2.0
   ```

## A caution about deleting tags

A tag that has been fetched is cached by `proxy.golang.org` and by every
consumer's module cache, and deleting it upstream does not remove it from
either. A bad release is fixed by publishing the next version, never by
retagging.
