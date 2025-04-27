package lunamigrate

import (
	"fmt"
	"os"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"luna/config"
)

var ForceCmd = &cobra.Command{
	Use:   "force [version]",
	Short: "Force set the migration version",
	Args:  cobra.ExactArgs(1),
	Run:   runForce,
}

func runForce(cmd *cobra.Command, args []string) {
	versionStr := args[0]
	version, err := strconv.Atoi(versionStr)
	if err != nil {
		fmt.Println("Invalid version number:", versionStr)
		os.Exit(1)
	}

	m, err := migrate.New(
		"file://"+config.ConfigValues.MigrationsDir,
		config.ConfigValues.DbURL,
	)
	if err != nil {
		fmt.Println("Migration init error:", err)
		os.Exit(1)
	}

	err = m.Force(version)
	if err != nil {
		fmt.Println("Force migration failed:", err)
		os.Exit(1)
	}

	fmt.Printf("Forced migration version to %d successfully.\n", version)
}
