package main

import (
    "os"
    "path/filepath"
    "fmt"
)

func DiscordFixPatch(location string) error {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()
    updateExePath := filepath.Join(os.Getenv("LOCALAPPDATA"), "Discord", "Update.exe")
    _, err := os.Stat(updateExePath)
    if err != nil {
        // Uninstall Discord Plugin
        dllPath := filepath.Join(location, "R2Northstar", "plugins", "DiscordRPC.dll")
        err := os.Remove(dllPath)
        if err != nil {
            return err
        }
    }
    return nil
}
