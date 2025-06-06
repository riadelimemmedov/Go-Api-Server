package config

// ! Config holds the entire configuration for the API server.
type Config struct {
	Server ServerConfig `mapstructure:"server"`
	App    AppConfig    `mapstructure:"app"`
	CORS   CORSConfig   `mapstructure:"cors"`
	RateLimit RateLimitConfig `mapstructure:"rate_limit"`
}

// ! ServerConfig holds the configuration for the API server.
type ServerConfig struct {
	Host         string `mapstructure:"server"`
	Port         string `mapstructure:"port"`
	Mode         string `mapstructure:"mode"`
	ReadTimeout  int    `mapstructure:"read_timeout"`
	WriteTimeout int    `mapstructure:"write_timeout"`
}

// ! AppConfig holds the application-level configuration.
type AppConfig struct {
	Name        string `mapstructure:"name"`
	Version     string `mapstructure:"version"`
	Environment string `mapstructure:"environment"`
}

// ! CORSConfig holds the CORS configuration for the API server.
type CORSConfig struct {
	Enabled        bool     `mapstructure:"enabled"`
	AllowedOrigins []string `mapstructure:"allowed_origins"`
	AllowedMethods []string `mapstructure:"allowed_methods"`
	AllowedHeaders []string `mapstructure:"allowed_headers"`
}

// ! RateLimitConfig holds the rate limiting configuration for the API server.
type RateLimitConfig struct {
	Enabled   bool `mapstructure:"enabled"`
	RPS       int  `mapstructure:"rps"`
	Burst     int  `mapstructure:"burst"`
	CleanupMS int  `mapstructure:"cleanup_ms"`
}
