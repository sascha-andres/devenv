package del_script

func (c Command) IsResponsible(commandName string) bool {
	return commandName == "delscript"
}
