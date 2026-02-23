package questions

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/excircle/kill-runner/pkg/cluster"
)

type q5vars struct {
	marker   string
	scenario string
	skip     bool
	prompt   string
	jobfile  string
}

var q5 = q5vars{
	marker:   "Q5",
	scenario: "5",
	skip:     false,
	prompt: fmt.Sprintf(`#--------------------------------------------
# Scenario 3 - Sort Pods By Creation Date
#--------------------------------------------

%sCONTEXT%s:   Management wants a txt file that contains a sorted list of pods by creation date
%sOBJECTIVE 1%s: 

- Using kubectl, create a file called %s inside of the %s directory

`, red, reset, green, reset, highlight("pod-ages.txt", "green"), highlight("Q5", "green")),
	jobfile: "job.yaml",
}

// Stageq5 runs the logic for question 3
func StageQ5() {
	fmt.Printf("[%s] Staging Scenario %s...\n", q5.marker, q5.scenario)

	// Check if q5 dir exists
	if _, err := os.Stat(q5.marker); !os.IsNotExist(err) {
		q5.skip = true
	}

	if !q5.skip {
		pwd, err := os.Getwd()
		if err != nil {
			log.Fatalf("Failed to get present working directory: %v", err)
		}

		folderPath := fmt.Sprintf("%s/%s", pwd, q5.marker)

		err = os.Mkdir(folderPath, 0775) // Set perms
		if err != nil {
			log.Fatalf("Failed to create folder: %v", err)
		}
	}

	fmt.Printf("[%s] Successfully created %s dir!\n", q5.marker, q5.marker)

	// Check if q5 Namespace exists
	namespace := fmt.Sprintf("q%s-ns", q5.scenario)
	if cluster.NamespaceExists(namespace) {
		fmt.Printf("[%s] Namespace %s already exists!\n", q5.marker, namespace)
		q5.skip = true
	} else {
		fmt.Printf("[%s] Namespace does not exist!\n", q5.marker)
		q5.skip = false
	}

	if !q5.skip {
		// Create q5 Namespace
		err := cluster.CreateNamespace(namespace, q5.marker)
		if err != nil {
			fmt.Println("Error setting up namespace:", err)
			return
		}

	}

	fmt.Printf("[%s] Successfully staged %s scenario!\n", q5.marker, q5.marker)
	fmt.Printf("[%s] Please run 'kr start %s'!\n", q5.marker, strings.ToLower(q5.marker))

}

func StartQ5() {
	fmt.Println(q5.prompt)
}

func UnstageQ5() {
	fmt.Printf("[%s] Unstaging Kubernetes Question %s...\n", q5.marker, q5.scenario)

	// Check if q5 Namespace exists
	namespace := fmt.Sprintf("q%s-ns", q5.scenario)
	if !cluster.NamespaceExists(namespace) {
		fmt.Printf("[%s] Namespace %s does not exist!\n", q5.marker, namespace)
	} else {
		// Remove q5 Namespace
		fmt.Printf("[%s] Deleting %s!\n", q5.marker, q5.marker)
		err := cluster.DestroyNamespace(namespace, q5.marker)
		if err != nil {
			log.Fatalf("Error deleting namespace:", err)
		}
	}

	// Remove q5 dir
	if _, err := os.Stat(q5.marker); os.IsNotExist(err) {
		fmt.Printf("[%s] Answer directory is already gone!\n", q5.marker)
	} else {
		err = os.RemoveAll(q5.marker)
		if err != nil {
			log.Fatalf("Failed to remove %s dir: %v", err, q5.marker)
		}
	}

	fmt.Printf("[%s] Successfully unstaged %s scenario!\n", q5.marker, q5.marker)
}

func ValidateQ5() {
	fmt.Printf("[%s] Validating Kubernetes Question %s...\n", q5.marker, q5.scenario)

	// Validate that the namespace exists
	namespace := fmt.Sprintf("q%s-ns", q5.scenario)
	if !cluster.NamespaceExists(namespace) {
		msg := fmt.Sprintf("[%s] Namespace %s does not exist!\n", q5.marker, namespace)
		fmt.Println(highlight(msg, "red"))
		os.Exit(0)
	} else {
		fmt.Printf("[%s] Namespace %s exists!\n", q5.marker, namespace)
	}

	// Validate that the service exists
	svcName := "service-am-i-ready"
	if !cluster.CheckService(namespace, svcName) {
		msg := fmt.Sprintf("[%s] Service %s does not exist!\n", q5.marker, svcName)
		fmt.Println(highlight(msg, "red"))
		os.Exit(0)
	} else {
		fmt.Printf("[%s] Service %s exists!\n", q5.marker, svcName)
	}

	// Validate that ready-if-service-ready pod exists
	podName := "ready-if-service-ready"
	if !cluster.PodExists(podName, namespace) {
		msg := fmt.Sprintf("[%s] Pod %s does not exist!\n", q5.marker, podName)
		fmt.Println(highlight(msg, "red"))
		os.Exit(0)
	} else {
		fmt.Printf("[%s] Pod %s exists!\n", q5.marker, podName)
	}

	// Validate that ready-if-service-ready pod is running
	if !cluster.CheckPodRunning(podName, namespace) {
		msg := fmt.Sprintf("[%s] Pod %s is not running!\n", q5.marker, podName)
		fmt.Println(highlight(msg, "red"))
		os.Exit(0)
	} else {
		fmt.Printf("[%s] Pod %s is running!\n", q5.marker, podName)
	}

	// Validate that am-i-ready pod exists
	podName = "am-i-ready"
	if !cluster.PodExists(podName, namespace) {
		msg := fmt.Sprintf("[%s] Pod %s does not exist!\n", q5.marker, podName)
		fmt.Println(highlight(msg, "red"))
		os.Exit(0)
	} else {
		fmt.Printf("[%s] Pod %s exists!\n", q5.marker, podName)
	}

	// Validate that am-i-ready pod is running
	if !cluster.CheckPodRunning(podName, namespace) {
		msg := fmt.Sprintf("[%s] Pod %s is not running!\n", q5.marker, podName)
		fmt.Println(highlight(msg, "red"))
		os.Exit(0)
	} else {
		fmt.Printf("[%s] Pod %s is running!\n", q5.marker, podName)
	}

	// Print validation complete
	success := fmt.Sprintf(`[%s] %sValidation complete!`, q5.marker, green)
	fmt.Println(success)
}
