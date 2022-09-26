package cli

import (
	"github.com/spf13/cobra"
	"sigs.k8s.io/release-utils/version"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "lord",
		DisableAutoGenTag: true,
		SilenceUsage:      true,
	}

	cmd.AddCommand(version.Version())
	return cmd
}
