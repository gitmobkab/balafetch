package logfilehelpers


import (
	"fmt"
	"os"
	"path"
)

func SetupLogFile() (string, error){
	LOG_FILE_PERM := os.FileMode.Perm(0744)

	const BalafetchDirName string = "balafetch" 
	const BalafetchLogFileName string = "balafetch.log"

	UserHomeDir, UserHomeDirErr := os.UserHomeDir()
	if UserHomeDirErr != nil {
		return "", fmt.Errorf("Failed to get User Home Directory: %w", UserHomeDirErr)
	}

	BalafetchDirPath := path.Join(UserHomeDir, BalafetchDirName)
	MkdirErr := os.MkdirAll(BalafetchDirPath, LOG_FILE_PERM) 
	if MkdirErr != nil {
		return "", fmt.Errorf("Failed to make Balafetch Dir: %w", MkdirErr)
	}

	logFilePath := path.Join(BalafetchDirPath, BalafetchLogFileName)
	return logFilePath, nil
}