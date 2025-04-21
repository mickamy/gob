package reset

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/mickamy/gob"
	"github.com/mickamy/gob/config"
)

var Cmd = &cobra.Command{
	Use:   "reset",
	Short: "Drop, recreate, and migrate the database",
	Long: `Reset the database by dropping it, recreating it, and applying all migrations.

Useful for local development when you want a fresh schema.
This command is equivalent to running 'gob drop', 'gob create', and 'gob migrate' in sequence.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.Load()
		if err != nil {
			fmt.Printf("❌  Failed to load config file at %s: %s\n", config.Path, err)
		}
		Run(cfg)
	},
}

func Run(cfg config.Config) {
	if err := gob.Reset(cfg); err != nil {
		fmt.Printf("❌  Failed to reset database '%s': %s\n", cfg.Database.Name, err)
		os.Exit(1)
	}

	fmt.Printf("✅  Database '%s' resetted successfully!\n", cfg.Database.Name)
}
