# devenv

devenv is a tool to manage projects containing multiple git repositories. With a config
you can setup all repositories with one command and even work with the local repositories
as if they where one ( eg creating branches in all referenced projects) using the shell.

[![Go Report Card](https://goreportcard.com/badge/github.com/sascha-andres/devenv)](https://goreportcard.com/report/github.com/sascha-andres/devenv) [![codebeat badge](https://codebeat.co/badges/8cba9986-0b0b-4142-8cbf-03c0e7c4f5d6)](https://codebeat.co/projects/github-com-sascha-andres-devenv-master)

## Prerequisites

You have to have a working git commandline installation.

## Known caveats

Developed and mostly tested under Linux. OSX and Windows are currently not thoroughly tested.

Commands to be executed before a bash is called is unsupported on Windows.

A shell must be configured for Windows.

## Main configuration

There are two main configuration options:

1. basepath
2. configpath

Those values can either be set in your home directory in a .devenv.[yaml,toml,json] file,
in the current directory or as environment variables.

A sample yaml file looks like this:

    ---

    basepath: /dev/src
    configpath: /home/user/.devenv

This project uses [viper](https://github.com/spf13/viper) for configuration and commandline
argument handling.

### basepath

This is the path in which projects are created. E setup project creates a folder within
and checks out the connected repositories relative to that directory.

### configpath

Where to store the project specific configurations.

## Project configuration

The project configuration contains references to the repositories and basic settings for the environment:

1. name
2. repositories
3. env
4. shell
5. commands

A valid project configuration looks like this:

    ---

    name: devenv
    repositories:
      - name: devenv
        url: git@github.com:sascha-andres/devenv.git
        path: src/devenv
    processes:
      shell: bash
      commands:
        - echo Hello
      env:
        VAR: VALUE

### name

Name must be a unique name for your project. It is used to identify the project configuration in
devenv calls.

### repositories

A list of repositories. Each repository must provide the following information:

1. name
2. path
3. url

A valid repository entry looks like this:

    name: devenv
    path: src/devenv
    url: git@github.com:sascha-andres/devenv.git

#### name

Is used to identify the repository for operations specific to one repository

#### path

A relative path where the repository is cloned to

#### url

Remote url to the repository

### processes.env

Key value pairs that are added to the shell and git processes

### processes.shell

Executable fot the shell. If you want to use the `fish` shell in one project as opposed to your
default shell specify here.

### processes.shellarguments

Passed to the shell as arguments

### processes.commands

Commands to execute before the shell is called.

## Variables

Each environment variable known to the devenv process can be access as go standard template in the following places:

* commands
* shell
* shell-arguments
* Environment variables for processes

Values in env override the environment variables. An additional variable called ENV_DIRECTORY is provided containing the path to the environment.

__While technically possible to inject references this is not supported!__

## Commands

Commands are top level commands to work with devenv itself. The following commands are supported:

* add
* shell
* clean
* setup

### Interactive shell

**Breaking change in 1.3.0**

Call `devenv -- <environment>` to start the interactive shell. With `devenv shell <environment>` the bash shell or configured shell will be called.

For commands see `In-App shell`

### add

Create a new project stub:

    devenv add project

This will create a new configuration in your configuration directory.

### shell

Open a system shell in the project directory with environment variables preconfigured.

### clean

Remove project from basepath. Checks for uncommited changes.

### setup

Create project directory and clone all referenced projects

## In-App shell

The in app shell provides easy methods to work with your git repositories. Non aliases are part of autocompletion

### addrepo command

|Info||
|---|---|
|Aliases|none|
|Description|Will ask for the required values of a repository and if cloning is successful saves the new repository to the project configuration|

Call with 

    addrepo

Required information:
* name
* path
* url

### branch

|Info||
|---|---|
|Aliases|br|
|Description|Will switch to the specified branch for each repository|
||__If the branch does not exist, it will be created__|

Call with

    branch <name>

You can add any number of additional arguments as long as they are valid arguments
to `git checkout`

Required information:
* name

### delrepo command

|Info||
|---|---|
|Aliases|none|
|Description|Will ask for the name of a repository and remove it from project when there are no changes|

Call with 

    delrepo <name>

Required information:
* name

### log command

|Info||
|---|---|
|Aliases|l|
|Description|Prints out the last ten commits decorated for each repository|

Call with 

    log

### pull command

|Info||
|---|---|
|Aliases|<|
|Description|Executes a pull for each repository|

Call with 

    pull
  
You can add any number of additional arguments as long as they are valid arguments
to `git pull`

### push command

|Info||
|---|---|
|Aliases|>|
|Description|Executes a push for each repository|

Call with 

    push
  
You can add any number of additional arguments as long as they are valid arguments
to `git push`

### status command

|Info||
|---|---|
|Aliases|st|
|Description|Prints the status of each referenced repository|

Call with 

    status
  
You can add any number of additional arguments as long as they are valid arguments
to `git status`

### Repository commands

Repository commands take the name of a repository as they work on a single repository and not on every referenced repository.

They are prefixed with `repo` which can be shortened to `r`.

The following commands are available for a single repository:

* branch
* commit
* log
* merge
* pull
* push
* status

#### branch

|Info||
|---|---|
|Aliases|br|
|Description|Will switch to the specified branch for each repository|
||__If the branch does not exist, it will be created__|

Call with

    repo <name> branch <name-of-branch>

You can add any number of additional arguments as long as they are valid arguments
to `git checkout`

Required information:
* name
* name-of-branch

#### log command

|Info||
|---|---|
|Aliases|l|
|Description|Prints out the log for the repository|

Call with 

    repo <name> log

Required information:
* name

#### merge command

|Info||
|---|---|
|Aliases|none|
|Description|Merge a branch|

Call with 

    repo <name> merge <branch-name>

You can add any number of additional arguments as long as they are valid arguments
to `git merge`

Required information:
* name

#### pull command

|Info||
|---|---|
|Aliases|<|
|Description|Executes a pull|

Call with 

    repo <name> pull
  
You can add any number of additional arguments as long as they are valid arguments
to `git pull`

Required information:
* name

#### push command

|Info||
|---|---|
|Aliases|>|
|Description|Executes a push for repository|

Call with 

    repo <name> push
  
You can add any number of additional arguments as long as they are valid arguments
to `git push`

Required information:
* name

#### status command

|Info||
|---|---|
|Aliases|st|
|Description|Prints the status of repository|

Call with 

    repo <name> status
  
You can add any number of additional arguments as long as they are valid arguments
to `git status`

Required information:
* name

## Code of conduct

You can find the code of conduct [here](conde_of_conduct.md), taken from [Contributor Covenant](http://contributor-covenant.org/)

## History

|Version|Description|
|---|---|
|v1.2.0|Make log parameters configurable|
| |Refactorings|
|v1.1.1|Fix for branch detection|
| |Refactorings|
|v1.1.0|Completion for shell|
| |Command arguments for shell|
| |Variables|
|v1.0.1|Enable/Disable repository|
|v1.0.0|Initial version|
