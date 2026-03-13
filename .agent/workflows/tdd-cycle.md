---
description: TDD Workflow: Requirements to Bazel Green
---

# TDD Workflow: Requirements to Bazel Green

1. **Scan Requirements:** Locate the first unchecked item in `requirements.md`.
2. **Draft Test:** Create a `*_test.go` file. 
3. **Sync Bazel:** Run `bazel run //:gazelle` so Bazel sees the new test file.
4. **Verify Red:** Run `bazel test //path/to/pkg/...`. Confirm it fails.
5. **Implement:** Write the minimum Go code in the corresponding `.go` file.
6. **Verify Green:** Run `bazel test //path/to/pkg/...`.
7. **Checklist Update:** If Green, mark the requirement as completed in `requirements.md`.
