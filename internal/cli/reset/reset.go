package reset

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/mickamy/godb"
	"github.com/mickamy/godb/config"
)

var (
	force bool
)

var Cmd = &cobra.Command{
	Use:   "reset",
	Short: "Drop, recreate, and migrate the database",
	Long: `Reset the database by dropping it, recreating it, and applying all migrations.

Useful for local development when you want a fresh schema.
This command is equivalent to running 'godb drop', 'godb create', and 'godb migrate' in sequence.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.Load()
		if err != nil {
			fmt.Printf("❌ Failed to load config file at %s: %s\n", config.Path, err)
		}
		Run(cfg, force)
	},
}

func init() {
	Cmd.Flags().BoolVarP(&force, "force", "f", false, "Terminate the connections to the database before dropping it")
}

func Run(cfg config.Config, force bool) {
	if err := godb.Reset(cfg, force); err != nil {
		fmt.Printf("❌ Failed to reset database '%s': %s\n", cfg.Database.Name, err)
		os.Exit(1)
	}

	fmt.Printf("✅ Database '%s' resetted successfully!\n", cfg.Database.Name)
}
