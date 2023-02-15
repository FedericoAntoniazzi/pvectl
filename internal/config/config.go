package config

import (
	"github.com/kkyr/fig"
)

// Config represents the schema of the configuration file
// Every item can be set via env var as well
type Config struct {
	Endpoint string `fig:"endpoint"` // Endpoint of the proxmox server

	AllowInsecureConnection bool `fig:"allowInsecureConnection" default:"false"` // Allow connecting to a non secure endpoint

	Auth struct {
		Method string `fig:"method" default:"apitoken"` // Authentication method. Accepted values: "apitoken", "login"

		TokenID string `fig:"tokenId"` // used only when method is apitoken
		Secret  string `fig:"secret"`  // used only when method is apitoken

		Username string `fig:"username"` // used only when method is login
		Password string `fig:"password"` // used only when method is login
	}
}

// GetConfig reads the configuration file
func GetConfig() (Config, error) {
	var config Config
	err := fig.Load(&config, fig.File("pvectl.yaml"), fig.Dirs("."), fig.UseEnv("PVECTL"))
	return config, err
}
