package types

type Field struct {
	Name         string              `yaml:"name"`
	Type         string              `yaml:"type"`
	Min          int                 `yaml:"min,omitempty"`
	Max          int                 `yaml:"max,omitempty"`
	Unique       bool                `yaml:"unique,omitempty"`
	Pattern      string              `yaml:"pattern,omitempty"`
	Format       string              `yaml:"format,omitempty"`
	Prefix       string              `yaml:"prefix,omitempty"`
	Suffix       string              `yaml:"suffix,omitempty"`
	Distribution Distribution        `yaml:"distribution,omitempty"`
	Options      []string            `yaml:"options,omitempty"`
	Dependencies map[string][]string `yaml:"dependencies,omitempty"`
	Default      interface{}         `yaml:"default,omitempty"`
}

type Distribution struct {
	Type   string  `yaml:"type"`
	Mean   float64 `yaml:"mean,omitempty"`
	StdDev float64 `yaml:"std_dev,omitempty"`
	Min    float64 `yaml:"min,omitempty"`
	Max    float64 `yaml:"max,omitempty"`
}

type Config struct {
	Fields []Field `yaml:"fields"`
}
