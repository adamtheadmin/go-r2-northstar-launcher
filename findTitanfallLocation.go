package main

import (
	"os"
)

var locations = []string{
	"C:/Program Files (x86)/Origin Games/Titanfall2",
	"C:/Program Files (x86)/Steam/steamapps/common/Titanfall2",
	"C:/Program Files/EA Games/Titanfall2",
}

func findTitanfallLocation() string {
	for _, location := range locations {
		if _, err := os.Stat(location); err == nil {
			return location
		}
	}
	panic("Could not find Titanfall 2 location")
}
