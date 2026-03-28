package ui

import (
	"time"
	"fmt"
)

const blue = "\033[34m"
const reset = "\033[0m"

func PrintTimeoutProgress(label string, timeout int, stream func(string), done <-chan struct{}) {
    ticker := time.NewTicker(time.Duration(100) * time.Millisecond)
    defer ticker.Stop()

    elapsed := 0.0
    for {
        select {
        case <-done:
            stream("") // clear the line
            return
        case <-ticker.C:
            elapsed += 0.1
            if timeout == 0 {
                stream(fmt.Sprintf("%s %s[%.1f/∞ seconds]%s", label, blue, elapsed, reset))
            } else {
                stream(fmt.Sprintf("%s %s[%.1f/%d seconds]%s", label, blue, elapsed, timeout, reset))
            }
        }
    }
}