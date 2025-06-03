# appsec-review-smoke-go

A deliberately vulnerable **smoke target** for [appsec-review](https://github.com/wsollers/appsec-review).
It exists so that a review of a tiny, fully understood Go program has one correct answer at every
stage. **Do not deploy, run against untrusted input, or reuse any of this code.**

The program is a toy HTTP server: `main.go` is the code-graph fixture program and serves `/`; the
admin console under `/admin` dispatches operator requests on the `op` parameter.

## License

MIT; see `LICENSE`.
