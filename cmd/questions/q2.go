package questions

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/excircle/kill-runner/pkg/cluster"
	"github.com/excircle/kill-runner/pkg/utils"
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
# Scenario 2 - Create A Single Pod
#----------------------------------

%sCONTEXT%s:   Create a Pod and create a BASH script to assist your manager check data
%sOBJECTIVE%s: 

- Create a single Pod of image %s in Namespace '%s'.
- The Pod should be named '%s' and the container should be named '%s'.
- Write a command that provides the status of this pod  into './Q2/pod1-status-command.sh'. The command should use kubectl.
`, red, reset, green, reset, highlight("httpd:2.4.41-alpine", "green"), highlight("q2-ns", "green"), highlight("pod1", "green"), highlight("pod1-container", "green")),
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

	// Validate that pod1 pod exists
	podName := "pod1"
	if !cluster.PodExists(podName, namespace) {
		msg := fmt.Sprintf("[%s] Pod %s does not exist!\n", q2.marker, podName)
		fmt.Println(highlight(msg, "red"))
		os.Exit(0)
	} else {
		fmt.Printf("[%s] Pod %s exists!\n", q2.marker, podName)
	}

	// Check that pod1 has a container called pod1-container
	containerName := "pod1-container"
	if !cluster.CheckPodForContainer(podName, containerName, namespace) {
		msg := fmt.Sprintf("[%s] Container %s does not exist in Pod %s!\n", q2.marker, containerName, podName)
		fmt.Println(highlight(msg, "red"))
		os.Exit(0)
	} else {
		fmt.Printf("[%s] Container %s exists in Pod %s!\n", q2.marker, containerName, podName)
	}

	// Check that pod1-container is running the httpd:2.4.41-alpine image
	imageName := "httpd:2.4.41-alpine"
	if !cluster.CheckContainerForImage(podName, containerName, namespace, imageName) {
		msg := fmt.Sprintf("[%s] Container %s in Pod %s is not running image %s!\n", q2.marker, containerName, podName, imageName)
		fmt.Println(highlight(msg, "red"))
		os.Exit(0)
	} else {
		fmt.Printf("[%s] Container %s in Pod %s is running image %s!\n", q2.marker, containerName, podName, imageName)
	}

	// Check that file ./Q2/pod1-status-command.sh contains string "Running"
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get present working directory: %v", err)
	}

	filePath := fmt.Sprintf("%s/%s/pod1-status-command.sh", pwd, q2.marker)
	if !utils.FileContainsString(filePath, "Running") {
		msg := fmt.Sprintf("[%s] %s does not contain string 'Running'!\n", q2.marker, filePath)
		fmt.Println(highlight(msg, "red"))
		os.Exit(0)
	} else {
		fmt.Printf("[%s] %s contains string 'Running'!\n", q2.marker, filePath)
	}

	// Print validation complete
	success := fmt.Sprintf(`[%s] %sValidation complete!`, q2.marker, green)
	fmt.Println(success)
}
