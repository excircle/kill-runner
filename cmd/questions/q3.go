package questions

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/excircle/kill-runner/pkg/cluster"
)

type q3vars struct {
	marker   string
	scenario string
	skip     bool
	prompt   string
	jobfile  string
}

var q3 = q3vars{
	marker:   "Q3",
	scenario: "3",
	skip:     false,
	prompt: fmt.Sprintf(`#----------------------------------
# Scenario 3 - Pod Ready if Service is reachable
#----------------------------------

%sCONTEXT%s:   Create some pods that implement liveness and readiness probes
%sOBJECTIVE 1%s: 

- Create a Pod called %s
- This Pod should be of image %s
- Configure a LivenessProbe which simply executes command %s.
- Configure a ReadinessProbe which does check if the url %s is reachable .
- Validate using %s for this

%s:

- Create a second pod called %s
- This Pod should be of image %s
- This Pod should have label %s

`, red, reset, green, reset, highlight("ready-if-service-ready", "green"), highlight("nginx:1.16.1-alpine", "green"), highlight("true", "green"), highlight("http://service-am-i-ready:80", "green"), highlight("wget -T2 -O- http://service-am-i-ready:80", "green"), highlight("OBJECTIVE 2:", "green"), highlight("am-i-ready", "green"), highlight("nginx:1.16.1-alpine", "green"), highlight("label", "green")),
	jobfile: "job.yaml",
}

// Stageq3 runs the logic for question 3
func StageQ3() {
	fmt.Printf("[%s] Staging Scenario %s...\n", q3.marker, q3.scenario)

	// Check if q3 dir exists
	if _, err := os.Stat(q3.marker); !os.IsNotExist(err) {
		q3.skip = true
	}

	if !q3.skip {
		pwd, err := os.Getwd()
		if err != nil {
			log.Fatalf("Failed to get present working directory: %v", err)
		}

		folderPath := fmt.Sprintf("%s/%s", pwd, q3.marker)

		err = os.Mkdir(folderPath, 0775) // Set perms
		if err != nil {
			log.Fatalf("Failed to create folder: %v", err)
		}
	}

	fmt.Printf("[%s] Successfully created %s dir!\n", q3.marker, q3.marker)

	// Check if q3 Namespace exists
	namespace := fmt.Sprintf("q%s-ns", q3.scenario)
	if cluster.NamespaceExists(namespace) {
		fmt.Printf("[%s] Namespace %s already exists!\n", q3.marker, namespace)
		q3.skip = true
	} else {
		fmt.Printf("[%s] Namespace does not exist!\n", q3.marker)
		q3.skip = false
	}

	if !q3.skip {
		// Create q3 Namespace
		err := cluster.CreateNamespace(namespace, q3.marker)
		if err != nil {
			fmt.Println("Error setting up namespace:", err)
			return
		}

	}

	// Create SVC that exposes the pod
	svcName := "service-am-i-ready"
	port := int32(80)
	targetPort := int32(80)
	exposeType := "NodePort"
	podName := "am-i-ready"
	if !cluster.CheckService(namespace, svcName) {
		fmt.Printf("[%s] Creating service %s!\n", q3.marker, svcName)
		err := cluster.ExposePod(namespace, podName, svcName, port, targetPort, exposeType)
		if err != nil {
			fmt.Println("Error creating service:", err)
			return
		}
	} else {
		fmt.Printf("[%s] Service %s exists!\n", q3.marker, svcName)
	}

	fmt.Printf("[%s] Successfully staged %s scenario!\n", q3.marker, q3.marker)
	fmt.Printf("[%s] Please run 'kr start %s'!\n", q3.marker, strings.ToLower(q3.marker))

}

func StartQ3() {
	fmt.Println(q3.prompt)
}

func UnstageQ3() {
	fmt.Printf("[%s] Unstaging Kubernetes Question %s...\n", q3.marker, q3.scenario)

	// Check if q3 Namespace exists
	namespace := fmt.Sprintf("q%s-ns", q3.scenario)
	if !cluster.NamespaceExists(namespace) {
		fmt.Printf("[%s] Namespace %s does not exist!\n", q3.marker, namespace)
	} else {
		// Remove q3 Namespace
		fmt.Printf("[%s] Deleting %s!\n", q3.marker, q3.marker)
		err := cluster.DestroyNamespace(namespace, q3.marker)
		if err != nil {
			log.Fatalf("Error deleting namespace:", err)
		}
	}

	// Remove q3 dir
	if _, err := os.Stat(q3.marker); os.IsNotExist(err) {
		fmt.Printf("[%s] Answer directory is already gone!\n", q3.marker)
	} else {
		err = os.RemoveAll(q3.marker)
		if err != nil {
			log.Fatalf("Failed to remove %s dir: %v", err, q3.marker)
		}
	}

	fmt.Printf("[%s] Successfully unstaged %s scenario!\n", q3.marker, q3.marker)
}

func ValidateQ3() {
	fmt.Printf("[%s] Validating Kubernetes Question %s...\n", q3.marker, q3.scenario)

	// Validate that the namespace exists
	namespace := fmt.Sprintf("q%s-ns", q3.scenario)
	if !cluster.NamespaceExists(namespace) {
		msg := fmt.Sprintf("[%s] Namespace %s does not exist!\n", q3.marker, namespace)
		fmt.Println(highlight(msg, "red"))
		os.Exit(0)
	} else {
		fmt.Printf("[%s] Namespace %s exists!\n", q3.marker, namespace)
	}

	// Validate that the service exists
	svcName := "service-am-i-ready"
	if !cluster.CheckService(namespace, svcName) {
		msg := fmt.Sprintf("[%s] Service %s does not exist!\n", q3.marker, svcName)
		fmt.Println(highlight(msg, "red"))
		os.Exit(0)
	} else {
		fmt.Printf("[%s] Service %s exists!\n", q3.marker, svcName)
	}

	// Validate that ready-if-service-ready pod exists
	podName := "ready-if-service-ready"
	if !cluster.PodExists(podName, namespace) {
		msg := fmt.Sprintf("[%s] Pod %s does not exist!\n", q3.marker, podName)
		fmt.Println(highlight(msg, "red"))
		os.Exit(0)
	} else {
		fmt.Printf("[%s] Pod %s exists!\n", q3.marker, podName)
	}

	// Validate that ready-if-service-ready pod is running
	if !cluster.CheckPodRunning(podName, namespace) {
		msg := fmt.Sprintf("[%s] Pod %s is not running!\n", q3.marker, podName)
		fmt.Println(highlight(msg, "red"))
		os.Exit(0)
	} else {
		fmt.Printf("[%s] Pod %s is running!\n", q3.marker, podName)
	}

	// Validate that am-i-ready pod exists
	podName = "am-i-ready"
	if !cluster.PodExists(podName, namespace) {
		msg := fmt.Sprintf("[%s] Pod %s does not exist!\n", q3.marker, podName)
		fmt.Println(highlight(msg, "red"))
		os.Exit(0)
	} else {
		fmt.Printf("[%s] Pod %s exists!\n", q3.marker, podName)
	}

	// Validate that am-i-ready pod is running
	if !cluster.CheckPodRunning(podName, namespace) {
		msg := fmt.Sprintf("[%s] Pod %s is not running!\n", q3.marker, podName)
		fmt.Println(highlight(msg, "red"))
		os.Exit(0)
	} else {
		fmt.Printf("[%s] Pod %s is running!\n", q3.marker, podName)
	}

	// Print validation complete
	success := fmt.Sprintf(`[%s] %sValidation complete!`, q3.marker, green)
	fmt.Println(success)
}
