package generate

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate [args]",
	Short: "Forward to golang-migrate CLI (binary or go tool)",
	Long: `Forwards arguments to the golang-migrate CLI.

If 'migrate' is not installed as a binary, it falls back to 'go tool migrate'.
`,
	Run: func(cmd *cobra.Command, args []string) {
		Migrate(args)
	},
}

func Migrate(args []string) error {
	binPath, err := exec.LookPath("migrate")
	var cmd *exec.Cmd

	if err == nil {
		// Found binary
		cmd = exec.Command(binPath, args...)
	} else {
		// Fallback to 'go tool migrate'
		fmt.Println("ℹ️  'migrate' not found in PATH. Falling back to 'go tool migrate'")
		cmd = exec.Command("go", append([]string{"tool", "migrate"}, args...)...)
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		var execErr *exec.Error
		if errors.As(err, &execErr) {
			fmt.Println("❌  'migrate' command not found in PATH.")
			os.Exit(1)
		}
		if errors.Is(err, os.ErrNotExist) {
			// Command not found
			fmt.Println("❌  'migrate' command not found in PATH.")
			os.Exit(1)
		} else {
			// Other error
			fmt.Printf("❌  Failed to run 'migrate': %s\n", err)
			os.Exit(1)
		}
	}

	fmt.Println("✅  Generated migration successfully!")

	return nil
}
