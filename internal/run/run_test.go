package run

import (
	"errors"
	"fmt"
	"testing"

	exitCodes "github.com/gitmobkab/balafetch/internal/exit_codes"
	"github.com/gitmobkab/balafetch/internal/model"
	"github.com/gitmobkab/balafetch/internal/ui"
)

// --- fake response builders ---

func fakeImagesListResponse(category string) []byte {
	return []byte(fmt.Sprintf(`{
		"query": {
			"pages": {
				"1": {
					"images": [
						{"title": "File:%s_card.png"}
					]
				}
			}
		}
	}`, category))
}

func fakeImageInfoResponse(url string) []byte {
	return []byte(fmt.Sprintf(`{
		"query": {
			"pages": {
				"1": {
					"imageinfo": [
						{"url": "%s"}
					]
				}
			}
		}
	}`, url))
}

// happyDeps returns a BalafetchDependencies where everything succeeds.
func happyDeps() BalafetchDependencies {
	callCount := 0
	return BalafetchDependencies{
		GetFromBalatroApi: func(params map[string]string, timeout int) ([]byte, error) {
			callCount++
			if callCount == 1 {
				return fakeImagesListResponse("Joker"), nil
			}
			return fakeImageInfoResponse("https://fake.url/card.png"), nil
		},
		GetFunc: func(url string, timeout int) ([]byte, error) {
			return []byte("fake image data"), nil
		},
		ExecFastfetch: func(logo string) error {
			return nil
		},
		streamer: ui.NoOpStreamer(),
	}
}

// --- tests ---

func TestRunBalafetch_ValidCategory_Success(t *testing.T) {
	ctx := model.Ctx{CardCategory: "joker", Timeout: 5}
	code, err := RunBalafetch(ctx, happyDeps())

	if code != exitCodes.SuccessCode {
		t.Errorf("expected exit code %d, got %d", exitCodes.SuccessCode, code)
	}
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestRunBalafetch_EmptyCategory_PicksRandom(t *testing.T) {
	// Empty category should not crash — picker selects a valid one
	ctx := model.Ctx{CardCategory: "", Timeout: 5}
	code, err := RunBalafetch(ctx, happyDeps())

	if code != exitCodes.SuccessCode {
		t.Errorf("expected exit code %d, got %d", exitCodes.SuccessCode, code)
	}
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestRunBalafetch_InvalidCategory_CommandError(t *testing.T) {
	ctx := model.Ctx{CardCategory: "notarealcategory", Timeout: 5}
	code, err := RunBalafetch(ctx, happyDeps())

	if code != exitCodes.CommandErrorCode {
		t.Errorf("expected exit code %d, got %d", exitCodes.CommandErrorCode, code)
	}
	if err != nil {
		t.Errorf("expected nil error for command error, got %v", err)
	}
}

func TestRunBalafetch_ApiFailure_FirstCall_RequestError(t *testing.T) {
	deps := happyDeps()
	deps.GetFromBalatroApi = func(params map[string]string, timeout int) ([]byte, error) {
		return nil, errors.New("connection refused")
	}

	ctx := model.Ctx{CardCategory: "joker", Timeout: 5}
	code, err := RunBalafetch(ctx, deps)

	if code != exitCodes.RequestFailureCode {
		t.Errorf("expected exit code %d, got %d", exitCodes.RequestFailureCode, code)
	}
	if err == nil {
		t.Error("expected an error, got nil")
	}
}

func TestRunBalafetch_ApiFailure_SecondCall_RequestError(t *testing.T) {
	callCount := 0
	deps := happyDeps()
	deps.GetFromBalatroApi = func(params map[string]string, timeout int) ([]byte, error) {
		callCount++
		if callCount == 1 {
			return fakeImagesListResponse("Joker"), nil
		}
		return nil, errors.New("timeout")
	}

	ctx := model.Ctx{CardCategory: "joker", Timeout: 5}
	code, err := RunBalafetch(ctx, deps)

	if code != exitCodes.RequestFailureCode {
		t.Errorf("expected exit code %d, got %d", exitCodes.RequestFailureCode, code)
	}
	if err == nil {
		t.Error("expected an error, got nil")
	}
}

func TestRunBalafetch_InvalidJson_FirstCall_ParseError(t *testing.T) {
	deps := happyDeps()
	deps.GetFromBalatroApi = func(params map[string]string, timeout int) ([]byte, error) {
		return []byte(`not valid json{{{`), nil
	}

	ctx := model.Ctx{CardCategory: "joker", Timeout: 5}
	code, err := RunBalafetch(ctx, deps)

	if code != exitCodes.ApiResponseParsingFailureCode {
		t.Errorf("expected exit code %d, got %d", exitCodes.ApiResponseParsingFailureCode, code)
	}
	if err == nil {
		t.Error("expected a parse error, got nil")
	}
}

func TestRunBalafetch_ImageDownloadFailure_RequestError(t *testing.T) {
	deps := happyDeps()
	deps.GetFunc = func(url string, timeout int) ([]byte, error) {
		return nil, errors.New("image download failed")
	}

	ctx := model.Ctx{CardCategory: "joker", Timeout: 5}
	code, err := RunBalafetch(ctx, deps)

	if code != exitCodes.RequestFailureCode {
		t.Errorf("expected exit code %d, got %d", exitCodes.RequestFailureCode, code)
	}
	if err == nil {
		t.Error("expected an error, got nil")
	}
}

func TestRunBalafetch_FastfetchFailure_CommandError(t *testing.T) {
	deps := happyDeps()
	deps.ExecFastfetch = func(logo string) error {
		return errors.New("fastfetch: command not found")
	}

	ctx := model.Ctx{CardCategory: "joker", Timeout: 5}
	code, err := RunBalafetch(ctx, deps)

	if code != exitCodes.CommandErrorCode {
		t.Errorf("expected exit code %d, got %d", exitCodes.CommandErrorCode, code)
	}
	if err != nil {
		t.Errorf("expected nil error for command error, got %v", err)
	}
}

func TestRunBalafetch_CaseInsensitiveCategory(t *testing.T) {
	cases := []string{"JOKER", "Joker", "joker", "jOkEr"}
	for _, cat := range cases {
		t.Run(cat, func(t *testing.T) {
			ctx := model.Ctx{CardCategory: cat, Timeout: 5}
			code, err := RunBalafetch(ctx, happyDeps())
			if code != exitCodes.SuccessCode {
				t.Errorf("category %q: expected %d, got %d", cat, exitCodes.SuccessCode, code)
			}
			if err != nil {
				t.Errorf("category %q: unexpected error: %v", cat, err)
			}
		})
	}
}