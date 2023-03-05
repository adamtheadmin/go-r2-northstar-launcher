package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

func Launch(location string) error {
	launchPath := filepath.Join(location, "NorthstarLauncher.exe")
	fmt.Printf("Launching %s\n", launchPath)
	cmd := exec.Command(launchPath)
	return cmd.Start()
}
