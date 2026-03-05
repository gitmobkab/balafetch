package imageutil

import (
	"testing"
	"github.com/gitmobkab/balafetch/internal/model"
)

func makeImagesListResponse(titles []string) model.ImagesListResponse {
	images := make([]struct {
		Title string `json:"title"`
	}, len(titles))
	for i, t := range titles {
		images[i].Title = t
	}

	var r model.ImagesListResponse
	r.Query.Pages = map[string]struct {
		Images []struct {
			Title string `json:"title"`
		} `json:"images"`
	}{
		"1": {Images: images},
	}
	return r
}

func makeImageInfoResponse(imageUrl string) model.ImageInfoResponse {
	var r model.ImageInfoResponse
	r.Query.Pages = map[string]struct {
		ImageInfo []struct {
			Url string `json:"url"`
		} `json:"imageinfo"`
	}{
		"1": {
			ImageInfo: []struct {
				Url string `json:"url"`
			}{{Url: imageUrl}},
		},
	}
	return r
}

func TestGetRandomImageTitle_ReturnsOneOfTheTitles(t *testing.T) {
	titles := []string{"File:Joker.png", "File:TheWorld.png", "File:TheFool.png"}
	response := makeImagesListResponse(titles)

	for range 50 {
		_ , err := GetRandomImageTitle(response)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	}
}

func TestGetRandomImageTitle_EmptyPages_ReturnsError(t *testing.T) {
	var r model.ImagesListResponse
	r.Query.Pages = map[string]struct {
		Images []struct {
			Title string `json:"title"`
		} `json:"images"`
	}{}

	_ , err := GetRandomImageTitle(r)
	if err == nil {
		t.Errorf("expected error string for empty pages, got nil")
	}
}

func TestGetRandomImageTitle_PagesWithoutImages_ReturnsError(t *testing.T) {
	var r model.ImagesListResponse
	r.Query.Pages = map[string]struct {
		Images []struct {
			Title string `json:"title"`
		} `json:"images"`
	}{
		"1": {Images: []struct {
			Title string `json:"title"`
		}{}},
	}

	_ , err := GetRandomImageTitle(r)
	if err == nil {
		t.Errorf("expected error string for pages without images, got nil")
	}
}

func TestGetImageUrl_ReturnsCorrectUrl(t *testing.T) {
	expected := "https://static.wikia.nocookie.net/balatrogame/card.png"
	response := makeImageInfoResponse(expected)

	result, _ := GetImageUrl(response)
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestGetImageUrl_EmptyPages_ReturnsEmpty(t *testing.T) {
	var r model.ImageInfoResponse
	r.Query.Pages = map[string]struct {
		ImageInfo []struct {
			Url string `json:"url"`
		} `json:"imageinfo"`
	}{}

	_ , err := GetImageUrl(r)
	if err == nil {
		t.Errorf("expected non nil error for empty pages, got nil")
	}
}

func TestGetImageUrl_PagesWithoutImageInfo_ReturnsEmpty(t *testing.T) {
	var r model.ImageInfoResponse
	r.Query.Pages = map[string]struct {
		ImageInfo []struct {
			Url string `json:"url"`
		} `json:"imageinfo"`
	}{
		"1": {ImageInfo: []struct {
			Url string `json:"url"`
		}{}},
	}

	_ , err := GetImageUrl(r)
	if err == nil {
		t.Errorf("expected non nil error for pages without image info, got nil")
	}
}