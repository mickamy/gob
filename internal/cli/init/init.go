package init

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/mickamy/gob"
)

var Cmd = &cobra.Command{
	Use:   "init",
	Short: "Create a new gob config with interactive prompts",
	Long: `Initialize a new gob project by selecting database engine,
connection settings, and migration path via interactive prompts.

This will generate a configuration file (e.g. .gob.yaml) for use with other gob commands.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return Run()
	},
}

func Run() error {
	if _, err := os.Stat(gob.ConfigPath); err == nil {
		fmt.Printf("⚠️  %s already exists. Overwrite? [y/N]: ", gob.ConfigPath)
		var res string
		if _, err := fmt.Scanln(&res); err != nil {
			return fmt.Errorf("failed to read input: %w", err)
		}
		if res != "y" && res != "Y" {
			fmt.Println("❌ Canceled.")
			return nil
		}
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("🚀 Welcome to gob init!")
	fmt.Println("🔧 Select database engine:")
	fmt.Println("   1) PostgreSQL")
	fmt.Println("   2) MySQL")
	fmt.Print("👉 Enter choice [1-2]: ")

	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	var driver string
	switch choice {
	case "1":
		driver = "postgres"
	case "2":
		driver = "mysql"
	default:
		fmt.Println("❌ Invalid choice.")
		return nil
	}

	fmt.Printf("✅ You selected: %s\n", driver)

	var host, port, user, password, name, migrationDir string

	fmt.Print("🌐 Enter database host (default: localhost): ")
	host, _ = reader.ReadString('\n')
	host = strings.TrimSpace(host)
	if host == "" {
		host = "localhost"
	}

	var defaultPort string
	if driver == "postgres" {
		defaultPort = "5432"
	} else {
		defaultPort = "3306"
	}
	fmt.Printf("🔌 Enter database port (default: %s): ", defaultPort)
	port, _ = reader.ReadString('\n')
	port = strings.TrimSpace(port)
	if port == "" {
		port = defaultPort
	}
	portNumber, err := strconv.Atoi(port)
	if err != nil {
		fmt.Println("❌ Invalid port number.")
		return nil
	}

	var defaultUser string
	if driver == "postgres" {
		defaultUser = "postgres"
	} else {
		defaultUser = "root"
	}
	fmt.Printf("👤 Enter database user (default: %s): ", defaultUser)
	user, _ = reader.ReadString('\n')
	user = strings.TrimSpace(user)
	if user == "" {
		user = defaultUser
	}

	fmt.Print("🔑 Enter database password (default: password): ")
	password, _ = reader.ReadString('\n')
	password = strings.TrimSpace(password)
	if password == "" {
		password = "password"
	}

	fmt.Print("📛 Enter database name: ")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)
	if name == "" {
		fmt.Println("❌ Database name is required.")
		return nil
	}

	fmt.Print("📂 Enter migration directory (default: migrations): ")
	migrationDir, _ = reader.ReadString('\n')
	migrationDir = strings.TrimSpace(migrationDir)
	if migrationDir == "" {
		migrationDir = "migrations"
	}

	cfg := gob.Config{
		Database: gob.Database{
			Driver:   driver,
			Host:     host,
			Port:     portNumber,
			User:     user,
			Password: password,
			Name:     name,
		},
		Migrations: gob.Migrations{
			Dir: migrationDir,
		},
	}

	data, err := yaml.Marshal(cfg)
	if err != nil {
		fmt.Printf("❌ Failed to marshal config: %v\n", err)
		return nil
	}

	err = os.WriteFile(gob.ConfigPath, data, 0644)
	if err != nil {
		fmt.Printf("❌ Failed to write config file: %v\n", err)
		return nil
	}

	fmt.Printf("✅ Successfully generated %s!\n", gob.ConfigPath)
	return nil
}
