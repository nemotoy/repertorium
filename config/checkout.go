package config

// TODO Validation

// CheckoutConfig ...
type CheckoutConfig struct {
	Access *AccessConfig `yaml:"access"`
	Target *TargetConfig `yaml:"target"`
	Output *OutputConfig `yaml:"output"`
}

// AccessConfig ...
type AccessConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

// TargetConfig ...
type TargetConfig struct {
	Owner  string `yaml:"owner"`
	Branch string `yaml:"branch"`
}

// OutputConfig ...
type OutputConfig struct {
	Path string `yaml:"path"`
}
