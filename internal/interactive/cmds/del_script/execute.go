package del_script

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sascha-andres/devenv"
	"github.com/sascha-andres/devenv/internal/helper"
	"github.com/spf13/viper"
	"log"
	"os"
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
		log.Printf("error: script does not exist in environment folder")
		return errors.New("no such script in environment folder")
	}

	var scriptDirectory = path.Join(viper.GetString("configpath"), e.Name)
	var scriptInScriptDirectory = path.Join(scriptDirectory, scriptName)
	if ok, _ := helper.Exists(scriptInScriptDirectory); !ok {
		log.Printf("error: script does not exist in config directory")
		return errors.New(fmt.Sprintf("no such script in configuration folder %s", scriptDirectory))
	}

	err := os.Remove(scriptInScriptDirectory)
	if err != nil {
		return err
	}

	return os.RemoveAll(path.Join(viper.GetString("basepath"), e.Name, scriptName))
}
