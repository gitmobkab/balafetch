package run

import (
	"fmt"
	"log"
	"os"
	"github.com/gitmobkab/balafetch/internal/log_file_helpers"
	"github.com/gitmobkab/balafetch/internal/exit_codes"
	"github.com/gitmobkab/balafetch/internal/model"
)


func FullBalafetchRun(ctx model.Ctx) (int){
	logFilePath, logFileSetupErr := logfilehelpers.SetupLogFile()
	if logFileSetupErr != nil{
		fmt.Println("Balafetch setup failed\n",logFileSetupErr)
		os.Exit(exitCodes.LogFileSetupFailureCode)
	}
	
	logFile, OpenFileErr := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if OpenFileErr != nil{
		fmt.Println("Unable to open log file:", OpenFileErr)
		os.Exit(exitCodes.LogFileSetupFailureCode)
	}

	log.SetOutput(logFile)
	
	balafetchRunExitCode, balafetchRunErr := RunBalafetch(ctx)
	switch {
	case balafetchRunExitCode != exitCodes.SuccessCode && balafetchRunExitCode != exitCodes.CommandErrorCode:
		RunFastfetchDefault()
		log.Printf("[ Error Code: %d ]\n",balafetchRunExitCode)
		log.Println(balafetchRunErr)
	}

	logFile.Close()
	return balafetchRunExitCode
}