package generate

import (
	"github.com/spf13/cobra"
)

type generateOptions struct {
	file string
}

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate something from lord",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	cmd.AddCommand(DockerfileTemplate())
	cmd.AddCommand(TektonTask())
	return cmd
}
