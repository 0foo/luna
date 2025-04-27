package lunamigrate

import (
	"github.com/spf13/cobra"
	"luna/config"
)

var Cmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	Long:  "migrate: a CLI tool for managing migrations",
}

var migrationsDir string

func init() {
	Cmd.PersistentFlags().StringVar(&migrationsDir, "migrations", "", "location of migrations directory (default is ./migrations)")

	// If --migrations flag is passed, override the config
	if migrationsDir != "" {
		config.ConfigValues.MigrationsDir = migrationsDir
	}

	// If still empty, fallback to default
	if config.ConfigValues.MigrationsDir == "" {
		config.ConfigValues.MigrationsDir = "./migrations"
	}

	Cmd.AddCommand(UpCmd)
	Cmd.AddCommand(DownCmd)
	Cmd.AddCommand(ForceCmd)
	Cmd.AddCommand(CreateCmd)
}
