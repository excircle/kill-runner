package questions

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/excircle/kill-runner/pkg/cluster"
)

type q7vars struct {
	marker   string
	scenario string
	skip     bool
	prompt   string
}

var q7 = q7vars{
	marker:   "Q7",
	scenario: "7",
	skip:     false,
	prompt: fmt.Sprintf(`#----------------------------------
# Scenario 7 - DaemonSet on all Nodes
#----------------------------------

%sCONTEXT%s:   Create a DaemonSet that runs on all nodes
%sOBJECTIVE 1%s: 

- Create a DaemonSet named %s
- Ensure it uses the image %s
- Ensure it has %s and %s.
- The Pods it creates should request %s and %s.
- These pods must run on %s.

`, red, reset, green, reset, highlight("ds-important", "green"), highlight("httpd:2.4-alpine", "green"), highlight("id=ds-important", "green"), highlight("uuid=18426a0b-5f59-4e10-923f-c0e078e82462", "green"), highlight("10 millicore cpu", "green"), highlight("10 mebibyte memory", "green"), highlight("all control planes and worker nodes", "green")),
}

// Stageq7 runs the logic for question 3
func StageQ7() {
	fmt.Printf("[%s] Staging Scenario %s...\n", q7.marker, q7.scenario)

	// Check if q7 dir exists
	if _, err := os.Stat(q7.marker); !os.IsNotExist(err) {
		q7.skip = true
	}

	if !q7.skip {
		pwd, err := os.Getwd()
		if err != nil {
			log.Fatalf("Failed to get present working directory: %v", err)
		}

		folderPath := fmt.Sprintf("%s/%s", pwd, q7.marker)

		err = os.Mkdir(folderPath, 0775) // Set perms
		if err != nil {
			log.Fatalf("Failed to create folder: %v", err)
		}
	}

	fmt.Printf("[%s] Successfully created %s dir!\n", q7.marker, q7.marker)

	// Check if q7 Namespace exists
	namespace := fmt.Sprintf("q%s-ns", q7.scenario)
	if cluster.NamespaceExists(namespace) {
		fmt.Printf("[%s] Namespace %s already exists!\n", q7.marker, namespace)
		q7.skip = true
	} else {
		fmt.Printf("[%s] Namespace does not exist!\n", q7.marker)
		q7.skip = false
	}

	if !q7.skip {
		// Create q7 Namespace
		err := cluster.CreateNamespace(namespace, q7.marker)
		if err != nil {
			fmt.Println("Error setting up namespace:", err)
			return
		}

	}

	fmt.Printf("[%s] Successfully staged %s scenario!\n", q7.marker, q7.marker)
	fmt.Printf("[%s] Please run 'kr start %s'!\n", q7.marker, strings.ToLower(q7.marker))

}

func StartQ7() {
	fmt.Println(q7.prompt)
}

func UnstageQ7() {
	fmt.Printf("[%s] Unstaging Kubernetes Question %s...\n", q7.marker, q7.scenario)

	// Check if q7 Namespace exists
	namespace := fmt.Sprintf("q%s-ns", q7.scenario)
	if !cluster.NamespaceExists(namespace) {
		fmt.Printf("[%s] Namespace %s does not exist!\n", q7.marker, namespace)
	} else {
		// Remove q7 Namespace
		fmt.Printf("[%s] Deleting %s!\n", q7.marker, q7.marker)
		err := cluster.DestroyNamespace(namespace, q7.marker)
		if err != nil {
			log.Fatalf("Error deleting namespace:", err)
		}
	}

	// Remove q7 dir
	if _, err := os.Stat(q7.marker); os.IsNotExist(err) {
		fmt.Printf("[%s] Answer directory is already gone!\n", q7.marker)
	} else {
		err = os.RemoveAll(q7.marker)
		if err != nil {
			log.Fatalf("Failed to remove %s dir: %v", err, q7.marker)
		}
	}

	fmt.Printf("[%s] Successfully unstaged %s scenario!\n", q7.marker, q7.marker)
}

