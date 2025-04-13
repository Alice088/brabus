package dto

type Global struct {
	Limits Limits `yaml:"limits"`
	Debug  bool   `yaml:"debug"`
}

type Limits struct {
	FailLimit int `yaml:"fail_limit"`
}
