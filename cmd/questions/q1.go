package questions

import (
	"fmt"
	"log"
	"os"

	"github.com/excircle/kill-runner/pkg/cluster"
)

var marker string = "Q1"
var red string = "\033[31m"   // ANSI escape code for red
var green string = "\033[32m" // ANSI escape code for green
var reset string = "\033[0m"  // Reset color
var namespace string = "q1-namespace"
var skip bool = false

var prompt string = fmt.Sprintf(`#----------------------------------
# Scenario 1 - Obtain Namespaces
#----------------------------------

%sCONTEXT%s:   The DevOps team would like to get the list of all Namespaces in the present working cluster.
%sOBJECTIVE%s: Get the list and save it to './q1/namespaces.txt'.
`, red, reset, green, reset)

// StageQ1 runs the logic for question 1
func StageQ1() {
	fmt.Printf("[%s] Staging Scenario 1...\n", marker)

	// Check if Q1 dir exists
	if _, err := os.Stat(marker); !os.IsNotExist(err) {
		skip = true
	}

	if !skip {
		pwd, err := os.Getwd()
		if err != nil {
			log.Fatalf("Failed to get present working directory: %v", err)
		}

		folderPath := fmt.Sprintf("%s/%s", pwd, marker)

		err = os.Mkdir(folderPath, 0775) // Set perms
		if err != nil {
			log.Fatalf("Failed to create folder: %v", err)
		}
	}

	fmt.Printf("[%s] Successfully created Q1 dir!\n", marker)

	// Check if Q1 Namespace exists
	if cluster.NamespaceExists(namespace) {
		fmt.Printf("[%s] Namespace %s already exists!\n", marker, namespace)
		skip = true
	} else {
		fmt.Printf("[%s] Namespace does not exist!\n", marker)
		skip = false
	}

	if !skip {
		// Create Q1 Namespace
		err := cluster.CreateNamespace(namespace, marker)
		if err != nil {
			fmt.Println("Error setting up namespace:", err)
			return
		}

	}
	fmt.Printf("[%s] Successfully staged Q1 scenario!\n", marker)
	fmt.Printf("[%s] Please run 'kr start q1'!\n", marker)

}

func StartQ1() {
	fmt.Println(prompt)
}

func UnstageQ1() {
	fmt.Printf("[%s] Unstaging Kubernetes Question 1...\n", marker)

	// Undo Q1 Namespace
	fmt.Printf("[%s] Deleting %s!\n", marker, marker)
	err := cluster.DestroyNamespace(namespace, marker)
	if err != nil {
		fmt.Println("Error deleting namespace:", err)
		return
	}

	// Remove Q1 dir
	err = os.RemoveAll(marker)
	if err != nil {
		log.Fatalf("Failed to remove Q1 dir: %v", err)
	}

	fmt.Printf("[%s] Successfully unstaged %s scenario!\n", marker, marker)
}
