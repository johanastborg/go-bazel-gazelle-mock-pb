# Phonebook API Project

This is a sample project for a Phonebook API built with Go, Protocol Buffers, and managed with the Bazel build system.

## Project Structure

The project has a modular structure utilizing Go modules and Bazel packages:

- `api/phonebook/v1`: Contains the Protobuf definitions (`person.proto`) and generated Go code tests.
- `go.mod`: Go module definition for `asymptotic.world`.
- `MODULE.bazel`: Bazel module file declaring dependencies on `rules_go`, `gazelle`, `protobuf`, and `rules_proto`.

## Prerequisites

To build and test this project, you need:

- **Bazel**: A fast, scalable, multi-language and extensible build system. (Version 8.0+ is recommended as defined by the dependencies).
- **Go**: The Go programming language (managed by rules_go in Bazel).

## Building

To build all the targets in the project, run:

```bash
bazel build //...
```

## Testing

To run all the tests in the project, run:

```bash
bazel test //...
```

This will compile the protobuf definitions, build the Go code, and execute the tests (e.g., `TestCreatePerson` in `person_test.go`).

## Generating Build Files with Gazelle

If you add new Go files or update dependencies, you can use Gazelle to automatically generate or update your `BUILD.bazel` files.

To update the build files, run:

```bash
bazel run //:gazelle
```

## Dependencies

- [rules_go](https://github.com/bazelbuild/rules_go): Bazel rules for building Go.
- [bazel-gazelle](https://github.com/bazelbuild/bazel-gazelle): Build file generator for Bazel projects, specifically for Go.
- [protobuf](https://github.com/protocolbuffers/protobuf): Protocol Buffers compiler and runtime.
- [rules_proto](https://github.com/bazelbuild/rules_proto): Bazel rules for building Protocol Buffers.
