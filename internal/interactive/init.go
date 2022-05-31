package interactive

import (
	"github.com/sascha-andres/devenv/internal/interactive/cmds/add_repo"
	"github.com/sascha-andres/devenv/internal/interactive/cmds/add_script"
	"github.com/sascha-andres/devenv/internal/interactive/cmds/branch"
	"github.com/sascha-andres/devenv/internal/interactive/cmds/commit"
	"github.com/sascha-andres/devenv/internal/interactive/cmds/del_repo"
	"github.com/sascha-andres/devenv/internal/interactive/cmds/del_script"
	"github.com/sascha-andres/devenv/internal/interactive/cmds/edit_script"
	"github.com/sascha-andres/devenv/internal/interactive/cmds/fetch"
	"github.com/sascha-andres/devenv/internal/interactive/cmds/log"
	"github.com/sascha-andres/devenv/internal/interactive/cmds/pull"
	"github.com/sascha-andres/devenv/internal/interactive/cmds/push"
	"github.com/sascha-andres/devenv/internal/interactive/cmds/repo_branch"
	"github.com/sascha-andres/devenv/internal/interactive/cmds/repo_commit"
	"github.com/sascha-andres/devenv/internal/interactive/cmds/repo_disable"
	"github.com/sascha-andres/devenv/internal/interactive/cmds/repo_enable"
	"github.com/sascha-andres/devenv/internal/interactive/cmds/repo_fetch"
	"github.com/sascha-andres/devenv/internal/interactive/cmds/repo_log"
	"github.com/sascha-andres/devenv/internal/interactive/cmds/repo_merge"
	"github.com/sascha-andres/devenv/internal/interactive/cmds/repo_pin"
	"github.com/sascha-andres/devenv/internal/interactive/cmds/repo_pull"
	"github.com/sascha-andres/devenv/internal/interactive/cmds/repo_push"
	"github.com/sascha-andres/devenv/internal/interactive/cmds/repo_shell"
	"github.com/sascha-andres/devenv/internal/interactive/cmds/repo_status"
	"github.com/sascha-andres/devenv/internal/interactive/cmds/repo_unpin"
	"github.com/sascha-andres/devenv/internal/interactive/cmds/scan"
	"github.com/sascha-andres/devenv/internal/interactive/cmds/shell"
	"github.com/sascha-andres/devenv/internal/interactive/cmds/status"
)

func init() {
	addCommands()
	addRepositoryCommands()
}

func addRepositoryCommands() {
	repositoryCommands = append(repositoryCommands, repo_disable.Command{})
	repositoryCommands = append(repositoryCommands, repo_fetch.Command{})
	repositoryCommands = append(repositoryCommands, repo_branch.Command{})
	repositoryCommands = append(repositoryCommands, repo_commit.Command{})
	repositoryCommands = append(repositoryCommands, repo_enable.Command{})
	repositoryCommands = append(repositoryCommands, repo_log.Command{})
	repositoryCommands = append(repositoryCommands, repo_merge.Command{})
	repositoryCommands = append(repositoryCommands, repo_pin.Command{})
	repositoryCommands = append(repositoryCommands, repo_pull.Command{})
	repositoryCommands = append(repositoryCommands, repo_push.Command{})
	repositoryCommands = append(repositoryCommands, repo_shell.Command{})
	repositoryCommands = append(repositoryCommands, repo_status.Command{})
	repositoryCommands = append(repositoryCommands, repo_unpin.Command{})
}

func addCommands() {
	commands = append(commands, add_repo.Command{})
	commands = append(commands, add_script.Command{})
	commands = append(commands, edit_script.Command{})
	commands = append(commands, del_script.Command{})
	commands = append(commands, fetch.Command{})
	commands = append(commands, branch.Command{})
	commands = append(commands, commit.Command{})
	commands = append(commands, del_repo.Command{})
	commands = append(commands, log.Command{})
	commands = append(commands, pull.Command{})
	commands = append(commands, push.Command{})
	commands = append(commands, scan.Command{})
	commands = append(commands, shell.Command{})
	commands = append(commands, status.Command{})
}
