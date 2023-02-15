package cmd

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"

	"github.com/FedericoAntoniazzi/pvectl/cmd/node"
	"github.com/FedericoAntoniazzi/pvectl/internal/config"
	"github.com/luthermonson/go-proxmox"
	"github.com/spf13/cobra"
)

const pvectlVersion = "dev"

// NewRootCmd creates a new pvectl command instance
func NewRootCmd(pveClient *proxmox.Client) *cobra.Command {
	// rootCmd represents the base command when called without any subcommands
	rootCmd := &cobra.Command{
		Use:   "pvectl",
		Short: "pvectl is a CLI tool to manage ProxmoxVE clusters",
		Long: `pvectl is a CLI tool to manage ProxmoxVE clusters

Find more information at https://github.com/FedericoAntoniazzi/pvectl`,
		Version: pvectlVersion,
	}

	// Add subcommands
	rootCmd.AddCommand(
		node.NewNodeCmd(pveClient),
	)

	// Add flags

	return rootCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cfg, err := config.GetConfig()
	if err != nil {
		os.Exit(2)
	}

	pveClient := createProxmoxClient(cfg)
	rootCmd := NewRootCmd(pveClient)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// createProxmoxClien creates a new proxmox client using provided configuration
func createProxmoxClient(cfg config.Config) *proxmox.Client {
	// Define the HTTP Client
	httpClient := http.DefaultClient
	httpClient.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: cfg.AllowInsecureConnection,
		},
	}

	// Default options
	options := []proxmox.Option{
		proxmox.WithUserAgent(fmt.Sprintf("pvectl/%s", pvectlVersion)),
		proxmox.WithClient(httpClient),
	}

	// Authentication methods
	if cfg.Auth.Method == "apitoken" {
		options = append(options, proxmox.WithAPIToken(cfg.Auth.TokenID, cfg.Auth.Secret))
	} else if cfg.Auth.Method == "login" {
		options = append(options, proxmox.WithLogins(cfg.Auth.Username, cfg.Auth.Password))
	}

	return proxmox.NewClient(cfg.Endpoint, options...)
}
