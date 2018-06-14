package config

// TODO Validation

// Config ...
type Config struct {
	Get *GetConfig `yaml:"get"`
	Gen *GenConfig `yaml:"gen"`
}

// GetConfig ...
type GetConfig struct {
	Checkout *CheckoutConfig `yaml:"checkout"`
}

// CheckoutConfig ...
type CheckoutConfig struct {
	Target *TargetConfig `yaml:"target"`
	Output *OutputConfig `yaml:"output"`
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

// GenConfig ...
type GenConfig struct {
	Testcode *TestcodeConfig `yaml:"testcode"`
}

// TestcodeConfig ...
type TestcodeConfig struct {
	Input *InputConfig `yaml:"input"`
}

// InputConfig ...
type InputConfig struct {
	Path string `yaml:"path"`
}
