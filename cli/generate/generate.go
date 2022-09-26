package generate

import (
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate something from lord",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	// cmd.AddCommand(Tekton())
	// cmd.AddCommand(Dockerfile())
	cmd.AddCommand(DockerfileTemplate())
	return cmd
}
