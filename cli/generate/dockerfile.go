package generate

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
	"go.sbr.pm/lord/builders"
	"go.sbr.pm/lord/config"
)

const (
	dockerfileTemplate = `{{ range .Build.Steps -}}
FROM {{ .Image }} AS builder-{{ .Name }}
WORKDIR /app
COPY . /app
USER root
RUN {{ .Command }}
{{- end}}

FROM {{ .BaseImage }}
{{ range .Build.Steps -}}
COPY --from=builder-{{ .Name }} /app/{{ .Name }} /{{ .Name }}
{{- end }}
{{ range $key, $value := .Labels -}}
LABEL {{ $key }}={{ $value }}
{{- end }}
{{ range .Envs -}}
ENV {{ .Name }}={{ .Value }}
{{- end }}
ENTRYPOINT ["/{{ .Name }}"]`
)

func DockerfileTemplate() *cobra.Command {
	opts := &generateOptions{}
	cmd := &cobra.Command{
		Use:   "dockerfile",
		Short: "Generate a dockerfile from lord",
		RunE: func(cmd *cobra.Command, args []string) error {
			components := args
			c, err := config.Load(opts.file)
			if err != nil {
				return err
			}
			if len(args) == 0 {
				for n, _ := range c.Config {
					components = append(components, n)
				}
			}
			for _, component := range components {
				t, err := generateDockerfile(component, c.Config[component])
				if err != nil {
					return err
				}
				fmt.Println(t)
			}
			return nil
		},
	}
	cmd.Flags().StringVarP(&opts.file, "filename", "f", "lord.yaml", "Lord file to parse")
	return cmd
}

type data struct {
	BaseImage string
	Name      string
	Build     builders.Build
	Labels    map[string]string
	Envs      []config.Env
}

func generateDockerfile(name string, c config.Component) (string, error) {
	b, err := builders.Generate(name, c)
	if err != nil {
		return "", err
	}
	templateData := data{
		Name:      name,
		Build:     b,
		BaseImage: "cgr.dev/chainguard/static:latest",
		Labels:    c.Labels,
		Envs:      c.Envs,
	}

	var d strings.Builder
	t := template.Must(template.New("dockerfile").Parse(dockerfileTemplate))
	if err := t.Execute(&d, templateData); err != nil {
		return "", err
	}
	return d.String(), nil
}
