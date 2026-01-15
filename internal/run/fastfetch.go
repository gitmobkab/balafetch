package run

import (
  "os"
  "os/exec"
)

func RunFastfetch(logo string) error {
  cmd := exec.Command("fastfetch", "-l", logo)
  cmd.Stdout = os.Stdout
  return cmd.Run()
}

func RunFastfetchDefault() error {
  cmd := exec.Command("fastfetch")
  cmd.Stdout = os.Stdout
  return cmd.Run()
}