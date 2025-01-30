package utils

var ConfigPath string = "kill.config"
var LoggingFile string = "kill.log"
var DbPath string = "killdb.sqlite"
var LogLevel int
var Kubeconfig string

func SetGlobalLogLevel(level int) {
	LogLevel = level
}

func SetGlobalKubeConfig(path string) {
	Kubeconfig = path
}
