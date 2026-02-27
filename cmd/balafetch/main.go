package main

import (
	"os"
	"github.com/gitmobkab/balafetch/internal/run"
	"github.com/gitmobkab/balafetch/internal/exit_codes"
	"github.com/gitmobkab/balafetch/internal/model"
	pflag "github.com/spf13/pflag"
)


func main(){
	var exitCode int = exitCodes.SuccessCode
	
	helpFlag := pflag.BoolP("help", "h", false, "Show help information")
	versionFlag := pflag.BoolP("version", "v", false, "Show version information")
	timeoutFlag := pflag.IntP("timeout", "t", 20, "Set the timeout for API requests in seconds")
	pflag.Parse()

	var app_ctx model.Ctx

	card_category := pflag.Arg(0)
	if card_category != "" {
		app_ctx.CardCategory = card_category
	} else {
		app_ctx.CardCategory = ""
	}

	app_ctx.Timeout = *timeoutFlag

	if *helpFlag {
		Help()
	} else if *versionFlag {
		Version()
	} else {
		exitCode = run.FullBalafetchRun(app_ctx)
	}

	os.Exit(exitCode)
}