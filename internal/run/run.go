package run

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/gitmobkab/balafetch/internal/api"
	"github.com/gitmobkab/balafetch/internal/exit_codes"
	"github.com/gitmobkab/balafetch/internal/imageutil"
	"github.com/gitmobkab/balafetch/internal/model"
	"github.com/gitmobkab/balafetch/internal/random"
	"github.com/gitmobkab/balafetch/internal/strings_helpers"
	"github.com/gitmobkab/balafetch/internal/data"
)


/*
BalafetchDependencies is a struct that holds the dependencies for the balafetch run.
This allows for easier testing and separation of concerns.
*/
type BalafetchDependencies struct {
	GetFromBalatroApi func(params map[string]string, timeout int) ([]byte, error)
	GetImage func(url string, timeout int) ([]byte, error)
	ExecFastfetch func(imagePath string) error
}

/*
DefaultBalafetchDependencies returns the default dependencies for the balafetch run.
This is used in the FullBalafetchRun function, but can be overridden for testing purposes.
*/
func DefaultBalafetchDependencies() BalafetchDependencies {
	return BalafetchDependencies{
		GetFromBalatroApi: api.GetFromBalatroApi,
		GetImage: api.GetRequest,
		ExecFastfetch: RunFastfetch,
	}
}

/*
run the balafetch script and return either (int,nil) or (int, error)

the first returned value is the program exit code.
the second returned value is either nil or the first error the funtion encounter

the exit codes are defined in internal/exit_codes/exit_codes.go,
*/
func RunBalafetch(ctx model.Ctx, Dependencies BalafetchDependencies) (int, error){
	global_picker := random.NewPicker(time.Now().Unix())

	CategoryTitle := ctx.CardCategory
	timeout := ctx.Timeout
	
	
	if CategoryTitle == "" {
		CategoryTitle = global_picker.PickRandomBalatroCardCategory()
	} else {
		NormalizedCategory := strings_helpers.LowerCase(CategoryTitle)
		GottenCategory, exists := data.CategoryResolution[NormalizedCategory]
		if !exists {
			fmt.Printf("balafetch could not resolve '%s' to a valid balatro card category\n", CategoryTitle)
			fmt.Println("Note: use double quotes for categories with spaces, e.g. \"tarot cards\"")
			data.DisplayCategoryHelp()
			return exitCodes.CommandErrorCode, nil
		} else {
			CategoryTitle = GottenCategory
		}
	}

	CategoryTitle = strings_helpers.TitleCase(CategoryTitle)
	
	ImagesListParams := map[string]string{
		"action": "query",
		"prop": "images",
		"titles": CategoryTitle,
		"imlimit": "max",
		"format": "json",
	}
	
	ResponseData, RequestErr := Dependencies.GetFromBalatroApi(ImagesListParams,timeout)
	if RequestErr != nil {
		return exitCodes.RequestFailureCode, RequestErr
	}

	var imagesList model.ImagesListResponse
	if err := json.Unmarshal(ResponseData, &imagesList); err != nil {
		return exitCodes.ApiResponseParsingFailureCode, err
	}

	imageTitle := imageutil.GetRandomImageTitle(imagesList)

	ImageInfoParams := map[string]string{
		"action":"query",
		"prop":"imageinfo",
		"titles": imageTitle,
		"iiprop":"url",
		"format":"json",
	}
	
	ImageData, RequestErr := Dependencies.GetFromBalatroApi(ImageInfoParams, timeout)
	if RequestErr != nil {
		return exitCodes.RequestFailureCode, RequestErr
	}

	var imageInfo model.ImageInfoResponse
	if err := json.Unmarshal(ImageData, &imageInfo); err != nil{
		return exitCodes.ApiResponseParsingFailureCode, err
	}

	image_url := imageutil.GetImageUrl(imageInfo)
	image_data, err := Dependencies.GetImage(image_url, timeout)
	if err != nil {
		return exitCodes.RequestFailureCode, err
	}

	f, err := os.CreateTemp("","balatro-*.png")
	
	if err != nil {
		return exitCodes.FileIOErrorCode, err
	}


	if _, err := f.Write(image_data); err != nil {
		return exitCodes.FileIOErrorCode, err
	}
	if err := f.Close(); err != nil {
		return exitCodes.FileIOErrorCode, err
	}

	defer os.Remove(f.Name())
	
	if err := Dependencies.ExecFastfetch(f.Name()); err != nil {
		fmt.Println("Error running fastfetch:", err)
		return exitCodes.CommandErrorCode, nil
	}

	return exitCodes.SuccessCode, nil
}

