package updater

import (
	"context"
	"fmt"
	"runtime"
	"time"

	selfupdate "github.com/creativeprojects/go-selfupdate"
	"github.com/gitmobkab/balafetch/internal/data"
)

func Update(timeout int) error {
	source, err := selfupdate.NewGitHubSource(selfupdate.GitHubConfig{})
	if err != nil {
		return fmt.Errorf("failed to create update source: %w", err)
	}

	assetName := fmt.Sprintf("balafetch-%s-%s", runtime.GOOS, runtime.GOARCH)
	if runtime.GOOS == "windows" {
		assetName += ".exe"
	}

	updater, err := selfupdate.NewUpdater(selfupdate.Config{
		Source:    source,
		Validator: &selfupdate.ChecksumValidator{UniqueFilename: "checksums.txt"},
		Filters:   []string{assetName},
	})
	if err != nil {
		return fmt.Errorf("failed to create updater: %w", err)
	}

	ctx := context.Background()
	if timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
		defer cancel()
	}

	exe, err := selfupdate.ExecutablePath()
	if err != nil {
		return fmt.Errorf("failed to locate executable: %w", err)
	}

	release, err := updater.UpdateCommand(ctx, exe, data.Version, selfupdate.ParseSlug("gitmobkab/balafetch"))
	if err != nil {
		return fmt.Errorf("update failed: %w", err)
	}

	if release.Equal(data.Version) {
		fmt.Printf("balafetch is already up to date (V%s)\n", data.Version)
		return nil
	}

	fmt.Printf("successfully updated balafetch to V%s\n", release.Version())
	return nil
}
