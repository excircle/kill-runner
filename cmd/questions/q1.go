package questions

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/excircle/kill-runner/pkg/cluster"
)

type q1vars struct {
	marker   string
	scenario string
	skip     bool
	prompt   string
}

var q1 = q1vars{
	marker:   "Q1",
	scenario: "1",
	skip:     false,
	prompt: fmt.Sprintf(`#----------------------------------
# Scenario 1 - Obtain Namespaces
#----------------------------------

%sCONTEXT%s:   The DevOps team would like to get the list of all Namespaces in the present working cluster.
%sOBJECTIVE%s: Get the list and save it to './q1/namespaces.txt'.
`, red, reset, green, reset),
}

// StageQ1 runs the logic for question 1
func StageQ1() {
	fmt.Printf("[%s] Staging Scenario 1...\n", q1.marker)

	// Check if Q1 dir exists
	if _, err := os.Stat(q1.marker); !os.IsNotExist(err) {
		q1.skip = true
	}

	if !q1.skip {
		pwd, err := os.Getwd()
		if err != nil {
			log.Fatalf("Failed to get present working directory: %v", err)
		}

		folderPath := fmt.Sprintf("%s/%s", pwd, q1.marker)

		err = os.Mkdir(folderPath, 0775) // Set perms
		if err != nil {
			log.Fatalf("Failed to create folder: %v", err)
		}
	}

	fmt.Printf("[%s] Successfully created Q1 dir!\n", q1.marker)

	// Check if Q1 Namespace exists
	namespace := fmt.Sprintf("q%s-ns", q1.scenario)
	if cluster.NamespaceExists(namespace) {
		fmt.Printf("[%s] Namespace %s already exists!\n", q1.marker, namespace)
		q1.skip = true
	} else {
		fmt.Printf("[%s] Namespace does not exist!\n", q1.marker)
		q1.skip = false
	}

	if !q1.skip {
		// Create Q1 Namespace
		err := cluster.CreateNamespace(namespace, q1.marker)
		if err != nil {
			fmt.Println("Error setting up namespace:", err)
			return
		}

	}
	fmt.Printf("[%s] Successfully staged Q1 scenario!\n", q1.marker)
	fmt.Printf("[%s] Please run 'kr start q1'!\n", q1.marker)

}

func StartQ1() {
	fmt.Println(q1.prompt)
}

func UnstageQ1() {
	fmt.Printf("[%s] Unstaging Kubernetes Question 1...\n", q1.marker)

	// Check if Q1 Namespace exists
	namespace := fmt.Sprintf("q%s-ns", q1.scenario)
	if !cluster.NamespaceExists(namespace) {
		fmt.Printf("[%s] Namespace %s does not exist!\n", q1.marker, namespace)
	} else {
		// Remove Q1 Namespace
		fmt.Printf("[%s] Deleting %s!\n", q1.marker, q1.marker)
		err := cluster.DestroyNamespace(namespace, q1.marker)
		if err != nil {
			log.Fatalf("Error deleting namespace:", err)
		}
	}

	// Remove Q1 dir
	if _, err := os.Stat(q1.marker); os.IsNotExist(err) {
		fmt.Printf("[%s] Answer directory is already gone!\n", q1.marker)
	} else {
		err = os.RemoveAll(q1.marker)
		if err != nil {
			log.Fatalf("Failed to remove Q1 dir: %v", err)
		}
	}

	fmt.Printf("[%s] Successfully unstaged %s scenario!\n", q1.marker, q1.marker)
}

func ValidateQ1() {
	fmt.Printf("[%s] Validating Kubernetes Question 1...\n", q1.marker)

	// Check if "Q1/namespaces.txt" exists
	if _, err := os.Stat(fmt.Sprintf("%s/namespaces.txt", q1.marker)); os.IsNotExist(err) {
		fmt.Printf("[%s] %s/namespaces.txt not found. TRY AGAIN!\n", q1.marker, q1.marker)
		os.Exit(0)
	} else {
		fmt.Printf("[%s] %s/namespaces.txt found!\n", q1.marker, q1.marker)
	}

	// Check if namespaces exist in the file
	namespaces := []string{
		"calico-apiserver", "calico-system", "default", "kube-node-lease",
		"kube-public", "kube-system", "local-path-storage", "q1-ns", "tigera-operator",
	}

	// Read the file into a map for quick lookup
	answerFile := fmt.Sprintf("%s/namespaces.txt", q1.marker)
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
	fmt.Printf("[%s] namespace check suceeded!\n", q1.marker)
	success := fmt.Sprintf(`[%s] %sValidation complete!`, q1.marker, green)
	fmt.Println(success)
}
