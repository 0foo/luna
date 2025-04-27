package lunadb

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"luna/config"

	_ "github.com/lib/pq" // Postgres driver
	"github.com/spf13/cobra"
)

var RawCmd = &cobra.Command{
	Use:   "raw [SQL]",
	Short: "Run a raw SQL command",
	Long:  "Run any raw SQL against the database",
	Args:  cobra.MinimumNArgs(1),
	Run:   runRaw,
}

func runRaw(cmd *cobra.Command, args []string) {
	sqlCmd := strings.Join(args, " ")

	db, err := sql.Open("postgres", config.ConfigValues.DbURL)
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		os.Exit(1)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		fmt.Println("Failed to ping database:", err)
		os.Exit(1)
	}

	sqlCmdLower := strings.ToLower(strings.TrimSpace(sqlCmd))
	if strings.HasPrefix(sqlCmdLower, "select") {
		rows, err := db.Query(sqlCmd)
		if err != nil {
			fmt.Println("Failed to execute query:", err)
			os.Exit(1)
		}
		defer rows.Close()

		cols, err := rows.Columns()
		if err != nil {
			fmt.Println("Failed to get columns:", err)
			os.Exit(1)
		}

		rowCount := 0

		for rows.Next() {
			rowCount++

			columns := make([]interface{}, len(cols))
			columnPointers := make([]interface{}, len(cols))

			for i := range columns {
				columnPointers[i] = &columns[i]
			}

			if err := rows.Scan(columnPointers...); err != nil {
				fmt.Println("Failed to scan row:", err)
				os.Exit(1)
			}

			rowMap := make(map[string]interface{})
			for i, colName := range cols {
				val := columns[i]
				if b, ok := val.([]byte); ok {
					rowMap[colName] = string(b)
				} else {
					rowMap[colName] = val
				}
			}

			fmt.Println(rowMap)
		}

		if rowCount == 0 {
			fmt.Println("(no rows returned)")
		}

		if err := rows.Err(); err != nil {
			fmt.Println("Error reading rows:", err)
			os.Exit(1)
		}
	} else {
		_, err = db.Exec(sqlCmd)
		if err != nil {
			fmt.Println("Failed to execute SQL:", err)
			os.Exit(1)
		}
		fmt.Println("SQL command executed successfully")
	}
}
