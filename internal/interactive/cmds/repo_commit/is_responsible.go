package repo_commit

func (c Command) IsResponsible(commandName string) bool {
	return commandName == "commit" || commandName == "ci"
}
