package killrunner

import (
	"fmt"
	"os"

	"github.com/excircle/kill-runner/cmd/questions" // Import the questions package

	"github.com/spf13/cobra"
)

// stageCmd represents the stage command
var validateCmd = &cobra.Command{
	Use:   "validate [question]",
	Short: "validate a Kubernetes exam question",
	Long: `This command validates a Kubernetes exam question, determining if the answer is correct. Example usage:

  kr validate q1    # Provide the prompt for question 1
  kr validate q2    # Provide the prompt for question 2
  kr validate q3    # Provide the prompt for question 3
`,
	Args: cobra.ExactArgs(1), // Ensure exactly one argument is provided
	Run: func(cmd *cobra.Command, args []string) {
		question := args[0]

		switch question {
		case "q1":
			questions.ValidateQ1()
		case "q2":
			questions.ValidateQ2()
		case "q3":
			questions.ValidateQ3()
		case "q4":
			questions.ValidateQ4()
		case "q6":
			questions.ValidateQ6()
		case "q7":
			questions.ValidateQ7()
		case "q8":
			questions.ValidateQ8()
		default:
			fmt.Println("Invalid question. Please use q1, q2, or q3.")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
}
