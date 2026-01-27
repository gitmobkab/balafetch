package main

import (
	"fmt"
	"os"
	"github.com/gitmobkab/balafetch/internal/run"
	"github.com/gitmobkab/balafetch/internal/exit_codes"
)


func main(){
	var flags = map[string]func() int{
		"-v": Version,
		"-h": Help,
	}

	cmdArgs := os.Args[1:]
	var exitCode int = 0

	if len(cmdArgs) == 0 {
		exitCode = run.FullBalafetchRun()
	} else if len(cmdArgs) > 1 {
		fmt.Println("Too many arguments provided. Use -h for help.")
		fmt.Println("Usage: balafetch [-h | -v]")
		exitCode = exitCodes.InvalidUsageFailureCode
	} else {
		if action, exists := flags[cmdArgs[0]]; exists {
			exitCode = action()
		} else {
			fmt.Printf("Unknown argument: '%s'. Use -h for help.\n", cmdArgs[0])	
			exitCode = exitCodes.InvalidUsageFailureCode	
		}
	}
	os.Exit(exitCode)
}