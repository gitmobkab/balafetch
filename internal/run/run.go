package run

import (
	"encoding/json"
	"os"
	"time"

	"github.com/gitmobkab/balafetch/internal/api"
	"github.com/gitmobkab/balafetch/internal/imageutil"
	"github.com/gitmobkab/balafetch/internal/model"
	"github.com/gitmobkab/balafetch/internal/random"
	"github.com/gitmobkab/balafetch/internal/exit_codes"
)

/*
run the balafetch script and return either (int,nil) or (int, error)

the first returned value is the program exit code.
the second returned value is either nil or the first error the funtion encounter

the exit codes are

0 -> success

1 -> any network bound error, timeout, dns resolution failed, etc
Note: a non 2xx response code from the api also count in this category

2 -> failure to parse the api response as json

3 -> File I/O bound error, unable to read/write to a file
Note: permission denied is the only error balafetch might get for this

4 -> Command not in user system

for each exit code, the exact error that trigerred returned for context
except 0 exit code

see internal/exit_codes/exit_codes.go for more details
*/
func RunBalafetch(timeout int) (int, error){
	global_picker := random.NewPicker(time.Now().Unix())

	var CategoryTitle string = global_picker.PickRandomBalatroCardCategory()
	
	ImagesListParams := map[string]string{
		"action": "query",
		"prop": "images",
		"titles": CategoryTitle,
		"imlimit": "max",
		"format": "json",
	}
	

	ResponseData, RequestErr := api.GetFromBalatroApi(ImagesListParams,timeout)
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
	
	ImageData, RequestErr := api.GetFromBalatroApi(ImageInfoParams, timeout)
	if RequestErr != nil {
		return exitCodes.RequestFailureCode, RequestErr
	}

	var imageInfo model.ImageInfoResponse
	if err := json.Unmarshal(ImageData, &imageInfo); err != nil{
		return exitCodes.ApiResponseParsingFailureCode, err
	}

	image_url := imageutil.GetImageUrl(imageInfo)
	image_data, err := api.GetRequest(image_url, timeout)
	if err != nil {
		return exitCodes.ApiResponseParsingFailureCode, err
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
	
	
	if err := RunFastfetch(f.Name()); err != nil {
		return exitCodes.CommandErrorCode, err
	}

	return exitCodes.SuccessCode, nil
}

