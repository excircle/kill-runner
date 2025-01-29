package utils

var ConfigPath string = "kill.config"
var LoggingFile string = "kill.log"
var DbPath string = "killdb.sqlite"
var LogLevel int

func SetGlobalLogLevel(level int) {
	LogLevel = level
}
