package main

import (
	"fmt"
	"log"
	"os"
	"github.com/gitmobkab/balafetch/internal/run"
)

func main(){
	logFile,fileErr := os.OpenFile("balafetch.log",os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if fileErr != nil {
		log.Fatal(fileErr)
	}


	log.SetOutput(logFile)

	exitCode, runErr := run.RunBalafetch()


	switch {
	case exitCode < run.CommandErrorCode:
		run.RunFastfetchDefault()
		log.Printf("[ Error Code: %d ]\n",exitCode)
		log.Println(runErr)
	case exitCode == run.CommandErrorCode:
		fmt.Println("Fastfetch not installed in your system")
	}

	logFile.Close()
	os.Exit(exitCode)
}