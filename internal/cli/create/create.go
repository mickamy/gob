package create

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/mickamy/gob"
	"github.com/mickamy/gob/config"
)

var Cmd = &cobra.Command{
	Use:   "create",
	Short: "Create the database defined in your gob config",
	Long:  "Creates a database using the connection settings defined in .gob.yaml",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.Load()
		if err != nil {
			fmt.Printf("❌ Failed to load config file at %s: %w\n", config.Path, err)
		}
		Run(cfg)
	},
}

func Run(cfg config.Config) {
	if err := gob.Create(cfg); err != nil {
		if errors.Is(err, gob.ErrCreateDatabaseExists) {
			fmt.Printf("✅ Database '%s' already exists.\n", cfg.Database.Name)
			return
		}
		fmt.Printf("❌ Failed to create database '%s': %s\n", cfg.Database.Name, err)
		os.Exit(1)
	}

	fmt.Printf("✅ Database '%s' created successfully!\n", cfg.Database.Name)
}
