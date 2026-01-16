package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"github.com/gitmobkab/balafetch/internal/run"
)

func SetupLogFile() (string, error){

	const BalafetchDirName string = "balafetch" 
	const BalafetchLogFileName string = "balafetch.log"

	UserHomeDir, UserHomeDirErr := os.UserHomeDir()
	if UserHomeDirErr != nil {
		fmt.Println(UserHomeDirErr)
		return "", fmt.Errorf("Failed to get User Home Directory: %w", UserHomeDirErr)
	}

	BalafetchDirPath := path.Join(UserHomeDir, BalafetchDirName)
	MkdirErr := os.MkdirAll(BalafetchDirPath, 0744) 
	// ok, it's weird or maybe i don't get it
	// but in order for us to open/create the log file,
	// the user must explicitely possess the rwx rights
	// so we don't make a dir with rw- permission
	// but rather rwx
	if MkdirErr != nil {
		return "", fmt.Errorf("Failed to make Balafetch Dir: %w", MkdirErr)
	}

	logFilePath := path.Join(BalafetchDirPath, BalafetchLogFileName)
	return logFilePath, nil
}

func main(){
	const LogFileSetupFailureCode int = 50

	logFilePath, logFileSetupErr := SetupLogFile()
	if logFileSetupErr != nil{
		fmt.Println("Balafetch setup failed\n",logFileSetupErr)
		os.Exit(LogFileSetupFailureCode)
	}
	
	logFile, OpenFileErr := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if OpenFileErr != nil{
		fmt.Println("Unable to open log file:", OpenFileErr)
		os.Exit(LogFileSetupFailureCode)
	}

	log.SetOutput(logFile)
	
	balafetchRunExitCode, balafetchRunErr := run.RunBalafetch()
	switch {
	case balafetchRunExitCode != run.SuccessCode && balafetchRunExitCode != run.CommandErrorCode:
		run.RunFastfetchDefault()
		log.Printf("[ Error Code: %d ]\n",balafetchRunExitCode)
		log.Println(balafetchRunErr)
	case balafetchRunExitCode == run.CommandErrorCode:
		fmt.Println("Fastfetch not installed in your system")
	}

	logFile.Close()
	os.Exit(balafetchRunExitCode)
}