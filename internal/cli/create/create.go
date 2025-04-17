package create

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mickamy/gob"
	"github.com/mickamy/gob/config"
)

var Cmd = &cobra.Command{
	Use:   "create",
	Short: "Create the database defined in your gob config",
	Long:  "Creates a database using the connection settings defined in .gob.yaml",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			fmt.Printf("❌ Failed to parse config file at %s: %w\n", config.Path, err)
		}
		return Run(cfg)
	},
}

func Run(cfg config.Config) error {
	if err := gob.Create(cfg); err != nil {
		if errors.Is(err, gob.ErrCreateDatabaseExists) {
			fmt.Printf("❌ Database %s already exists. Use --force to overwrite.\n", cfg.Database.Name)
			return nil
		}
		fmt.Printf("❌ Failed to create database %s: %s\n", cfg.Database.Name, err)
	}

	fmt.Printf("✅ Database %s created successfully!\n", cfg.Database.Name)
	return nil
}
