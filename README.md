# appsec-review-smoke-go

A deliberately vulnerable **smoke target** for [appsec-review](https://github.com/wsollers/appsec-review).
It exists so that a review of a tiny, fully understood Go program has one correct answer at every
stage. **Do not deploy, run against untrusted input, or reuse any of this code.**

The program is a toy HTTP server: `main.go` is the code-graph fixture program and serves `/`; the
admin console under `/admin` dispatches operator requests on the `op` parameter.

## Build

One Go module, one `package main`, no network, no generated code:

```sh
GOFLAGS=-mod=vendor go build -o bin/ ./...
```

The declared module requirements are vendored with `go mod vendor`; no package of theirs is
imported, so `vendor/` holds only `modules.txt`. The build is expected to succeed in the
appsec-review `go` build image (go 1.23.1) as an unprivileged user with networking disabled.

## History

`tools/make-history.sh` rebuilds this repository's commit graph from `tools/history/steps/` with
fixed authors and dates, so the history is reproducible and can be reviewed as code.

## License

MIT; see `LICENSE`.
