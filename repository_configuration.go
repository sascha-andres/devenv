package devenv

// RepositoryConfiguration contains information about linked repositories
type RepositoryConfiguration struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
