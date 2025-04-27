package lunamigrate

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"luna/config"
)

var CreateCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "Create a new migration",
	Args:  cobra.ExactArgs(1),
	Run:   runCreate,
}

func runCreate(cmd *cobra.Command, args []string) {
	name := strings.TrimSpace(args[0])
	if name == "" {
		fmt.Println("Migration name cannot be empty.")
		os.Exit(1)
	}

	timestamp := time.Now().Format("20060102150405") // YYYYMMDDHHMMSS
	baseName := fmt.Sprintf("%s_%s", timestamp, name)

	upFile := filepath.Join(config.ConfigValues.MigrationsDir, baseName+".up.sql")
	downFile := filepath.Join(config.ConfigValues.MigrationsDir, baseName+".down.sql")

	// Create empty .up.sql file
	if err := os.WriteFile(upFile, []byte("-- Write your UP migration here\n"), 0644); err != nil {
		fmt.Println("Failed to create UP migration file:", err)
		os.Exit(1)
	}

	// Create empty .down.sql file
	if err := os.WriteFile(downFile, []byte("-- Write your DOWN migration here\n"), 0644); err != nil {
		fmt.Println("Failed to create DOWN migration file:", err)
		os.Exit(1)
	}

	fmt.Println("Created migration files:")
	fmt.Println(" -", upFile)
	fmt.Println(" -", downFile)
}
