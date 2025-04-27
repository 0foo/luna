package seedutil

import (
	"database/sql"
	"fmt"
	"luna/config"
	"strings"
)

func InsertSeedData(query string, values []interface{}) error {
	// Open the database
	db, err := sql.Open("postgres", config.ConfigValues.DbURL)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	fmt.Println("Database opened sucessfully")

	// Execute the insert
	res, err := db.Exec(query, values...)
	if err != nil {
		return fmt.Errorf("failed to execute insert: %w", err)
	}
	
	fmt.Println("Database seed operation successful")
	fmt.Println("Database result: ")
	fmt.Println(res)
	return nil
}

// BuildBatchInsertStatement builds a batch INSERT SQL and values from multiple rows
func BuildBatchInsertStatement(tableName string, rows []map[string]string) (string, []interface{}, error) {
	if len(rows) == 0 {
		return "", nil, fmt.Errorf("no rows to insert")
	}

	// Consistent column order
	columns := make([]string, 0, len(rows[0]))
	for col := range rows[0] {
		columns = append(columns, fmt.Sprintf("\"%s\"", col))
	}

	valueGroups := make([]string, 0, len(rows))
	values := make([]interface{}, 0)

	paramCounter := 1

	// Now for each row
	for _, row := range rows {
		groupPlaceholders := make([]string, 0, len(columns))

		// Always iterate columns in order
		for _, colExpr := range columns {
			// Remove quotes from column name
			colName := strings.Trim(colExpr, "\"")

			groupPlaceholders = append(groupPlaceholders, fmt.Sprintf("$%d", paramCounter))
			values = append(values, row[colName])
			paramCounter++
		}

		valueGroups = append(valueGroups, fmt.Sprintf("(%s)", strings.Join(groupPlaceholders, ", ")))
	}

	query := fmt.Sprintf(
		"INSERT INTO \"%s\" (%s) VALUES %s",
		tableName,
		strings.Join(columns, ", "),
		strings.Join(valueGroups, ", "),
	)

	return query, values, nil
}
