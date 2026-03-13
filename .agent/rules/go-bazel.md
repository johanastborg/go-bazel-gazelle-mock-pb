---
trigger: always_on
---

# Go & Bazel Engineering Rules

- **Build System:** Use `bazel` exclusively. Never run `go build` or `go test` directly.
- **Dependency Management:** - When `go.mod` changes, run `bazel run //:get_deps`.
    - Always follow up with `bazel run //:gazelle` to update `BUILD.bazel` files.
- **TDD Requirement Parsing:**
    - Read requirements from `requirements.md`.
    - Map each requirement to a specific `go_test` target.
- **Code Style:** All Go code must be formatted via `gofmt` before being committed to the repo.