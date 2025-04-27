package config

// Config holds all configuration values
type Config struct {
	DbURL         string `mapstructure:"db_url"`
	MigrationsDir string `mapstructure:"migrations_dir"`
	SeedsDir      string `mapstructure:"seeds_dir"`
}

var ConfigValues Config
