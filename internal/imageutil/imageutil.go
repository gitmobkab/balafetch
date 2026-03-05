package imageutil

import (
	"errors"
	"time"
	"github.com/gitmobkab/balafetch/internal/model"
	"github.com/gitmobkab/balafetch/internal/random"
)

// GetRandomImageTitle takes the response of the images list api call
// and returns the title of a random image from the list of the first page only.
// it return an empty string and an error if there's no pages or if the pages don't contain any images.
// otherwise it returns the picked title image and nil.
func GetRandomImageTitle(imagesResponse model.ImagesListResponse) (string, error){
	
	var imageTitle string
	picker := random.NewPicker(time.Now().Unix())
	for _, page := range imagesResponse.Query.Pages {
		images := page.Images
		if len(images) == 0 {
			return "", errors.New("no images found in the first page")
		}
		choosenImageIndex := picker.PickInt(len(images))
		imageTitle = images[choosenImageIndex].Title
		return imageTitle, nil
	}
	return "",errors.New("API response doesn't contain any page")
}

// GetImageUrl takes the response of the image info api call and returns the url of the image.
// it returns an empty string if the response doesn't contain any page or if the pages don't contain any image info.
// it works similarly to GetRandomImageTitle, it only looks at the first page and the first image info of that page.
func GetImageUrl(imageInfo model.ImageInfoResponse) (string, error){
	var url string

	for _,page := range imageInfo.Query.Pages {
		for _,image_info := range page.ImageInfo {
			url = image_info.Url
			return url, nil
		}
		return "", errors.New("no image info found in the first page")
	}
	return "", errors.New("API response doesn't contain any page")
}
