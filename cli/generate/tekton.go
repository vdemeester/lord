package generate

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"go.sbr.pm/lord/config"
)

var (
	taskTemplate = `apiVersion: tekton.dev/v1beta1
kind: Task
metadata:
  name: {{ name }}
spec:
  params:
  - name: reference
  - name: base-image
    default: {{ build.defaultBaseImage }}
  workspaces:
  - name: source
  
`
)

func TektonTask() *cobra.Command {
	opts := &generateOptions{}
	cmd := &cobra.Command{
		Use:   "tekton-task",
		Short: "Generate a task from lord",
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
			fmt.Println("components to generate:", components)
			fmt.Printf("config: %+v\n", c)
			fmt.Println("---")
			for _, component := range components {
				t, err := generateTektonTask(c.Config[component])
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

func generateTektonTask(c config.Component) (string, error) {
	var sb strings.Builder
	return sb.String(), nil
}
