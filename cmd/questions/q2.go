package questions

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/excircle/kill-runner/pkg/cluster"
)

type q2vars struct {
	marker   string
	scenario string
	skip     bool
	prompt   string
}

var q2 = q2vars{
	marker:   "Q2",
	scenario: "2",
	skip:     false,
	prompt: fmt.Sprintf(`#----------------------------------
# Scenario 2 - Scale Down Stateful Set
#----------------------------------

%sCONTEXT%s:   Management has asked you scale down pods in order to save resource
%sOBJECTIVE%s: 

- Find the deployment called %s
- Ensure that the deployment is running %s
`, red, reset, green, reset, highlight("o3db", "green"), highlight("1 replica", "green")),
}

// Stageq2 runs the logic for question 2
func StageQ2() {
	fmt.Printf("[%s] Staging Scenario %s...\n", q2.marker, q2.scenario)

	// Check if Q2 dir exists
	if _, err := os.Stat(q2.marker); !os.IsNotExist(err) {
		q2.skip = true
	}

	if !q2.skip {
		pwd, err := os.Getwd()
		if err != nil {
			log.Fatalf("Failed to get present working directory: %v", err)
		}

		folderPath := fmt.Sprintf("%s/%s", pwd, q2.marker)

		err = os.Mkdir(folderPath, 0775) // Set perms
		if err != nil {
			log.Fatalf("Failed to create folder: %v", err)
		}
	}

	fmt.Printf("[%s] Successfully created %s dir!\n", q2.marker, q2.marker)

	// Check if Q2 Namespace exists
	namespace := fmt.Sprintf("q%s-ns", q2.scenario)
	if cluster.NamespaceExists(namespace) {
		fmt.Printf("[%s] Namespace %s already exists!\n", q2.marker, namespace)
		q2.skip = true
	} else {
		fmt.Printf("[%s] Namespace does not exist!\n", q2.marker)
		q2.skip = false
	}

	if !q2.skip {
		// Create q2 Namespace
		err := cluster.CreateNamespace(namespace, q2.marker)
		if err != nil {
			fmt.Println("Error setting up namespace:", err)
			return
		}

	}

	// Create o3db deployment using CreateDeployment() func
	deploymentName := "o3db"
	deployImage := "nginx"
	replicas := int32(3)
	if cluster.CheckDeployment(namespace, deploymentName) {
		fmt.Printf("[%s] Deployment %s exists!\n", q2.marker, deploymentName)
	} else {
		fmt.Printf("[%s] Creating deployment %s!\n", q2.marker, deploymentName)
		err := cluster.CreateDeployment(namespace, deploymentName, deployImage, replicas)
		if err != nil {
			fmt.Println("Error creating deployment:", err)
			return
		}
	}

	fmt.Printf("[%s] Successfully staged %s scenario!\n", q2.marker, q2.marker)
	fmt.Printf("[%s] Please run 'kr start %s'!\n", q2.marker, strings.ToLower(q2.marker))

}

func StartQ2() {
	fmt.Println(q2.prompt)
}

func UnstageQ2() {
	fmt.Printf("[%s] Unstaging Kubernetes Question %s...\n", q2.marker, q2.scenario)

	// Check if Q2 Namespace exists
	namespace := fmt.Sprintf("q%s-ns", q2.scenario)
	if !cluster.NamespaceExists(namespace) {
		fmt.Printf("[%s] Namespace %s does not exist!\n", q2.marker, namespace)
	} else {
		// Remove Q2 Namespace
		fmt.Printf("[%s] Deleting %s!\n", q2.marker, q2.marker)
		err := cluster.DestroyNamespace(namespace, q2.marker)
		if err != nil {
			log.Fatalf("Error deleting namespace:", err)
		}
	}

	// Remove Q2 dir
	if _, err := os.Stat(q2.marker); os.IsNotExist(err) {
		fmt.Printf("[%s] Answer directory is already gone!\n", q2.marker)
	} else {
		err = os.RemoveAll(q2.marker)
		if err != nil {
			log.Fatalf("Failed to remove %s dir: %v", err, q2.marker)
		}
	}

	fmt.Printf("[%s] Successfully unstaged %s scenario!\n", q2.marker, q2.marker)
}

func ValidateQ2() {
	fmt.Printf("[%s] Validating Kubernetes Question %s...\n", q2.marker, q2.scenario)

	// Validate that the namespace exists
	namespace := fmt.Sprintf("q%s-ns", q2.scenario)
	if !cluster.NamespaceExists(namespace) {
		msg := fmt.Sprintf("[%s] Namespace %s does not exist!\n", q2.marker, namespace)
		fmt.Println(highlight(msg, "red"))
		os.Exit(0)
	} else {
		fmt.Printf("[%s] Namespace %s exists!\n", q2.marker, namespace)
	}

	// Validate that the deployment exists
	deploymentName := "o3db"
	if !cluster.CheckDeployment(namespace, deploymentName) {
		msg := fmt.Sprintf("[%s] Deployment %s does not exist!\n", q2.marker, deploymentName)
		fmt.Println(highlight(msg, "red"))
		os.Exit(0)
	} else {
		fmt.Printf("[%s] Deployment %s exists!\n", q2.marker, deploymentName)
	}

	// Validate that the deployment has 1 replica
	expectedReplicas := int32(1)
	if !cluster.CheckReplicaCount(namespace, deploymentName, expectedReplicas) {
		msg := fmt.Sprintf("[%s] Deployment %s does not have %d replicas!\n", q2.marker, deploymentName, expectedReplicas)
		fmt.Println(highlight(msg, "red"))
		os.Exit(0)
	} else {
		fmt.Printf("[%s] Deployment %s has %d replicas!\n", q2.marker, deploymentName, expectedReplicas)
	}

	// Print validation complete
	success := fmt.Sprintf(`[%s] %sValidation complete!`, q2.marker, green)
	fmt.Println(success)
}
