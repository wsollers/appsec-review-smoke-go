# Security notes

This repository is a deliberately vulnerable smoke target. It is never deployed, so there is no
disclosure process. The entries below record the issues that were fixed in this history, so that
the record of what changed and why is in one place.

## 2026-04-16: path traversal in `readAsset` (CWE-22)

The `asset` op joined the operator-supplied name onto the asset root and read the result, so a
name containing `../` read files outside the asset directory. `readAsset` now cleans the joined
path with `filepath.Clean` and refuses any result that does not start with the asset root.

## 2026-03-30: command injection in `runTool` (CWE-78)

The `tool` op built a shell command line from the operator-supplied tool name and ran it through
`/bin/sh -c`, so shell metacharacters in the name ran arbitrary commands. `runTool` now accepts
only the names `status` and `version` and performs them directly; there is no command string.

## 2026-03-01: `github.com/gorilla/websocket` 1.4.0 (CVE-2020-27813, GO-2020-0019)

gorilla/websocket before 1.4.1 could overflow the frame length while reading a frame and let a
peer exhaust the server's memory. The requirement was moved to 1.5.3, the current release.
