package lunadb

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "db",
	Short: "Run database commands",
	Long:  "db: a cli tool for running db commands",
}


func init() {
	Cmd.AddCommand(RawCmd)
	Cmd.AddCommand(ListCmd)
}
