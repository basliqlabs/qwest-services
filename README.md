# Qwest Services

## Run this project

This will install `Task` task runner.

```sh
npm install -g @go-task/cli
```

With `Task` installed, you can run the command below:

```sh
task dev
```

## Executing tests

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
