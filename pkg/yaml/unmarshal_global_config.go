package yaml

import (
	"brabus/pkg/dto"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func UnmarshalGlobalConfig() dto.Global {
	data, err := os.ReadFile("./worker/global.yaml")
	if err != nil {
		log.Fatalf("Error during reading file: %v\n", err)
	}

	var conf dto.Global
	if err := yaml.Unmarshal(data, &conf); err != nil {
		log.Fatalf("Error during parsing YAML: %v\n", err)
	}
	return conf
}
