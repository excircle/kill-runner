package questions

import (
	"fmt"

	"github.com/excircle/kill-runner/pkg/cluster"
)

var marker string = "Q1"

// StageQ1 runs the logic for question 1
func StageQ1() {
	fmt.Printf("[%s] Staging Scenario 1...\n", marker)

	// Creates additional namespace to check for in scenario
	err := cluster.CreateNamespace("q1-namespace", marker)
	if err != nil {
		fmt.Println("Error setting up namespace:", err)
		return
	}

	fmt.Printf("[%s] Successfully staged Q1 scenario!\n", marker)
}

func UnstageQ1() {
	fmt.Printf("[%s] Unstaging Kubernetes Question 1...\n", marker)

	// Undo Q1 scenario
	err := cluster.DestroyNamespace("q1-namespace", marker)
	if err != nil {
		fmt.Println("Error deleting namespace:", err)
		return
	}

	fmt.Printf("[%s] Successfully unstaged %s scenario!\n", marker, marker)
}
