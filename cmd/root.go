package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// NewRootCmd creates a new pvectl command instance
func NewRootCmd() *cobra.Command {
	// rootCmd represents the base command when called without any subcommands
	rootCmd := &cobra.Command{
		Use:   "pvectl",
		Short: "pvectl is a CLI tool to manage ProxmoxVE clusters",
		Long: `pvectl is a CLI tool to manage ProxmoxVE clusters

Find more information at https://github.com/FedericoAntoniazzi/pvectl`,
	}

	// Add subcommands
	// rootCmd.AddCommand()

	// Add flags

	return rootCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd := NewRootCmd()

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
