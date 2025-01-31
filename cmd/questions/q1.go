package questions

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/excircle/kill-runner/pkg/cluster"
)

var marker string = "Q1"
var red string = "\033[31m"   // ANSI escape code for red
var green string = "\033[32m" // ANSI escape code for green
var reset string = "\033[0m"  // Reset color
var namespace string = "q1-namespace"
var skip bool = false
var doNothing bool

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

	// Check if Q1 Namespace exists
	if !cluster.NamespaceExists(namespace) {
		fmt.Printf("[%s] Namespace %s does not exist!\n", marker, namespace)
	} else {
		// Remove Q1 Namespace
		fmt.Printf("[%s] Deleting %s!\n", marker, marker)
		err := cluster.DestroyNamespace(namespace, marker)
		if err != nil {
			log.Fatalf("Error deleting namespace:", err)
		}
	}

	// Remove Q1 dir
	if _, err := os.Stat(marker); os.IsNotExist(err) {
		fmt.Printf("[%s] Answer directory is already gone!\n", marker)
	} else {
		err = os.RemoveAll(marker)
		if err != nil {
			log.Fatalf("Failed to remove Q1 dir: %v", err)
		}
	}

	fmt.Printf("[%s] Successfully unstaged %s scenario!\n", marker, marker)
}

func ValidateQ1() {
	fmt.Printf("[%s] Validating Kubernetes Question 1...\n", marker)

	// Check if "Q1/namespaces.txt" exists
	if _, err := os.Stat(fmt.Sprintf("%s/namespaces.txt", marker)); os.IsNotExist(err) {
		fmt.Printf("[%s] %s/namespaces.txt not found. TRY AGAIN!\n", marker, marker)
		os.Exit(0)
	} else {
		fmt.Printf("[%s] %s/namespaces.txt found!\n", marker, marker)
	}

	// Check if namespaces exist in the file
	namespaces := []string{
		"calico-apiserver", "calico-system", "default", "kube-node-lease",
		"kube-public", "kube-system", "local-path-storage", "q1-namespace", "tigera-operator",
	}

	// Read the file into a map for quick lookup
	answerFile := fmt.Sprintf("%s/namespaces.txt", marker)
	file, err := os.Open(answerFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Search for terms
	for _, ns := range namespaces {
		found := false
		for _, line := range lines {
			if strings.Contains(line, ns) {
				found = true
			}
		}
		if !found {
			fmt.Printf("[FAILURE] namespace not found. TRY AGAIN!\n")
			os.Exit(0)
		}
	}
	// Print validation complete
	fmt.Printf("[%s] namespace check suceeded!\n", marker)
	fmt.Printf("[%s] Validation complete!\n", marker)
}
