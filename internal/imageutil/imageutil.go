package imageutil

import (
	"time"
	"github.com/gitmobkab/balafetch/internal/model"
	"github.com/gitmobkab/balafetch/internal/random"
)

func GetRandomImageTitle(imagesResponse model.ImagesListResponse) string{
	
	var imageTitle string
	picker := random.NewPicker(time.Now().Unix())
	for _, page := range imagesResponse.Query.Pages {
		images := page.Images
		choosenImageIndex := picker.PickInt(len(images))
		imageTitle = images[choosenImageIndex].Title
		break
	}
	return imageTitle
}

func GetImageUrl(imageInfo model.ImageInfoResponse) string{
	var url string

	for _,page := range imageInfo.Query.Pages {
		for _,image_info := range page.ImageInfo {
			url = image_info.Url
			break
		}
		break
	}
	return url
}
