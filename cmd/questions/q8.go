package questions

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/excircle/kill-runner/pkg/cluster"
)

type q8vars struct {
	marker   string
	scenario string
	skip     bool
	prompt   string
}

var q8 = q8vars{
	marker:   "Q8",
	scenario: "8",
	skip:     false,
	prompt: fmt.Sprintf(`#----------------------------------
# Scenario 8 - Multi Containers and Pod shared Volume
#----------------------------------

%sCONTEXT%s:   Create a pod that runs multiple containers and shares a volume between them.
%sOBJECTIVE 1%s: 

- Create a Pod named %s
- The Pod should have three containers: %s
- Ensure volume %s is attached to every container inside of pod %s 
- Container c1 should be of image %s
- Container c1 should have an environment named %s with a value reflecting the node name it is on
- Container c2 should be of image %s
- Container c3 should be of image %s
- The command %s should be written by c2 to the shared volume

`, red, reset, green, reset, highlight("multi-container-playground", "green"), highlight("c1, c2, and c3", "green"), highlight("myvol", "green"), highlight("multi-container-playground", "green"), highlight("nginx:1.17.6-alpine", "green"), highlight("MY_NODE_NAME", "green"), highlight("redis:7.4.2-bookworm", "green"), highlight("httpd:2.4.41-alpine", "green"), highlight("while true; do date >> /your/vol/path/date.log; sleep 1;", "green")),
}

// Stageq8 runs the logic for question 8
func StageQ8() {
	fmt.Printf("[%s] Staging Scenario %s...\n", q8.marker, q8.scenario)

	// Check if q8 dir exists
	if _, err := os.Stat(q8.marker); !os.IsNotExist(err) {
		q8.skip = true
	}

	if !q8.skip {
		pwd, err := os.Getwd()
		if err != nil {
			log.Fatalf("Failed to get present working directory: %v", err)
		}

		folderPath := fmt.Sprintf("%s/%s", pwd, q8.marker)

		err = os.Mkdir(folderPath, 0775) // Set perms
		if err != nil {
			log.Fatalf("Failed to create folder: %v", err)
		}
	}

	fmt.Printf("[%s] Successfully created %s dir!\n", q8.marker, q8.marker)

	// Check if q8 Namespace exists
	namespace := fmt.Sprintf("q%s-ns", q8.scenario)
	if cluster.NamespaceExists(namespace) {
		fmt.Printf("[%s] Namespace %s already exists!\n", q8.marker, namespace)
		q8.skip = true
	} else {
		fmt.Printf("[%s] Namespace does not exist!\n", q8.marker)
		q8.skip = false
	}

	if !q8.skip {
		// Create q8 Namespace
		err := cluster.CreateNamespace(namespace, q8.marker)
		if err != nil {
			fmt.Println("Error setting up namespace:", err)
			return
		}

	}

	fmt.Printf("[%s] Successfully staged %s scenario!\n", q8.marker, q8.marker)
	fmt.Printf("[%s] Please run 'kr start %s'!\n", q8.marker, strings.ToLower(q8.marker))

}

func StartQ8() {
	fmt.Println(q8.prompt)
}

func UnstageQ8() {
	fmt.Printf("[%s] Unstaging Kubernetes Question %s...\n", q8.marker, q8.scenario)

	// Check if q8 Namespace exists
	namespace := fmt.Sprintf("q%s-ns", q8.scenario)
	if !cluster.NamespaceExists(namespace) {
		fmt.Printf("[%s] Namespace %s does not exist!\n", q8.marker, namespace)
	} else {
		// Remove q8 Namespace
		fmt.Printf("[%s] Deleting %s!\n", q8.marker, q8.marker)
		err := cluster.DestroyNamespace(namespace, q8.marker)
		if err != nil {
			log.Fatalf("Error deleting namespace:", err)
		}
	}

	// Remove q8 dir
	if _, err := os.Stat(q8.marker); os.IsNotExist(err) {
		fmt.Printf("[%s] Answer directory is already gone!\n", q8.marker)
	} else {
		err = os.RemoveAll(q8.marker)
		if err != nil {
			log.Fatalf("Failed to remove %s dir: %v", err, q8.marker)
		}
	}

	fmt.Printf("[%s] Successfully unstaged %s scenario!\n", q8.marker, q8.marker)
}

