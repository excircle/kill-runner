package killrunner

import (
	"fmt"
	"os"

	"github.com/excircle/kill-runner/cmd/questions" // Import the questions package

	"github.com/spf13/cobra"
)

// stageCmd represents the stage command
var unstageCmd = &cobra.Command{
	Use:   "unstage [question]",
	Short: "Unstage a Kubernetes exam question",
	Long: `This command unstages a Kubernetes exam question, executing the relevant
logic for the specified question. Example usage:

  kr unstage q1    # Takes down the staged Kubenetes objects for question 1
  kr unstage q2    # Takes down the staged Kubenetes objects for question 2
  kr unstage q3    # Takes down the staged Kubenetes objects for question 3
`,
	Args: cobra.ExactArgs(1), // Ensure exactly one argument is provided
	Run: func(cmd *cobra.Command, args []string) {
		question := args[0]

		switch question {
		case "q1":
			questions.UnstageQ1()
		case "q2":
			questions.UnstageQ2()
		case "q3":
			questions.UnstageQ3()
		case "q4":
			questions.UnstageQ4()
		case "q6":
			questions.UnstageQ6()
		default:
			fmt.Println("Invalid question. Please use q1, q2, or q3.")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(unstageCmd)
}
