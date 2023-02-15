package node

import (
	"errors"
	"log"
	"time"

	"github.com/FedericoAntoniazzi/pvectl/internal/printer"
	"github.com/luthermonson/go-proxmox"
	"github.com/spf13/cobra"
)

// mergeNodeStatus aggregates the NodeStatus structs retrieved from the /nodes and /cluster/status endpoints.
func mergeNodeStatus(node, clusterNode *proxmox.NodeStatus) (*proxmox.NodeStatus, error) {
	if node.ID != clusterNode.ID {
		return node, errors.New("cannot merge node status: node id does not match")
	}

	node.NodeID = clusterNode.NodeID
	node.Name = clusterNode.Name
	node.IP = clusterNode.IP
	node.Online = clusterNode.Online
	node.Local = clusterNode.Local

	return node, nil
}

// NewNodeListCmd represents the command to list nodes in a Proxmox cluster
func NewNodeListCmd(pveClient *proxmox.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List nodes",
		Long:  "List nodes",
		Run: func(cmd *cobra.Command, args []string) {
			nodes, err := pveClient.Nodes()
			if err != nil {
				log.Fatal(err)
			}

			cluster, err := pveClient.Cluster()
			if err != nil {
				log.Fatal(err)
			}

			for _, node := range nodes {
				for _, clusterNode := range cluster.Nodes {
					if node.ID == clusterNode.ID {
						node, err = mergeNodeStatus(node, clusterNode)
						if err != nil {
							log.Println(err)
						}
					}
				}
			}

			tp := printer.NewTabbedPrinter()
			tp.SetHeaders("NODE", "ADDRESS", "UPTIME")
			for _, n := range nodes {
				uptime := time.Duration(n.Uptime) * time.Second
				tp.AddRow(n.Node, n.IP, uptime)
			}
			tp.Print()
		},
	}

	// Add subcommands
	// rootCmd.AddCommand()

	// Add flags

	return cmd
}
