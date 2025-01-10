# Qwest Services

## Setup development environment

We are using `Task` as our task runner, to install `Task` run the command below:

```sh
go install github.com/go-task/task/v3/cmd/task@v3.40.1
```
After installing `Task`, you can set up the project with:

```sh
# installs Air and Swag CLI and also sets up Git hooks
task setup
```

### Air
The `setup` command will install `Air` hot reload. You can read more about it at its [GitHub](https://github.com/air-verse/air).

### Git hooks
The `setup` command will set up pre-commit git hooks that:

- Format all Go files
- Run go vet for potential errors
- Run all tests in the changed directories
- Prevent committing unwanted files (.exe, .test, .out, .log, .env)

## Execute tests

You can run all tests using the command below:

```sh
task test -- <TEST_PATH>
```

### Notes:

- The `--` separator is required before passing arguments
- `TEST_PATH` is the path to the test file or package you want to test (e.g. `./pkg/richerror` or `./pkg/...`)

To run tests with coverage profile, run the following:

```sh
task test:cover -- <TEST_PATH>
```

--- 
## Roadmap

- [x] Add logger
- [x] Create Docker image
- [x] Envelope panic recovered errors
- [ ] Add swagger
- [ ] Add Auth service
