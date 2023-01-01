package logging

import (
	"io/ioutil"
	"log"
	"os"
)

var (
	// Trace is the logger for the most detailed messages. Should be used for debugging info only.
	// Trace messages will only be visible on dev machines.
	Trace *log.Logger

	// Info is the logger used for non-critical output.
	Info *log.Logger

	// Warning is the logger used for warnings / important messages to the user.
	Warning *log.Logger

	// Error is the logger used for serious / very important messages to the user. Usually you would write a message to
	// the Error logger before you stop the program due to a critical condition.
	Error *log.Logger
)

func init() {
	const flags = log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile | log.Lmsgprefix

	if "development" == os.Getenv("target") {
		Trace = log.New(os.Stdout, "TRACE: ", flags)
	} else {
		Trace = log.New(ioutil.Discard, "TRACE: ", flags)
	}

	Info = log.New(os.Stdout, "INFO: ", flags)
	Warning = log.New(os.Stdout, "WARNING: ", flags)
	Error = log.New(os.Stderr, "ERROR: ", flags)
}
