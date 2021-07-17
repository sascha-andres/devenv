package del_repo

func (c Command) IsResponsible(commandName string) bool {
	return commandName == "delrepo"
}
