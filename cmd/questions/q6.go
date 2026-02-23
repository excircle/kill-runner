package questions

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/excircle/kill-runner/pkg/cluster"
)

type q6vars struct {
	marker   string
	scenario string
	skip     bool
	prompt   string
}

var q6 = q6vars{
	marker:   "Q6",
	scenario: "6",
	skip:     false,
	prompt: fmt.Sprintf(`#----------------------------------
# Scenario 6 - RBAC ServiceAccount Role RoleBinding
#----------------------------------

%sCONTEXT%s:   Create RBAC Service Account bindings that can be used with roles
%sOBJECTIVE 1%s: 

- Create a new ServiceAccount called %s
- Create a new Role called %s
- Create a new RoleBinding called %s
- The new SA to only create %s and %s

`, red, reset, green, reset, highlight("processor", "green"), highlight("processor", "green"), highlight("processor", "green"), highlight("Secrets", "green"), highlight("Config Maps", "green")),
}

// Stageq6 runs the logic for question 3
func StageQ6() {
	fmt.Printf("[%s] Staging Scenario %s...\n", q6.marker, q6.scenario)

	// Check if q6 dir exists
	if _, err := os.Stat(q6.marker); !os.IsNotExist(err) {
		q6.skip = true
	}

	if !q6.skip {
		pwd, err := os.Getwd()
		if err != nil {
			log.Fatalf("Failed to get present working directory: %v", err)
		}

		folderPath := fmt.Sprintf("%s/%s", pwd, q6.marker)

		err = os.Mkdir(folderPath, 0775) // Set perms
		if err != nil {
			log.Fatalf("Failed to create folder: %v", err)
		}
	}

	fmt.Printf("[%s] Successfully created %s dir!\n", q6.marker, q6.marker)

	// Check if q6 Namespace exists
	namespace := fmt.Sprintf("q%s-ns", q6.scenario)
	if cluster.NamespaceExists(namespace) {
		fmt.Printf("[%s] Namespace %s already exists!\n", q6.marker, namespace)
		q6.skip = true
	} else {
		fmt.Printf("[%s] Namespace does not exist!\n", q6.marker)
		q6.skip = false
	}

	if !q6.skip {
		// Create q6 Namespace
		err := cluster.CreateNamespace(namespace, q6.marker)
		if err != nil {
			fmt.Println("Error setting up namespace:", err)
			return
		}

	}

	fmt.Printf("[%s] Successfully staged %s scenario!\n", q6.marker, q6.marker)
	fmt.Printf("[%s] Please run 'kr start %s'!\n", q6.marker, strings.ToLower(q6.marker))

}

func StartQ6() {
	fmt.Println(q6.prompt)
}

func UnstageQ6() {
	fmt.Printf("[%s] Unstaging Kubernetes Question %s...\n", q6.marker, q6.scenario)

	// Check if q6 Namespace exists
	namespace := fmt.Sprintf("q%s-ns", q6.scenario)
	if !cluster.NamespaceExists(namespace) {
		fmt.Printf("[%s] Namespace %s does not exist!\n", q6.marker, namespace)
	} else {
		// Remove q6 Namespace
		fmt.Printf("[%s] Deleting %s!\n", q6.marker, q6.marker)
		err := cluster.DestroyNamespace(namespace, q6.marker)
		if err != nil {
			log.Fatalf("Error deleting namespace:", err)
		}
	}

	// Remove q6 dir
	if _, err := os.Stat(q6.marker); os.IsNotExist(err) {
		fmt.Printf("[%s] Answer directory is already gone!\n", q6.marker)
	} else {
		err = os.RemoveAll(q6.marker)
		if err != nil {
			log.Fatalf("Failed to remove %s dir: %v", err, q6.marker)
		}
	}

	fmt.Printf("[%s] Successfully unstaged %s scenario!\n", q6.marker, q6.marker)
}

func ValidateQ6() {
	fmt.Printf("[%s] Validating Kubernetes Question %s...\n", q6.marker, q6.scenario)

	// Validate that the namespace exists
	namespace := fmt.Sprintf("q%s-ns", q6.scenario)
	if !cluster.NamespaceExists(namespace) {
		msg := fmt.Sprintf("[%s] Namespace %s does not exist!\n", q6.marker, namespace)
		fmt.Println(highlight(msg, "red"))
		os.Exit(0)
	} else {
		fmt.Printf("[%s] Namespace %s exists!\n", q6.marker, namespace)
	}

	// Validate that the ServiceAccount exists
	saAcct := "processor"
	if !cluster.CheckSaAcctExists(namespace, saAcct) {
		msg := fmt.Sprintf("[%s] ServiceAccount %s does not exist!\n", q6.marker, saAcct)
		fmt.Println(highlight(msg, "red"))
		os.Exit(0)
	} else {
		fmt.Printf("[%s] ServiceAccount %s exists!\n", q6.marker, saAcct)
	}

	// Validate that the Role exists
	role := "processor"
	if !cluster.CheckRoleExists(namespace, role) {
		msg := fmt.Sprintf("[%s] Role %s does not exist!\n", q6.marker, role)
		fmt.Println(highlight(msg, "red"))
		os.Exit(0)
	} else {
		fmt.Printf("[%s] Role %s exists!\n", q6.marker, role)
	}

	// Validate that the RoleBinding exists
	roleBinding := "processor"
	if !cluster.CheckRoleBindingExists(namespace, roleBinding) {
		msg := fmt.Sprintf("[%s] RoleBinding %s does not exist!\n", q6.marker, roleBinding)
		fmt.Println(highlight(msg, "red"))
		os.Exit(0)
	} else {
		fmt.Printf("[%s] RoleBinding %s exists!\n", q6.marker, roleBinding)
	}

	// Validate that the ServiceAccount can create Secrets
	action := "secrets"
	if !cluster.CanI(namespace, saAcct, action) {
		msg := fmt.Sprintf("[%s] ServiceAccount %s cannot create %s!\n", q6.marker, saAcct, action)
		fmt.Println(highlight(msg, "red"))
		os.Exit(0)
	} else {
		fmt.Printf("[%s] ServiceAccount %s can create %s!\n", q6.marker, saAcct, action)
	}

	// Validate that the ServiceAccount can create ConfigMaps
	action = "configmaps"
	if !cluster.CanI(namespace, saAcct, action) {
		msg := fmt.Sprintf("[%s] ServiceAccount %s cannot create %s!\n", q6.marker, saAcct, action)
		fmt.Println(highlight(msg, "red"))
		os.Exit(0)
	} else {
		fmt.Printf("[%s] ServiceAccount %s can create %s!\n", q6.marker, saAcct, action)
	}

	// Print validation complete
	success := fmt.Sprintf(`[%s] %sValidation complete!`, q6.marker, green)
	fmt.Println(success)
}
