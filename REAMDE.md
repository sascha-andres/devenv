# devenv

A development environment automater

## Commands

### Command line

devenv add <dev-env name>
devenv bash <dev-env name>
devenv shell <dev-env name>
devenv clean <dev-env name>

### In-App shell

repo add <repo-name> <path> <url>
repo delete <repo-name>
repo commit <repo-name> <commit-msg>
repo push <repo-name>
repo pull <repo-name>
repo merge <repo-name> <commit-msg>
repo branch <repo-name> <branch name>
repo close <repo-name> <branch name> <target branch>
commit <commit-msg>
push
pull
branch <repo-name> <branch name>
close <repo-name> <branch name> <target branch>

## Config file format

    ---

    environment_base_path: /home/user/dev/environments
    environment_config_path: /home/user/dev/configs
    
## Env config format

    ---

    name: <dev-env name>
    branch_prefix: <prefix your new branches with>

    repositories:
      - name: repo1
        path: src/repo1
        url: git@gitserver:path-to-git.git