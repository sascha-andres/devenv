package devenv

// EnvironmentConfiguration contains information aout the project
type EnvironmentConfiguration struct {
	Name         string                    `yaml:"name"`
	BranchPrefix string                    `yaml:"branch-prefix"`
	Repositories []RepositoryConfiguration `yaml:"repositories"`
}
