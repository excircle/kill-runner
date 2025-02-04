package questions

import (
	"fmt"
	"log"
	"os"

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
# Scenario 1 - Schedule Pod On Control Plane Node
#----------------------------------

%sCONTEXT%s:   Demonstrate your ability to schedule a Kubernetes pod on a control plane node.
%sOBJECTIVE%s: 
- Create a single Pod of image %s
- The pod should be called %s
- The container should be called %s
- This pod must be on the %s
`, red, reset, green, reset, highlight("httpd:2.4.41-alpine", "green"), highlight("pod1", "green"), highlight("pod1-container", "green"), highlight("control plane node", "green")),
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

	// Check that 'pod1' exists
	if !cluster.PodExists("pod1", "q1-ns") {
		fmt.Printf("[%s] Pod 'pod1' does not exist!\n", q1.marker)
		os.Exit(1)
	} else {
		fmt.Printf("[%s] Pod 'pod1' exists!\n", q1.marker)
	}

	// Check that pod1 is running httpd:2.4.41-alpine
	if !cluster.CheckContainerForImage("pod1", "pod1-container", "q1-ns", "httpd:2.4.41-alpine") {
		fmt.Printf("[%s] Pod 'pod1' is not running httpd:2.4.41-alpine!\n", q1.marker)
		os.Exit(1)
	} else {
		fmt.Printf("[%s] Pod 'pod1' is running httpd:2.4.41-alpine!\n", q1.marker)
	}

	// Check pod1 for container 'pod1-container'
	if !cluster.CheckPodForContainer("pod1", "pod1-container", "q1-ns") {
		fmt.Printf("[%s] Pod 'pod1' does not have container 'pod1-container'!\n", q1.marker)
		os.Exit(1)
	} else {
		fmt.Printf("[%s] Pod 'pod1' has container 'pod1-container'!\n", q1.marker)
	}
	// Check that pod1 is on the control plane node
	if node, ok := cluster.CheckPodsNode("q1-ns", "pod1"); !ok {
		fmt.Printf("[%s] Pod 'pod1' is not on the control plane node!\n", q1.marker)
		os.Exit(1)
	} else {
		fmt.Printf("[%s] Pod 'pod1' is on the control plane node %s!\n", q1.marker, node)
	}

	// Print validation complete
	fmt.Printf("[%s] namespace check suceeded!\n", q1.marker)
	success := fmt.Sprintf(`[%s] %sValidation complete!`, q1.marker, green)
	fmt.Println(success)
}
