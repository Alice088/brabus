package dto

type Config struct {
	Limits Limits `yaml:"limit"`
	Debug  bool   `yaml:"debug"`
}

type Limits struct {
	FailLimit uint8 `yaml:"fail_limit"`
}
