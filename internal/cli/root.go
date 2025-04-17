package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/mickamy/gob/internal/cli/version"
)

var cmd = &cobra.Command{
	Use:   "gob",
	Short: "A lightweight DB management CLI for Go projects",
	Long: `gob is a database management tool designed for Go projects.

It helps you handle database lifecycle tasks such as creation, migration, rollback, seeding, and more — all from a single, developer-friendly CLI.`,
}

func init() {
	cmd.AddCommand(version.Cmd)
}

func Execute() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