func ValidateQ7() {
	fmt.Printf("[%s] Validating Kubernetes Question %s...\n", q7.marker, q7.scenario)

	// Validate that the namespace exists
	namespace := fmt.Sprintf("q%s-ns", q7.scenario)
	if !cluster.NamespaceExists(namespace) {
		msg := fmt.Sprintf("[%s] Namespace %s does not exist!\n", q7.marker, namespace)
		fmt.Println(highlight(msg, "red"))
		os.Exit(0)
	} else {
		fmt.Printf("[%s] Namespace %s exists!\n", q7.marker, namespace)
	}

	// Validate that the DaemonSet exists
	daemonset := "ds-important"
	if !cluster.DaemonSetExists(namespace, daemonset) {
		msg := fmt.Sprintf("[%s] DaemonSet %s does not exist!\n", q7.marker, daemonset)
		fmt.Println(highlight(msg, "red"))
		os.Exit(0)
	} else {
		fmt.Printf("[%s] DaemonSet %s exists!\n", q7.marker, daemonset)
	}

	// Validate that the DaemonSet uses the correct image
	image := "httpd:2.4-alpine"
	if !cluster.DaemonSetUsesImage(namespace, daemonset, image) {
		msg := fmt.Sprintf("[%s] DaemonSet %s does not use the correct image!\n", q7.marker, daemonset)
		fmt.Println(highlight(msg, "red"))
		os.Exit(0)
	} else {
		fmt.Printf("[%s] DaemonSet %s uses the correct image!\n", q7.marker, daemonset)
	}

	// Validate that the DaemonSet has the correct labels
	labelKey := "id"
	labelValue := "ds-important"
	if !cluster.DaemonSetHasLabel(namespace, daemonset, labelKey, labelValue) {
		msg := fmt.Sprintf("[%s] DaemonSet %s does not have the correct label!\n", q7.marker, daemonset)
		fmt.Println(highlight(msg, "red"))
		os.Exit(0)
	} else {
		fmt.Printf("[%s] DaemonSet %s has the correct label!\n", q7.marker, daemonset)
	}

	labelKey = "uuid"
	labelValue = "18426a0b-5f59-4e10-923f-c0e078e82462"
	if !cluster.DaemonSetHasLabel(namespace, daemonset, labelKey, labelValue) {
		msg := fmt.Sprintf("[%s] DaemonSet %s does not have the correct label!\n", q7.marker, daemonset)
		fmt.Println(highlight(msg, "red"))
		os.Exit(0)
	} else {
		fmt.Printf("[%s] DaemonSet %s has the correct label!\n", q7.marker, daemonset)
	}

	// Validate that the DaemonSet requests the correct resources
	cpuRequest := "10m"
	memoryRequest := "10Mi"
	if !cluster.DaemonSetHasResourceRequests(namespace, daemonset, cpuRequest, memoryRequest) {
		msg := fmt.Sprintf("[%s] DaemonSet %s does not request the correct resources!\n", q7.marker, daemonset)
		fmt.Println(highlight(msg, "red"))
		os.Exit(0)
	} else {
		fmt.Printf("[%s] DaemonSet %s requests the correct resources!\n", q7.marker, daemonset)
	}

	// Validate that the DaemonSet runs on all nodes
	if !cluster.DaemonSetRunningOnAllNodes(namespace, daemonset) {
		msg := fmt.Sprintf("[%s] DaemonSet %s does not run on all nodes!\n", q7.marker, daemonset)
		fmt.Println(highlight(msg, "red"))
		os.Exit(0)
	} else {
		fmt.Printf("[%s] DaemonSet %s runs on all nodes!\n", q7.marker, daemonset)
	}

	// Print validation complete
	success := fmt.Sprintf(`[%s] %sValidation complete!`, q7.marker, green)
	fmt.Println(success)
}
