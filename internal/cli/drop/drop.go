package drop

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/mickamy/godb"
	"github.com/mickamy/godb/config"
)

var Cmd = &cobra.Command{
	Use:   "drop",
	Short: "Drop the database defined in your godb config",
	Long:  "Drops a database using the connection settings defined in .godb.yaml",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.Load()
		if err != nil {
			fmt.Printf("❌ Failed to load config file at %s: %s\n", config.Path, err)
		}
		Run(cfg)
	},
}

func Run(cfg config.Config) {
	if err := godb.Drop(cfg); err != nil {
		fmt.Printf("❌ Failed to drop database '%s': %s\n", cfg.Database.Name, err)
		os.Exit(1)
	}

	fmt.Printf("✅ Database '%s' dropped successfully!\n", cfg.Database.Name)
}
