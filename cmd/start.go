package killrunner

import (
	"fmt"
	"os"

	"github.com/excircle/kill-runner/cmd/questions" // Import the questions package

	"github.com/spf13/cobra"
)

// stageCmd represents the stage command
var startCmd = &cobra.Command{
	Use:   "start [question]",
	Short: "start a Kubernetes exam question",
	Long: `This command starts a Kubernetes exam question, providing the prompt required 
	to complete the question. Example usage:

  kr start q1    # Provide the prompt for question 1
  kr start q2    # Provide the prompt for question 2
  kr start q3    # Provide the prompt for question 3
`,
	Args: cobra.ExactArgs(1), // Ensure exactly one argument is provided
	Run: func(cmd *cobra.Command, args []string) {
		question := args[0]

		switch question {
		case "q1":
			questions.StartQ1()
		case "q2":
			questions.StartQ2()
		case "q3":
			questions.StartQ3()
		case "q4":
			questions.StartQ4()
		case "q6":
			questions.StartQ6()
		case "q7":
			questions.StartQ7()
		case "q8":
			questions.StartQ8()
		default:
			fmt.Println("Invalid question. Please use q1, q2, or q3.")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
