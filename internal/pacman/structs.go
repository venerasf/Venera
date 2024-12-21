package pacman

type Target struct {
	Script      string  `yaml:"script"`
	Description string  `yaml:"description"`
	Version     float64 `yaml:"version"`
	Hash        string  `yaml:"hash"`
	Path        string  `yaml:"path"`
	Tags        []string `yaml:"tags"`
}

type Pack struct {
	Author      string  `yaml:"author"`
	Description string  `yaml:"description"`
	Version     float64 `yaml:"version"`
	Target      []Target `yaml:"target"`
}