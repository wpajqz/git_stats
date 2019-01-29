package version

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RunCommand version command return api version for tracking binary in different server
func RunCommand(apiVersion, gitCommit, built string) *cobra.Command {
	return &cobra.Command{
		Use:     "version",
		Short:   "The version of stats",
		Aliases: []string{"v"},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(os.Stdout, "Server:")
			fmt.Fprintln(os.Stdout, " Api Version:         ", apiVersion)
			fmt.Fprintln(os.Stdout, " Git commit:          ", gitCommit)
			fmt.Fprintln(os.Stdout, " Built:               ", built)
		},
	}
}
