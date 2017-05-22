package devenv

type EnvironmentConfiguration struct {
	Name         string                    `yaml:"name"`
	BranchPrefix string                    `yaml:"name"`
	Repositories []RepositoryConfiguration `yaml:"repositories"`
}
