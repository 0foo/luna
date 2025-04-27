package lunaseed

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "seed",
	Short: "Run database seed",
	Long:  "seed: a CLI tool for managing seeds",
}

func init() {
	Cmd.AddCommand(RunCmd)
}
