package pkg

import "log"

func LoggerErr(msg string) {
	log.Printf("\033[31mError:\033[0m %s\n", msg)
}
func LoggerWarn(msg string) {
	log.Printf("\033[33mWarning:\033[0m %s\n", msg)
}
func Logger(msg string) {
	log.Println(msg)
}
