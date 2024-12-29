package yaml

import (
	"brabus/configs"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func UnmarshalGlobalConfig() config.Global {
	data, err := os.ReadFile("./configs/global.yaml")
	if err != nil {
		log.Fatalf("Error during reading file: %v\n", err)
	}

	var conf config.Global
	if err := yaml.Unmarshal(data, &conf); err != nil {
		log.Fatalf("Error during parsing YAML: %v\n", err)
	}
	return conf
}
