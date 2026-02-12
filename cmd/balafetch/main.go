package main

import (
	"os"
	"github.com/gitmobkab/balafetch/internal/run"
	"github.com/gitmobkab/balafetch/internal/exit_codes"
	pflag "github.com/spf13/pflag"
)


func main(){
	helpFlag := pflag.BoolP("help", "h", false, "Show help information")
	versionFlag := pflag.BoolP("version", "v", false, "Show version information")
	timeoutFlag := pflag.IntP("timeout", "t", 20, "Set the timeout for API requests in seconds")

	pflag.Parse()

	var exitCode int = exitCodes.SuccessCode

	if *helpFlag {
		Help()
	} else if *versionFlag {
		Version()
	} else {
		exitCode = run.FullBalafetchRun(*timeoutFlag)
	}

	os.Exit(exitCode)
}