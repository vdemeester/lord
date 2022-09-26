package cli

import (
	"github.com/spf13/cobra"
	"go.sbr.pm/lord/cli/generate"
	"sigs.k8s.io/release-utils/version"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "lord",
		DisableAutoGenTag: true,
		SilenceUsage:      true,
	}

	cmd.AddCommand(generate.New())
	cmd.AddCommand(version.Version())
	return cmd
}
