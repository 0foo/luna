package lunadb

import (
	"database/sql"
	"fmt"
	"os"

	"luna/config"

	_ "github.com/lib/pq" // Postgres driver
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list-tables",
	Short: "List all non-system tables in the database",
	Long:  "Lists all user-defined tables excluding PostgreSQL system tables",
	Run:   runListTables,
}

func runListTables(cmd *cobra.Command, args []string) {
	db, err := sql.Open("postgres", config.ConfigValues.DbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	query := `
		SELECT schemaname, tablename
		FROM pg_catalog.pg_tables
		WHERE schemaname NOT IN ('pg_catalog', 'information_schema')
		ORDER BY schemaname, tablename;
	`

	rows, err := db.Query(query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to query tables: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	fmt.Println("Tables:")
	for rows.Next() {
		var schema string
		var table string
		if err := rows.Scan(&schema, &table); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to scan row: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("- %s.%s\n", schema, table)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Row iteration error: %v\n", err)
		os.Exit(1)
	}
}
