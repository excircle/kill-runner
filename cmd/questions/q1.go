package questions

import (
	"fmt"
)

// StageQ1 runs the logic for question 1
func StageQ1() {
	fmt.Println("Staging Kubernetes Question 1...")

	// Example: Interacting with Kubernetes
	// err := cluster.SetupNamespace("q1-namespace")
	// if err != nil {
	// 	fmt.Println("Error setting up namespace:", err)
	// 	return
	// }

	fmt.Println("Successfully staged q1!")
}

func UnstageQ1() {
	fmt.Println("Unstaging Kubernetes Question 1...")

	// Example: Interacting with Kubernetes
	// err := cluster.DeleteNamespace("q1-namespace")
	// if err != nil {
	// 	fmt.Println("Error deleting namespace:", err)
	// 	return
	// }

	fmt.Println("Successfully unstaged q1!")
}
