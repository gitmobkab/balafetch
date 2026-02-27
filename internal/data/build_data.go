package data

import "fmt"

var (
	Version string = "dev"
	BuildTime string = "local"
	CommitHash string = "dirty"
)

/* 
returns a formatted version string with build information,
the format is:

V<Version> (<OS>/<ARCH>) [<BuildTime>] *<CommitHash>
*/
func GetVersion() string {
	return fmt.Sprintf("V%s (%s/%s) [%s] *%s", Version, OS, ARCH, BuildTime, CommitHash)
}