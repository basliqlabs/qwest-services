# Qwest Services

## Run this project

We are using `Task` as our task runner, to install `Task` run the command below:

```sh
npm install -g @go-task/cli
```

> If you don't have `npm` installed on your machine, please
> follow [Task installation guides](https://taskfile.dev/installation/).

> This project uses [Air live reload](https://github.com/air-verse/air), so please run the following command to install it:

```sh
go install github.com/air-verse/air@latest
```

With `Task` and `Air` installed, you can run the command below to jump start the project:

```sh
task dev
```

## Setup development environment

After cloning the repository, install the git hooks:

```sh 
task setup-hooks
```

This will set up pre-commit hooks that:

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

- [ ] Add logger
- [ ] Add swagger
- [ ] Create Docker image
- [ ] Add Auth service
- [ ] Envelope panic recovered errors