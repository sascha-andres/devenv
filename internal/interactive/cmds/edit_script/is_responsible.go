package edit_script

func (c Command) IsResponsible(commandName string) bool {
	return commandName == "editscript"
}
