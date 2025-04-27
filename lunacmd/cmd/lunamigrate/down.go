package lunamigrate

import (
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq" // Postgres driver
	"github.com/spf13/cobra"
	"luna/config"
)

var DownCmd = &cobra.Command{
	Use:   "down",
	Short: "Rollback the last migration",
	Run:   runDown,
}

func runDown(cmd *cobra.Command, args []string) {
	m, err := migrate.New(
		"file://"+config.ConfigValues.MigrationsDir, // using ConfigValues
		config.ConfigValues.DbURL,                   // using ConfigValues
	)
	if err != nil {
		fmt.Println("Migration init error:", err)
		os.Exit(1)
	}

	if err := m.Steps(-1); err != nil && err != migrate.ErrNoChange {
		fmt.Println("Migration DOWN failed:", err)
		os.Exit(1)
	}

	fmt.Println("Migrations DOWN (rollback) completed.")
}
