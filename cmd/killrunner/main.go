package main

import (
	"fmt"
	"log"
	"os"

	killrunner "github.com/excircle/kill-runner/cmd"
	"github.com/excircle/kill-runner/pkg/utils"
	"gopkg.in/yaml.v3"
)

func main() {

	//----------------------------------
	// Init Configuration File
	//----------------------------------
	err := utils.ValidateConfig(utils.ConfigPath)
	if err != nil {
		log.Fatalf("Error initializing config: %v", err)
	}

	//----------------------------------
	// Init Logging
	//----------------------------------
	confStruct := utils.DefaultConfig()

	conf, err := os.ReadFile(utils.ConfigPath)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	err = yaml.Unmarshal(conf, &confStruct)
	if err != nil {
		log.Fatalf("Error unmarshalling config file: %v", err)
	}

	utils.SetGlobalLogLevel(confStruct.KillRunner.Config.Logging)
	err = utils.InitLog()
	if err != nil {
		log.Fatalf("Error initializing logging: %v", err)
	}

	utils.ClearTempLogBuffer()

	//----------------------------------
	// Init SQLite DB
	//----------------------------------
	db, err := utils.ValidateDB()
	if err != nil {
		utils.LogEvent(2, "Error initializing database:", err)
	}
	defer db.Close()

	utils.LogEvent(0, "kill-runner is ready to use.")

	//----------------------------------
	// Check Kube Connectivity
	//----------------------------------
	utils.SetGlobalKubeConfig(confStruct.KillRunner.Config.Kubeconfig)
	err = utils.KubeConnect(utils.Kubeconfig)
	if err != nil {
		utils.LogEvent(2, "Error connecting to Kubernetes:", err)
		log.Fatalf("Error connecting to Kubernetes: %v", err)
	} else {
		utils.LogEvent(0, fmt.Sprintf("Successfully connected to Kubernetes using %s.", utils.Kubeconfig))
	}

	//----------------------------------
	// CLI Exec
	//----------------------------------
	killrunner.Execute()
}
