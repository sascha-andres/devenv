package devenv

type RepositoryConfiguration struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
	URL string `yaml:"url"`
}