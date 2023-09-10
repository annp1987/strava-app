package config

import (
	"github.com/spf13/viper"
	"strings"
)

// Config holds configuration details of the application
type Config struct {
	ServicePort  string
	SQLiteDBFile string
	ClientID     string
	ClientSecret string
	RedirectURL  string
}

// NewConfig generator
func NewConfig() *Config {
	confer := viper.New()
	confer.AutomaticEnv()
	confer.SetEnvPrefix("")
	confer.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// set default service port is 8080
	confer.SetDefault("port", "8080")
	confer.SetDefault("sqlite.db", "strava.db")
	confer.SetDefault("client.id", "113411")
	confer.SetDefault("client.secret", "bdc9afce6522c921248a91e84f73102a6cf21680")
	confer.SetDefault("redirect.url", "https://parkrun.online")
	c := &Config{
		ServicePort:  confer.GetString("port"),
		SQLiteDBFile: confer.GetString("sqlite.db"),
		ClientID:     confer.GetString("client.id"),
		ClientSecret: confer.GetString("client.secret"),
		RedirectURL:  confer.GetString("redirect.url"),
	}
	return c
}
