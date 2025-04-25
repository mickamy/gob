package rollback

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/mickamy/godb"
	"github.com/mickamy/godb/config"
)

var (
	step = 1
)

var Cmd = &cobra.Command{
	Use:   "rollback",
	Short: "Rollback the most recent database migration",
	Long: `Rollback applied database migrations.

By default, it rolls back the latest migration. 
You can specify the number of steps to rollback using the '--step' or '-s' flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.Load()
		if err != nil {
			fmt.Printf("❌ Failed to load config file at %s: %s\n", config.Path, err)
			os.Exit(1)
		}
		Run(cfg, step)
	},
}

func init() {
	Cmd.Flags().IntVarP(&step, "step", "s", 1, "Number of steps to rollback")
}

func Run(cfg config.Config, step int) {
	if err := godb.Rollback(cfg, step); err != nil {
		if errors.Is(err, godb.ErrMigrateNoChange) {
			fmt.Printf("✅ No changes to rollback database '%s'.\n", cfg.Database.Name)
			return
		}
		fmt.Printf("❌ Failed to rollback database '%s': %s\n", cfg.Database.Name, err)
		os.Exit(1)
	}

	fmt.Printf("✅ Database '%s' rolled back successfully!\n", cfg.Database.Name)
}
