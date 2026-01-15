package model


type ImagesListResponse struct {
	Query struct {
		Pages map[string]struct {
			Images []struct {
				Title string `json:"title"`
			} `json:"images"`
		} `json:"pages"`
	} `json:"query"`
}

type ImageInfoResponse struct {
	Query struct {
		Pages map[string]struct {
			ImageInfo []struct {
				Url string `json:"url"`
			} `json:"imageinfo"`
		} `json:"pages"`
	} `json:"query"`
}


