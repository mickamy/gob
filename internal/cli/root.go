package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/mickamy/gob/internal/cli/create"
	"github.com/mickamy/gob/internal/cli/drop"
	"github.com/mickamy/gob/internal/cli/generate"
	initPkg "github.com/mickamy/gob/internal/cli/init"
	"github.com/mickamy/gob/internal/cli/migrate"
	"github.com/mickamy/gob/internal/cli/version"
)

var cmd = &cobra.Command{
	Use:   "gob",
	Short: "A lightweight DB management CLI for Go projects",
	Long: `gob is a database management tool designed for Go projects.

It helps you handle database lifecycle tasks such as creation, migration, rollback, seeding, and more — all from a single, developer-friendly CLI.`,
}

func init() {
	cmd.AddCommand(create.Cmd)
	cmd.AddCommand(drop.Cmd)
	cmd.AddCommand(generate.Cmd)
	cmd.AddCommand(initPkg.Cmd)
	cmd.AddCommand(migrate.Cmd)
	cmd.AddCommand(version.Cmd)
}

func Execute() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
