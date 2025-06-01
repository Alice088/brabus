package yaml

import (
	"brabus/pkg/dto"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func UnmarshalGlobalConfig() dto.Config {
	data, err := os.ReadFile("./configs/global.yaml")
	if err != nil {
		log.Fatalf("Error during reading file: %v\n", err)
	}

	var conf dto.Config
	if err := yaml.Unmarshal(data, &conf); err != nil {
		log.Fatalf("Error during parsing YAML: %v\n", err)
	}
	return conf
}
