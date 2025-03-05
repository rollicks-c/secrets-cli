package ui

import (
	"fmt"
	"github.com/gen2brain/beeep"
	"github.com/rollicks-c/secrets-cli/internal/config"
	"os/exec"
	"runtime"
	"strings"
)

func PutInClipboard(data string) error {
	cmd := exec.Command("xsel", "--clipboard", "--input")
	cmd.Stdin = strings.NewReader(data)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func NotifyUser(msg string) error {
	err := beeep.Alert(config.AppLabel, msg, "assets/warning.png")
	if err != nil {
		return err
	}
	return nil
}

func OpenURL(url string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	default:
		return fmt.Errorf("unsupported platform")
	}

	return cmd.Run()
}
