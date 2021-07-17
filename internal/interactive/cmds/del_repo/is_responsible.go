package del_repo

func (c DeleteRepositoryCommand) IsResponsible(commandName string) bool {
	return commandName == "delrepo"
}
