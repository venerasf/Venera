package pacman

type Pack struct {
	Author      string  `yaml:"author"`
	Description string  `yaml:"description"`
	Version     float64 `yaml:"version"`
	Target      []struct {
		Script      string  `yaml:"script"`
		Description string  `yaml:"description"`
		Version     float64 `yaml:"version"`
		Hash        string  `yaml:"hash"`
		Path        string  `yaml:"path"`
		Tags        []string `yaml:"tags"`
	} `yaml:"target"`
}