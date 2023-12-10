# sih-ipfs

## Pre requisites
- Taskfile (https://taskfile.dev/#/installation)
- Golang >= 1.16 (https://golang.org/doc/install)

- This project uses spf13/cobra for cli. To learn more about it, visit [cobra](https://github.com/spf13/cobra)

## Setup
All the common commands are in the Taskfile.yml. To run any command, just run `task <command>`. For example, to run the tests, run `task test`.
- After installing taskfile, run `task dev-setup` to install all the dependencies.

## Project structure
- `cmd` contains the code that will add the flags and other things
  - `cmd/root.go` contains the root command
  - `cmd/client.go` contains the client command
- `pkg` contains the code that will be used by the cmd to run the project
- `bin` will be created after running `task build` and will contain the binary file

## Running the project
- Run `task run` to start the project in development mode.

## Help
- Run `task --list` to see all the available commands.