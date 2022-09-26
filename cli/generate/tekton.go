package generate

import (
	"github.com/spf13/cobra"
)

func TektonTask() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tekton-task",
		Short: "Generate a task from lord",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	return cmd
}
