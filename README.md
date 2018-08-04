[![CircleCI](https://circleci.com/gh/jmatsu/deploygate-cli-go/tree/master.svg?style=svg)](https://circleci.com/gh/jmatsu/deploygate-cli-go/tree/master)

# dpg

    dpg - Golang implementation of  DeployGate API Client CLI
    DeployGate API reference is https://docs.deploygate.com/reference#deploygate-api

## Usage

The basic syntax is:

   dpg [global options] command [command options] [arguments...]

Global options:

   --help, -h     show help
   --version, -v  print the version

### COMMANDS


`help, h` option is avaiable for all commands.

     app           Application-based Operation API
     distribution  Application-based Operation API
     help, h       Shows a list of commands or help for one command


## Installation

```
go get github.com/jmatsu/dpg
```