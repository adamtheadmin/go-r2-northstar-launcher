package main

import (
    "io/ioutil"
    "strconv"
    "path/filepath"
    "os"
    "bytes"
)

const versionFilename = "VERSION"

type VersionManager struct {
    Location string
}

func (vm *VersionManager) Get() (int) {
    file := filepath.Join(vm.Location, versionFilename)
    content, err := ioutil.ReadFile(file)
    if err != nil {
        if !os.IsNotExist(err) {
            return 0
        }
        return 0
    }
    version, err := strconv.Atoi(string(bytes.TrimSpace(content)))
    if err != nil {
        panic(err)
    }
    return version
}

func (vm *VersionManager) Set(version int) {
    file := filepath.Join(vm.Location, versionFilename)
    content := []byte(strconv.Itoa(version))
    err := ioutil.WriteFile(file, content, 0644)
    if err != nil {
        panic(err)
    }
}
