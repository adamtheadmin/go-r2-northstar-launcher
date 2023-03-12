package main

import (
	"errors"
	"os"
)

var locations = []string {
	"C:/Program Files (x86)/Origin Games/Titanfall2",
	"C:/Program Files (x86)/Steam/steamapps/common/Titanfall2",
	"C:/Program Files/EA Games/Titanfall2",
}

func findTitanfallLocation() (string, error) {
	for _, location := range locations {
		if _, err := os.Stat(location); err == nil {
			return location, nil
		}
	}
	return "", errors.New("Could not find Titanfall 2 location")
}
