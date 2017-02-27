package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/adolphlwq/tt/cmd"
)

func main() {
	checkOS()
	cmd.Execute()
}

func checkOS() {
	if runtime.GOOS != "linux" {
		fmt.Println("tt can only run in linux")
		os.Exit(-1)
	}
}
