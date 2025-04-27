package lunaseed

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"luna/cmd/lunaseed/seedutil"

	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

func GetTableName(seedFile string) string {
	base := filepath.Base(seedFile)           // ex: "user.yaml"
	name := strings.TrimSuffix(base, ".yaml") // remove .yaml
	name = strings.TrimSuffix(name, ".yml")   // (support .yml too)
	return name
}

var RunCmd = &cobra.Command{
	Use:   "run <seedfile> <count>",
	Short: "Run a seed file",
	Args:  cobra.ExactArgs(2),
	Run:   RunMe,
}

func RunMe(cmd *cobra.Command, args []string) {
	seedFile := args[0]
	count, err := strconv.Atoi(args[1])
	tableName := GetTableName(seedFile)

	if err != nil {
		fmt.Printf("Invalid count: %v\n", err)
		return
	}

	fields, err := seedutil.LoadSeedFile(seedFile)
	if err != nil {
		fmt.Printf("Failed to load seed file: %v\n", err)
		return
	}

	// get fake data
	rows, err := seedutil.GetFakeData(fields, count)

	if err != nil {
		fmt.Println("Error getting Faker data:", err)
		return
	}
	// fmt.Println(rows)
	// fmt.Println("------------------------------")

	// Build the insert statement
	query, values, err := seedutil.BuildBatchInsertStatement(tableName, rows)

	if err != nil {
		fmt.Println("Error building batch insert:", err)
		return
	}
	fmt.Println(query, values)

	// insert data into db
	err = seedutil.InsertSeedData(query, values)
	if err != nil {
		fmt.Println("Database seed failed:", err)
		os.Exit(1)
	}

	fmt.Println("Seeding complete")
}
