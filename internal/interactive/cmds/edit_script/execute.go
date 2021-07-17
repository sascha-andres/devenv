package edit_script

import (
	"github.com/pkg/errors"
	"github.com/sascha-andres/devenv"
	"github.com/sascha-andres/devenv/internal/helper"
	"github.com/sascha-andres/devenv/internal/os_helper"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
)

func (c Command) Execute(e *devenv.EnvironmentConfiguration, executeScriptDirectory, repositoryName string, args []string) error {
	_ = executeScriptDirectory

	if !(runtime.GOOS == "linux" || runtime.GOOS == "darwin") {
		log.Printf("editscript is not supported on %s", runtime.GOOS)
		return nil
	}
	scriptName := strings.Join(args, "")
	if ok, _ := helper.Exists(path.Join(viper.GetString("basepath"), e.Name, scriptName)); !ok {
		log.Printf("error: script does not exist")
		return errors.New("no such script")
	}

	environmentVariables := os_helper.GetEnvironmentVariables()
	value, found := environmentVariables["EDITOR"]
	if !found {
		log.Println("no EDITOR variable found")
		return errors.New("no EDITOR variable found")
	}
	command := exec.Command(value, []string{path.Join(viper.GetString("basepath"), e.Name, scriptName)}...)
	env := helper.BuildEnvironment(environmentVariables)
	command.Dir = path.Join(viper.GetString("basepath"), e.Name)
	command.Env = env
	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr
	_, err := os_helper.StartAndWait(command)
	return err
}
