package node

import (
	"github.com/luthermonson/go-proxmox"
	"github.com/spf13/cobra"
)

// NewNodeCmd represents the root of the command to manage nodes
func NewNodeCmd(pveClient *proxmox.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "node",
		Short: "Manage node(s)",
		Long:  "Manage node(s)",
	}

	// Add subcommands
	cmd.AddCommand(
		NewNodeListCmd(pveClient),
	)

	// Add flags

	return cmd
}
