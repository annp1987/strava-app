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
	c := &Config{
		ServicePort:  confer.GetString("port"),
		SQLiteDBFile: confer.GetString("sqlite.db"),
		ClientID:     confer.GetString("client.id"),
		ClientSecret: confer.GetString("client.secret"),
		RedirectURL:  confer.GetString("redirect.url"),
	}
	return c
}
