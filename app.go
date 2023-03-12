package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	location, err := findTitanfallLocation()
	if err != nil {
    	panic(err)
    }
	vm := &VersionManager{Location: location}
	fmt.Println("Found Titanfall at: " + location)
	fmt.Println("Downloading R2 Release list...")
	releases := getReleases()
	latest := releases[0]
	currentVersion := vm.Get()
	latestVersion, err := strconv.Atoi(regexp.MustCompile("[^0-9]").ReplaceAllString(latest.Name, ""))
	if err != nil {
		panic(err)
	}
	if currentVersion != latestVersion {
		fmt.Println("Installing R2Northstar at latest version... ")
		err := R2Updater(location, latest)
		if err != nil {
			panic(err)
		}
		vm.Set(latestVersion)
	} else {
		fmt.Println("Latest version already installed!")
	}
	ea()
	fmt.Println("Launching R2Northstar")
	Launch(location)
}
