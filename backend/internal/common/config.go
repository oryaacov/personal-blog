package common

import "github.com/spf13/viper"

var Configuration *Config

func init() {
	viper.SetDefault("DB_DRIVER", "postgres")
	viper.SetDefault("DB_USERNAME", "postgres")
	viper.SetDefault("ALLOWED_HTTP_METHODS", "GET, OPTIONS")
	viper.SetDefault("ALLOWED_HTTP_HEADERS", "Origin, Authorization, Content-Type")
	viper.SetDefault("ALLOWED_ORIGINS", "*")
	viper.SetDefault("BASE_HTTP_URL", "/")
	viper.SetDefault("STATIC_FILES_PATH", "/blog")
}

func BuildConfig() *Config {
	return &Config{
		ConnectionString: viper.GetString("DB_USERNAME"),
		Database:         viper.GetString("DB_DRIVER"),
		WebServer: WebServerConfig{
			AllowedMethods:      viper.GetString("ALLOWED_HTTP_METHODS"),
			AllowedOrigins:      viper.GetString("ALLOWED_ORIGINS"),
			AllowedHeaders:      viper.GetString("ALLOWED_HTTP_HEADERS"),
			StaticFilesLocation: viper.GetString("BASE_HTTP_URL"), BaseURL: viper.GetString("STATIC_FILES_PATH")}}
}

type WebServerConfig struct {
	AllowedMethods      string
	AllowedHeaders      string
	AllowedOrigins      string
	BaseURL             string
	StaticFilesLocation string
}

type Config struct {
	ConnectionString string
	Database         string
	WebServer        WebServerConfig
}
