package golang

import (
	"fmt"

	"go.sbr.pm/lord/builders"
	"go.sbr.pm/lord/config"
)

func init() {
	builders.Register("go", goBuild)
}

func goBuild(name string, c config.Component) builders.Build {
	goimage := "cgr.dev/chainguard/go:1.19"
	gopackage := "."
	with := c.Build.With
	if v, ok := with["image"]; ok {
		goimage = v.(string)
	}
	if v, ok := with["package"]; ok {
		gopackage = v.(string)
	}
	b := builders.Build{
		Steps: []builders.Step{{
			Name:    name,
			Image:   goimage,
			Command: fmt.Sprintf("go build -v -o %s %s", name, gopackage),
		}},
	}
	return b
}
