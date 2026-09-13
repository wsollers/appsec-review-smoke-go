# Changelog

All notable changes to this project are documented in this file. The format follows
[Keep a Changelog](https://keepachangelog.com/en/1.1.0/).

## [Unreleased]

### Added

- Admin console under `/admin` with ping, tool and asset ops.
- Pinned module requirements for the websocket transport and the planned command-line client.

### Changed

- Builds offline with `GOFLAGS=-mod=vendor`.
- `tools/make-history.sh` reproduces the commit graph.

### Security

- 2026-03-01: Upgraded `github.com/gorilla/websocket` from 1.4.0 to 1.5.3 (CVE-2020-27813, GO-2020-0019: integer overflow in the frame reader lets a peer exhaust memory).
- 2026-03-30: Fixed command injection in `runTool` (CWE-78); tool names are now matched against a fixed list and nothing reaches a shell.
- 2026-04-16: Fixed path traversal in `readAsset` (CWE-22); the joined path is cleaned and confined to the asset root.
