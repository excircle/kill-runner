package killrunner

import (
	"fmt"
	"os"

	"github.com/excircle/kill-runner/cmd/questions" // Import the questions package

	"github.com/spf13/cobra"
)

// stageCmd represents the stage command
var stageCmd = &cobra.Command{
	Use:   "stage [question]",
	Short: "Stage a Kubernetes exam question",
	Long: `This command stages a Kubernetes exam question, executing the relevant
logic for the specified question. Example usage:

  kr stage q1    # Runs the logic for question 1
  kr stage q2    # Runs the logic for question 2
  kr stage q3    # Runs the logic for question 3
`,
	Args: cobra.ExactArgs(1), // Ensure exactly one argument is provided
	Run: func(cmd *cobra.Command, args []string) {
		question := args[0]

		switch question {
		case "q1":
			questions.StageQ1()
		case "q2":
			questions.StageQ2()
		case "q3":
			questions.StageQ3()
		default:
			fmt.Println("Invalid question. Please use q1, q2, or q3.")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(stageCmd)
}
