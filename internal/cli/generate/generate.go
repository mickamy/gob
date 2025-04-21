package generate

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"g"},
	Short:   "Generate migration files or other database-related assets",
	Long: `The 'generate' command helps you scaffold migration files and other assets
for your database workflows.

Use this to generate SQL migration files or prepare directories used by tools like golang-migrate.
Currently, it only supports generating SQL migration files w/ golang-migrate.
`,
}

func init() {
	Cmd.AddCommand(
		migrationCmd,
	)
}
