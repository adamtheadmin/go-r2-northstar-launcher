package main

import (
    "archive/zip"
    "io"
    "net/http"
    "os"
    "path/filepath"
    "bytes"
)

func R2Updater(location string, release Release) error {
    var zipFileURL = release.Assets[0].BrowserDownloadURL
    resp, err := http.Get(zipFileURL)
    if err != nil {
        return err
    }
    body, err := io.ReadAll(resp.Body)
    defer resp.Body.Close()
    if err != nil {
        return err;
    }

    r, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
    if err != nil {
        return err
    }

    for _, f := range r.File {
        filePath := filepath.Join(location, f.Name)
        if f.FileInfo().IsDir() {
            os.MkdirAll(filePath, os.ModePerm)
            continue
        }

        rc, err := f.Open()
        if err != nil {
            return err
        }

        defer rc.Close()

        f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
        if err != nil {
            return err
        }

        defer f.Close()

        _, err = io.Copy(f, rc)
        if err != nil {
            return err
        }
    }

    if err := DiscordFixPatch(location); err != nil {
        return err;
    }

    return nil
}
