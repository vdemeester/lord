package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"go.sbr.pm/lord/cli/generate"
	"sigs.k8s.io/release-utils/version"
	"sigs.k8s.io/yaml"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "lord",
		DisableAutoGenTag: true,
		SilenceUsage:      true,
	}

	cmd.AddCommand(Parse())
	cmd.AddCommand(generate.New())
	cmd.AddCommand(version.Version())
	return cmd
}

type config map[string]interface{}

type parseOptions struct {
	file string
}

func Parse() *cobra.Command {
	opts := &parseOptions{}
	cmd := &cobra.Command{
		Use:   "parse",
		Short: "Parse a lord file",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(opts.file)
			data, err := os.ReadFile(opts.file)
			if err != nil {
				return err
			}
			c := config{}
			if err := yaml.Unmarshal(data, &c); err != nil {
				return err
			}
			fmt.Println("c", len(c), c)
			fmt.Println("c[config]", len(c["config"].(map[string]interface{})), c["config"])
			fmt.Println("c[config][openshift-pipeline-controller][build]", c["config"].(map[string]interface{})["openshift-pipeline-controller"].(map[string]interface{})["build"])
			return nil
		},
	}
	cmd.Flags().StringVarP(&opts.file, "filename", "f", "", "Lord file to parse")
	return cmd
}
