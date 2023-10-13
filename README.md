# lw-cli (liquidweb-cli)
Official command line interface for the LiquidWeb API
```
CLI interface for LiquidWeb.

Command line interface for interacting with LiquidWeb services via
LiquidWeb's Public API.

If this is your first time running, you will need to setup at least
one auth context. An auth context contains authentication data for
accessing your LiquidWeb account. As such one auth context represents
one LiquidWeb account. You can have multiple auth contexts defined.

To setup your first auth context, you run 'auth init'. For further
information on auth contexts, be sure to checkout 'help auth' for a
list of capabilities.

As always, consult the various subcommands for specific features and
capabilities.

Usage:
  lw-cli [command]

Available Commands:
  asset         All things assets
  auth          authentication actions
  cloud         Interact with LiquidWeb's Cloud platform
  completion    Generate completion script
  dedicated     All things dedicated server
  default-flags Manage default flags
  help          Help about any command
  network       network actions
  plan          Process YAML plan file
  ssh           SSH to a Server
  version       show build information

Flags:
      --config string        config file (default is $HOME/.liquidweb-cli.yaml)
  -h, --help                 help for lw-cli`
      --use-context string   forces current context, without persisting the context change

Use "lw-cli [command] --help" for more information about a command.
```
## Obtaining prebuilt binaries

Head on over to the [releases page](https://github.com/liquidweb/liquidweb-cli/releases)  to get prebuilt binaries for your platform.

## Building from source

You can build lw-cli from source by running `make build` from the root of this repository. The resulting program will be located at `./_exe/lw-cli`.
You can also build+install lw-cli onto your system in the ordinary `go install` way. To do this, either just run `go install` from the root of this repository,
or `make install`. If you run `make` with no arguments, this will be the default action.

## First Time Setup
The first time you use lw-cli, you will need to setup an auth context. An auth context holds authentication related data for a specific LiquidWeb account. You can follow a guided questionnaire to add your auth contexts if you pass arguments `auth init` to lw-cli. By default contexts are stored in `~/.liquidweb-cli.yaml` or `%APPDATA%/.liquidweb-cli.yaml` on Windows.

## Adding auth contexts later
If you end up wanting to add an auth context later on, you can do so with `auth add-context`. You can find the usage documentation in `help auth add-context`.

## Removing auth contexts later
If you end up wanting to remove an auth context later on, you can do so with `auth remove-context`. You can find the usage documentation in `help auth remove-context`.

## Modifying auth contexts later
If you end up wanting to modify an auth context later on, you can do so with `auth update-context`. You can find the usage documentation in `help auth update-context`.

## LiquidWeb Cloud
The Cloud features you can use in manage.liquidweb.com on your Cloud Servers you can do with this command line tool. See `help cloud` for a full list of features and capabilities.

## Plans

A plan is a pre-defined yaml with optional template variables that can be used to
repeate specific tasks.  Fields in the yaml file match the params you would send
to the command.

For current commands supported via plans, take a look in the `examples/plans` directory of this repo.

Example:

`lw-cli plan --file plan.yaml`

```
---
cloud:
   server:
      create:
         - type: "SS.VPS"
           password: "{{- generatePassword 25 -}}"
           template: "UBUNTU_1804_UNMANAGED"
           zone: 27
           hostname: "web1.somehost.org"
           ips: 1
           public-ssh-key: "your public ssh key here
           config-id: 88
           bandwidth: "SS.5000"
```

You can find more examples of plans in `examples/plans`.

### Plan Variables

Plan yaml can make use of golang's template variables.  Allows variables to be passed on the
command line and it can access environment variables.

#### Environment Variables
Envonrment variables are defined as `.Env.VARNAME`.  On most linux systems and shells you can
get the logged in user with `{{ .Env.USER }}`.

#### User Defined Variables
If you wanted to pass user defined variables on the command line you would use the `--var` flag
(multiple `--var` flags can be passed).  For example, if you wanted to generate the hostname of
`web3.somehost.org` you would use the following command and yaml:

`lw-cli plan --file play.yaml --var node=3 --var role=web`

```
    hostname: "{{- .Var.role -}}{{- .Var.node -}}.somehost.org"
```


#### Functions

The following functions are defined to be called from a plan template:

- generatePassword <length>

For example, to generate a random 25 character password: 

```password: "{{- generatePassword 25 -}}"```

- now

Gives access to a Golang `time` object using your local machine's clock.

Simple example:

```
    hostname: "web1.{{- now.Year -}}{{- now.Month -}}{{- now.Day -}}.somehost.org"
```

- hex

Convert a number to hexidecimal.

