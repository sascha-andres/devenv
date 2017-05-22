package devenv

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/spf13/viper"

	yaml "gopkg.in/yaml.v1"
)

var (
	shPath   string
	shExists bool
)

// EnvironmentConfiguration contains information aout the project
type (
	EnvironmentConfiguration struct {
		Name         string                    `yaml:"name"`
		BranchPrefix string                    `yaml:"branch-prefix"`
		Repositories []RepositoryConfiguration `yaml:"repositories"`
		Environment  map[string]string         `yaml:"env"`
		Shell        string                    `yaml:"shell"`
	}
	environ []string
)

// LoadFromFile takes a YAML file and unmarhals its data
func (ev *EnvironmentConfiguration) LoadFromFile(path string) error {
	reader, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error opening project config: %#v\n", err)
	}
	file, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatalf("Error loading project config: %#v\n", err)
	}
	if err := yaml.Unmarshal(file, ev); err != nil {
		log.Fatalf("Error reading project config: %#v\n", err)
	}
	return nil
}

// StartShell executes configured shell or default shell (sh)
func (ev *EnvironmentConfiguration) StartShell() error {
	command := exec.Command("bash", "-l")
	env := environ(os.Environ())
	for key := range ev.Environment {
		env.unset(key)
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

// ProjectIsCreated checks whether project is checked out
func ProjectIsCreated(projectName string) bool {
	if ok, err := exists(path.Join(viper.GetString("basepath"), projectName)); ok && err == nil {
		return true
	}
	return false
}

// unset removes an environment key for a new process
func (e *environ) unset(key string) {
	for i := range *e {
		if strings.HasPrefix((*e)[i], key+"=") {
			(*e)[i] = (*e)[len(*e)-1]
			*e = (*e)[:len(*e)-1]
			break
		}
	}
}

// exists returns whether the given file or directory exists or not
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func init() {
	var err error
	shPath, err = exec.LookPath("sh")
	if err != nil {
		return
	}
	shExists = true
}
