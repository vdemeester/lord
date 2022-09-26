package generate

import (
	"github.com/spf13/cobra"
)

func DockerfileTemplate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dockerfile-template",
		Short: "Generate a dockerfile from lord",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	return cmd
}
