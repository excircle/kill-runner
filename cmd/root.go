package killrunner

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kr",
	Short: "Kill-Runner CLI tool for running Killer.sh scenarios",
	Long: `Kill-Runner (kr) runs Killer.sh scenarios against a target Kubernetes cluster.

Find more information at: https://github.com/excircle/kill-runner`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("(Kill-Runner Help Menu) kr runs Killer.sh scenarios against a target Kubernetes cluster.")
		fmt.Println("\nFind more information at: https://github.com/excircle/kill-runner")
		fmt.Println("\nBasic Commands (Beginner):")
		fmt.Println("  state           Checks the state of the current Kill Runner challenge")
		fmt.Println("\nUse \"kr [command] --help\" for more information about a command.")
	},
}

// Execute runs the root command
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
