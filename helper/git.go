package helper

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/sascha-andres/devenv"
	"github.com/spf13/viper"
)

var (
	gitExecutable string
)

// Git calls the system git in the project diorectory with specified arguments
func Git(ev devenv.EnvironmentConfiguration, args ...string) error {
	command := exec.Command("git", args...)
	env := Environ(os.Environ())
	for key := range ev.Environment {
		env.Unset(key)
	}
	for key, value := range ev.Environment {
		log.Printf("Setting '%s' to '%s'", key, value)
		env = append(env, fmt.Sprintf("%s=%s", key, value))
	}
	command.Dir = path.Join(viper.GetString("basepath"), ev.Name)
	command.Env = env
	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr
	if err := command.Start(); err != nil {
		return fmt.Errorf("Error running bash: %#v", err)
	}
	if err := command.Wait(); err != nil {
		return fmt.Errorf("Error waiting for bash: %#v", err)
	}
	return nil
}

func init() {
	var err error
	gitExecutable, err = exec.LookPath("git")
	if err != nil {
		log.Fatalf("Could not locate git: '%#v'", err)
	}
}
