package interactive

import (
	add_repo2 "github.com/sascha-andres/devenv/internal/interactive/cmds/add_repo"
	add_script2 "github.com/sascha-andres/devenv/internal/interactive/cmds/add_script"
	branch2 "github.com/sascha-andres/devenv/internal/interactive/cmds/branch"
	commit2 "github.com/sascha-andres/devenv/internal/interactive/cmds/commit"
	del_repo2 "github.com/sascha-andres/devenv/internal/interactive/cmds/del_repo"
	log2 "github.com/sascha-andres/devenv/internal/interactive/cmds/log"
	pull2 "github.com/sascha-andres/devenv/internal/interactive/cmds/pull"
	push2 "github.com/sascha-andres/devenv/internal/interactive/cmds/push"
	repo_branch2 "github.com/sascha-andres/devenv/internal/interactive/cmds/repo_branch"
	repo_commit2 "github.com/sascha-andres/devenv/internal/interactive/cmds/repo_commit"
	repo_disable2 "github.com/sascha-andres/devenv/internal/interactive/cmds/repo_disable"
	repo_enable2 "github.com/sascha-andres/devenv/internal/interactive/cmds/repo_enable"
	repo_log2 "github.com/sascha-andres/devenv/internal/interactive/cmds/repo_log"
	repo_merge2 "github.com/sascha-andres/devenv/internal/interactive/cmds/repo_merge"
	repo_pin2 "github.com/sascha-andres/devenv/internal/interactive/cmds/repo_pin"
	repo_pull2 "github.com/sascha-andres/devenv/internal/interactive/cmds/repo_pull"
	repo_push2 "github.com/sascha-andres/devenv/internal/interactive/cmds/repo_push"
	repo_shell2 "github.com/sascha-andres/devenv/internal/interactive/cmds/repo_shell"
	repo_status2 "github.com/sascha-andres/devenv/internal/interactive/cmds/repo_status"
	repo_unpin2 "github.com/sascha-andres/devenv/internal/interactive/cmds/repo_unpin"
	scan2 "github.com/sascha-andres/devenv/internal/interactive/cmds/scan"
	shell2 "github.com/sascha-andres/devenv/internal/interactive/cmds/shell"
	status2 "github.com/sascha-andres/devenv/internal/interactive/cmds/status"
)

func init() {
	addCommands()
	addRepositoryCommands()
}

func addRepositoryCommands() {
	repositoryCommands = append(repositoryCommands, repo_disable2.Command{})
	repositoryCommands = append(repositoryCommands, repo_branch2.Command{})
	repositoryCommands = append(repositoryCommands, repo_commit2.Command{})
	repositoryCommands = append(repositoryCommands, repo_enable2.Command{})
	repositoryCommands = append(repositoryCommands, repo_log2.Command{})
	repositoryCommands = append(repositoryCommands, repo_merge2.Command{})
	repositoryCommands = append(repositoryCommands, repo_pin2.Command{})
	repositoryCommands = append(repositoryCommands, repo_pull2.Command{})
	repositoryCommands = append(repositoryCommands, repo_push2.Command{})
	repositoryCommands = append(repositoryCommands, repo_shell2.Command{})
	repositoryCommands = append(repositoryCommands, repo_status2.Command{})
	repositoryCommands = append(repositoryCommands, repo_unpin2.Command{})
}

func addCommands() {
	commands = append(commands, add_repo2.Command{})
	commands = append(commands, add_script2.Command{})
	commands = append(commands, branch2.Command{})
	commands = append(commands, commit2.Command{})
	commands = append(commands, del_repo2.DeleteRepositoryCommand{})
	commands = append(commands, log2.Command{})
	commands = append(commands, pull2.Command{})
	commands = append(commands, push2.Command{})
	commands = append(commands, scan2.Command{})
	commands = append(commands, shell2.Command{})
	commands = append(commands, status2.Command{})
}
