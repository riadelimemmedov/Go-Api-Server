package config

import (
	"os"
	"strings"

	logger "github.com/riadalimemmedov/api-server-project/api-server-project/pkg/log"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// ! Load loads the configuration from the specified config file based on the environment.
func Load() *Config {
	// Initialize the configuration
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath(".")

	// Set default values
	setDefaults()

	// Enable environment variables
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Determine environment first
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = envLocal
	}

	// Set config file based on environment
	configFile := _getConfigFilename(env)

	// Read the specific config file
	viper.SetConfigName(configFile)
	if err := viper.ReadInConfig(); err != nil {
		logger.GetLogger().Error("Error reading config file: %v\n", zap.Error(err))
	}

	logger.GetLogger().Info("Using environment", zap.String("env", env))
	logger.GetLogger().Info("Using config file", zap.String("configFile", configFile))

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		logger.GetLogger().Error("Error unmarshalling config: %v", zap.Error(err))
	}
	logger.GetLogger().Info("Configuration loaded successfully")
	return &config
}

// ! setDefaults sets the default values for the configuration.
func setDefaults() {
	// Server defaults
	viper.SetDefault("server.host", "0.0.0.0")
	viper.SetDefault("server.port", "3000")
	viper.SetDefault("server.mode", "debug")
	viper.SetDefault("server.read_timeout", 30)
	viper.SetDefault("server.write_timeout", 30)

	//App defaults
	viper.SetDefault("app.name", "Go API")
	viper.SetDefault("app.version", "1.0.0")
	viper.SetDefault("app.environment", "development")

	//CORS defaults
	viper.SetDefault("cors.enabled", true)
	viper.SetDefault("cors.allowed_origins", []string{"*"})
	viper.SetDefault("cors.allowed_methods", []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	viper.SetDefault("cors.allowed_headers", []string{"*"})

	//Rate limit defaults
	viper.SetDefault("rate_limit.enabled", true)
	viper.SetDefault("rate_limit.rps", 100)
	viper.SetDefault("rate_limit.burst", 200)
	viper.SetDefault("rate_limit.cleanup_ms", 60000)
}

// ! IsLocal checks if the application is running in local environment.
func (c *Config) IsLocal() bool {
	return c.App.Environment == envLocal
}

// ! IsDevelopment checks if the application is running in development environment.
func (c *Config) IsDevelopment() bool {
	return c.App.Environment == envDevelopment
}

// ! IsProduction checks if the application is running in production environment.
func (c *Config) IsProduction() bool {
	return c.App.Environment == envProduction
}

// ! GetConfigFilename returns the configuration filename based on the environment.
func _getConfigFilename(env string) string {
	switch strings.ToLower(env) {
	case "development":
		return "config.development"
	case "production":
		return "config.production"
	case "local":
		return "config"
	default:
		logger.GetLogger().Warn("Unknown environment: %s, using default config", zap.String("env", env))
		return "config"
	}
}
