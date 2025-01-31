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
# Scenario 3 - Create A Single Job
#----------------------------------

%sCONTEXT%s:   Create a Job and Validate It's Output
%sOBJECTIVE%s: 

- Create a Job template located at %s
- This Job should run image %s and execute %s
- It should run a total of %s and should %s in parallel.
- Each pod created by the Job should have the name and label 'id' == %s.
`, red, reset, green, reset, highlight("Q3/job.yaml", "green"), highlight("busybox:1.31.0", "green"), highlight("sleep 2 && echo 'done'", "green"), highlight("3 times", "green"), highlight("execute 2 runs", "green"), highlight("awesome-job", "green")),
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

	// Validate that job.yaml exists in Q3 dir
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get present working directory: %v", err)
	}

	jobPath := fmt.Sprintf("%s/%s/job.yaml", pwd, q3.marker)
	if _, err := os.Stat(jobPath); os.IsNotExist(err) {
		msg := fmt.Sprintf("[%s] %s does not exist!\n", q3.marker, q3.jobfile)
		fmt.Println(highlight(msg, "red"))
		os.Exit(0)
	} else {
		fmt.Printf("[%s] %s exists!\n", q3.marker, q3.jobfile)
	}

	// Validate job exists
	jobName := "awesome-job"
	if !cluster.CheckJobExists(namespace, jobName) {
		msg := fmt.Sprintf("[%s] Job %s does not exist!\n", q3.marker, jobName)
		fmt.Println(highlight(msg, "red"))
		os.Exit(0)
	} else {
		fmt.Printf("[%s] Job %s exists!\n", q3.marker, jobName)
	}

	// Validate job container image
	jobImage := "busybox:1.31.0"
	if !cluster.CheckContainerJobsForImage(namespace, "awesome-job", jobImage) {
		msg := fmt.Sprintf("[%s] Job does not contain image %s!\n", q3.marker, jobImage)
		fmt.Println(highlight(msg, "red"))
		os.Exit(0)
	} else {
		fmt.Printf("[%s] Job contains image %s!\n", q3.marker, jobImage)
	}

	// Validate job spec
	parallelism := int32(2)
	completions := int32(3)
	label := "awesome-job"
	jobStatus, issue := cluster.CheckJobSpec(namespace, jobName, parallelism, completions, label)
	if !jobStatus {
		msg := fmt.Sprintf("[%s] Job spec does not match! Issue with %s\n", q3.marker, issue)
		fmt.Println(highlight(msg, "red"))
		os.Exit(0)
	} else {
		fmt.Printf("[%s] Job spec matches!\n", q3.marker)
	}

	// Print validation complete
	success := fmt.Sprintf(`[%s] %sValidation complete!`, q3.marker, green)
	fmt.Println(success)
}
