package builders

import (
	"fmt"

	"go.sbr.pm/lord/config"
)

var (
	builders = map[string]Builder{}
)

type Build struct {
	Steps []Step
}

type Step struct {
	Name    string
	Image   string
	Command string
}

type Builder func(string, config.Component) Build

func Register(n string, b Builder) {
	builders[n] = b
}

func Generate(name string, c config.Component) (Build, error) {
	builder, ok := builders[c.Build.Type]
	if !ok {
		return Build{}, fmt.Errorf("Builder %s not supported", c.Build.Type)
	}
	return builder(name, c), nil
}
