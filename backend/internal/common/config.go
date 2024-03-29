package common

import "github.com/spf13/viper"

func init() {
	viper.SetDefault("DATABASE_ADDRESS", "localhost:27017")
	viper.SetDefault("DATABASE_NAME", "personal-blog")
	viper.SetDefault("ARTICLE_COLLECTION", "articles")
	viper.SetDefault("THUMBNAILS_COLLECTION", "thumbnails")
	viper.SetDefault("ALLOWED_HTTP_METHODS", "GET, OPTIONS")
	viper.SetDefault("ALLOWED_HTTP_HEADERS", "Origin, Authorization, Content-Type")
	viper.SetDefault("ALLOWED_ORIGINS", "*")
	viper.SetDefault("BASE_HTTP_URL", "localhost")
	viper.SetDefault("STATIC_FILES_PATH", "/blog")
	viper.SetDefault("HTTP_PORT", 8081)
}

func BuildConfig() Config {
	return Config{
		Database: DatabaseConfig{
			Address: viper.GetString("DATABASE_ADDRESS"),
			Name:    viper.GetString("DATABASE_NAME"),
			Collections: Collections{
				Article:    viper.GetString("ARTICLE_COLLECTION"),
				Thumbnails: viper.GetString("THUMBNAILS_COLLECTION"),
			},
		},
		WebServer: WebServerConfig{
			AllowedMethods:      viper.GetString("ALLOWED_HTTP_METHODS"),
			AllowedOrigins:      viper.GetString("ALLOWED_ORIGINS"),
			AllowedHeaders:      viper.GetString("ALLOWED_HTTP_HEADERS"),
			BaseURL:             viper.GetString("BASE_HTTP_URL"),
			StaticFilesLocation: viper.GetString("STATIC_FILES_PATH"),
			Port:                viper.GetInt("HTTP_PORT")}}
}

type WebServerConfig struct {
	AllowedMethods      string
	AllowedHeaders      string
	AllowedOrigins      string
	BaseURL             string
	StaticFilesLocation string
	Port                int
}

type Collections struct {
	Thumbnails string
	Article    string
}

type DatabaseConfig struct {
	Address     string
	Name        string
	Collections Collections
}

type Config struct {
	Database  DatabaseConfig
	WebServer WebServerConfig
}
