package devenv

import (
	"bytes"
	"html/template"
)

// applyVariables uses GO's templating to apply the variables
func (ev *EnvironmentConfiguration) applyVariables(input string) (string, error) {
	templ, err := template.New("").Parse(input)
	if err != nil {
		return "", err
	}
	b := bytes.NewBuffer(nil)
	vars, err := ev.GetVariables()
	if err != nil {
		return "", err
	}
	if err = templ.Execute(b, vars); err != nil {
		return "", err
	}
	return b.String(), nil
}