func ValidateQ8() {
	fmt.Printf("[%s] Validating Kubernetes Question %s...\n", q8.marker, q8.scenario)

	// Validate that the namespace exists
	namespace := fmt.Sprintf("q%s-ns", q8.scenario)
	if !cluster.NamespaceExists(namespace) {
		msg := fmt.Sprintf("[%s] Namespace %s does not exist!\n", q8.marker, namespace)
		fmt.Println(highlight(msg, "red"))
		os.Exit(0)
	} else {
		fmt.Printf("[%s] Namespace %s exists!\n", q8.marker, namespace)
	}

	// Validate that multi-container-playground Pod exists
	pod := "multi-container-playground"
	if !cluster.PodExists(pod, namespace) {
		msg := fmt.Sprintf("[%s] Pod %s does not exist!\n", q8.marker, pod)
		fmt.Println(highlight(msg, "red"))
		os.Exit(0)
	} else {
		fmt.Printf("[%s] Pod %s exists!\n", q8.marker, pod)
	}

	// Check multi-container-playground for 'c1' container
	container := "c1"
	if !cluster.CheckPodForContainer(pod, container, namespace) {
		msg := fmt.Sprintf("[%s] Container %s does not exist in Pod %s!\n", q8.marker, container, pod)
		fmt.Println(highlight(msg, "red"))
		os.Exit(0)
	} else {
		fmt.Printf("[%s] Container %s exists in Pod %s!\n", q8.marker, container, pod)
	}

	// Check multi-container-playground for 'c2' container
	container = "c2"
	if !cluster.CheckPodForContainer(pod, container, namespace) {
		msg := fmt.Sprintf("[%s] Container %s does not exist in Pod %s!\n", q8.marker, container, pod)
		fmt.Println(highlight(msg, "red"))
		os.Exit(0)
	} else {
		fmt.Printf("[%s] Container %s exists in Pod %s!\n", q8.marker, container, pod)
	}

	// Check multi-container-playground for 'c3' container
	container = "c3"
	if !cluster.CheckPodForContainer(pod, container, namespace) {
		msg := fmt.Sprintf("[%s] Container %s does not exist in Pod %s!\n", q8.marker, container, pod)
		fmt.Println(highlight(msg, "red"))
		os.Exit(0)
	} else {
		fmt.Printf("[%s] Container %s exists in Pod %s!\n", q8.marker, container, pod)
	}

	// Ensure all containers inside of pod multi-container-playground have volume myvol attached
	volume := "myvol"
	for _, container := range []string{"c1", "c2", "c3"} {
		if !cluster.CheckContainerForVolume(namespace, pod, container, volume) {
			msg := fmt.Sprintf("[%s] Container %s does not have volume %s attached!\n", q8.marker, container, volume)
			fmt.Println(highlight(msg, "red"))
			os.Exit(0)
		} else {
			fmt.Printf("[%s] Container %s has volume %s attached!\n", q8.marker, container, volume)
		}
	}

	// Ensure container c1 is of image nginx:1.17.6-alpine
	container = "c1"
	image := "nginx:1.17.6-alpine"
	if !cluster.CheckContainerForImage(pod, container, namespace, image) {
		msg := fmt.Sprintf("[%s] Container %s is not of image %s!\n", q8.marker, container, image)
		fmt.Println(highlight(msg, "red"))
		os.Exit(0)
	} else {
		fmt.Printf("[%s] Container %s is of image %s!\n", q8.marker, container, image)
	}

	// Ensure container c2 is of image 7.4.2-bookworm
	container = "c2"
	image = "redis:7.4.2-bookworm"
	if !cluster.CheckContainerForImage(pod, container, namespace, image) {
		msg := fmt.Sprintf("[%s] Container %s is not of image %s!\n", q8.marker, container, image)
		fmt.Println(highlight(msg, "red"))
		os.Exit(0)
	} else {
		fmt.Printf("[%s] Container %s is of image %s!\n", q8.marker, container, image)
	}

	// Ensure container c3 is of image httpd:2.4.41-alpine
	container = "c3"
	image = "httpd:2.4.41-alpine"
	if !cluster.CheckContainerForImage(pod, container, namespace, image) {
		msg := fmt.Sprintf("[%s] Container %s is not of image %s!\n", q8.marker, container, image)
		fmt.Println(highlight(msg, "red"))
		os.Exit(0)
	} else {
		fmt.Printf("[%s] Container %s is of image %s!\n", q8.marker, container, image)
	}

	// Print validation complete
	success := fmt.Sprintf(`[%s] %sValidation complete!`, q8.marker, green)
	fmt.Println(success)
}
