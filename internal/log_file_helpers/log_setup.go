package logfilehelpers


import (
	"fmt"
	"os"
	"path"
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