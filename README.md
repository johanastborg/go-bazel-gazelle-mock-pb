# Phonebook API

This project implements a Phonebook API using Protocol Buffers (protobuf) and Go. It uses [Bazel](https://bazel.build/) as its build and test system, along with `rules_go`, `gazelle`, and `protobuf` dependencies.

## Project Structure

- `api/phonebook/v1/`: Contains the Protocol Buffer definitions (`.proto` files) and associated Go tests.
- `MODULE.bazel`, `BUILD.bazel`, `go.mod`: Bazel and Go module configuration files.

## Building the Project

To build the entire project, run:

```bash
bazel build //...
```

If you add new dependencies, files, or packages, remember to run Gazelle to update the `BUILD.bazel` files:

```bash
bazel run //:gazelle
```

## Using Test-Driven Development (TDD) with Bazel

Test-Driven Development (TDD) relies on a short, iterative development cycle where you write tests before writing the implementation.

In this project, you can use Bazel to run your tests efficiently. Bazel caches test results and only re-runs tests if the underlying code or dependencies have changed.

### Running Tests

To run all tests in the project, use the following command:

```bash
bazel test //...
```

To run a specific test target (for example, the person tests):

```bash
bazel test //api/phonebook/v1:v1_test
```

### Continuous Testing for TDD

For a smoother TDD workflow, it is highly recommended to use [ibazel](https://github.com/bazelbuild/bazel-watcher) (a tool that watches your source files and automatically re-runs Bazel commands when they change) or similar file watchers.

With `ibazel`, you can continuously run tests as you save files:

```bash
ibazel test //...
```

Alternatively, you can use a simple file watcher like `entr` or `watch` combined with the standard Bazel command if you prefer not to install `ibazel`:

```bash
# Example using entr
find . -name "*.go" -o -name "*.proto" | entr -c bazel test //...
```

### Typical TDD Workflow

1.  **Write a Failing Test:** Add a new test case in a `_test.go` file (e.g., in `api/phonebook/v1/person_test.go`).
2.  **Run the Test:** Run `bazel test //...` (or use `ibazel`) and watch it fail. This confirms your test is valid and actually testing the absence of the feature.
3.  **Write the Implementation:** Add the minimal code necessary (e.g., update the `.proto` file or add Go code) to make the test pass.
4.  **Run the Test Again:** Run `bazel test //...` and see it pass.
5.  **Refactor:** Clean up your code, confident that the tests will catch any regressions.
