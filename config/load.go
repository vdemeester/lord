package config

import (
	"os"

	"sigs.k8s.io/yaml"
)

type Config struct {
	Config map[string]Component
}

type Component struct {
	Build     Build
	Test      Test
	Artifacts []Artifact
	Labels    map[string]string
	Envs      []Env
}

type Build struct {
	Type string
	With map[string]interface{}
}

type Test struct {
	Type string
	With map[string]interface{}
}

type Artifact struct {
	Name string
}

type Env struct {
	Name  string
	Value string
}

func Load(path string) (Config, error) {
	var c Config
	data, err := os.ReadFile(path)
	if err != nil {
		return c, err
	}
	if err := yaml.Unmarshal(data, &c); err != nil {
		return c, err
	}
	return c, nil
}
