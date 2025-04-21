package migrate

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/mickamy/gob"
	"github.com/mickamy/gob/config"
)

var Cmd = &cobra.Command{
	Use:   "migrate",
	Short: "Apply database migrations",
	Long:  "Apply all up migrations using the configured database and migration directory.",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.Load()
		if err != nil {
			fmt.Printf("❌ Failed to load config file at %s: %s\n", config.Path, err)
			os.Exit(1)
		}
		Run(cfg)
	},
}

func Run(cfg config.Config) {
	if err := gob.Migrate(cfg); err != nil {
		if errors.Is(err, gob.ErrMigrateNoChange) {
			fmt.Printf("✅ No changes to apply for database '%s'.\n", cfg.Database.Name)
			return
		}
		fmt.Printf("❌ Failed to apply migrations for database '%s': %s\n", cfg.Database.Name, err)
		os.Exit(1)
	}

	fmt.Printf("✅ Migrations applied successfully for database '%s'!\n", cfg.Database.Name)
}
