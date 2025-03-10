package dto

type Global struct {
	Limits Limits `yaml:"limits"`
}

type Limits struct {
	FailLimit int `yaml:"fail_limit"`
}
