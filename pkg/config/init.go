package config

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Limit  Limit  `yaml:"limit"`
	Metric Metric `yaml:"metric"`
	Debug  bool   `yaml:"debug"`
}

type Limit struct {
	Fail uint8 `yaml:"fail"`
}

type Metric struct {
	CpuStatCollectDuration int64 `yaml:"cpu_stat_collect_duration"`
}

func Init() Config {
	data, err := os.ReadFile("./configs/global.yaml")
	if err != nil {
		log.Fatalf("Error during reading file: %v\n", err)
	}

	var conf Config
	if err := yaml.Unmarshal(data, &conf); err != nil {
		log.Fatalf("Error during parsing YAML: %v\n", err)
	}

	if err = conf.Valid(); err != nil {
		log.Fatalf("Error during validation: %v\n", err)
	}
	return conf
}

func (conf Config) Valid() error {
	if conf.Limit.Fail == 0 {
		return errors.New("limit -> fail: not set")
	}

	if conf.Metric.CpuStatCollectDuration == 0 {
		return errors.New("metric -> cpu_stat_collect_duration: not set")
	}

	return nil
}
