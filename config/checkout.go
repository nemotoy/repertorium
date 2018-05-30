package config

// TODO Validation

// Config ...
type Config struct {
	Get *GetConfig `yaml:"get"`
}

// GetConfig ...
type GetConfig struct {
	Checkout *CheckoutConfig `yaml:"checkout"`
}

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
	Owner string `yaml:"owner"`
}

// OutputConfig ...
type OutputConfig struct {
	Path string `yaml:"path"`
}
