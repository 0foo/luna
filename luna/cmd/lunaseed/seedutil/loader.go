package seedutil

import (
	"os"
	"gopkg.in/yaml.v3"
)

type SeedField map[string]string

func LoadSeedFile(path string) (SeedField, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var fields SeedField
	if err := yaml.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	return fields, nil
}
