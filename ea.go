package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func sleep(timeout int) {
	time.Sleep(time.Duration(timeout) * time.Millisecond)
}

func ea() {
	const EAAppPath string = "C:\\Program Files\\Electronic Arts\\EA Desktop\\EA Desktop\\EALauncher.exe"
	_, err := os.Stat(EAAppPath)
	if err != nil {
		return
	}
	fmt.Println("Launching EA App")
	cmd := exec.Command(EAAppPath)
	cmd.Start()
	sleep(5000)
}
