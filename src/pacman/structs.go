package pacman

type Pack struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Author      string   `yaml:"author"`
	Version     float64  `yaml:"version"`
	Date        string   `yaml:"date"`
	Target      []Target `yaml:"target"`
}
type Target struct {
	Script  string  `yaml:"script"`
	Version float64 `yaml:"version"`
	Hash    string  `yaml:"hash"`
	Path    string  `yaml:"path"`
}
