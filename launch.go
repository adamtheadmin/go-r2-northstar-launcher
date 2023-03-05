package main

import (
    "os/exec"
    "path/filepath"
    "fmt"
)

func Launch(location string) error {
    launchPath := filepath.Join(location, "NorthstarLauncher.exe")
    fmt.Printf("Launching %s\n", launchPath)
    cmd := exec.Command(launchPath)
    return cmd.Start()
}
