package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

const (
	//TTVERSION tt default version
	TTVERSION = "0.2"
)

type versionInfo struct {
	TTVersion string
	GoVersion string
	OSArch    string
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show version of tt",
	Run:   versionRun,
}

func versionRun(cmd *cobra.Command, args []string) {
	info := getTTVersion()
	fmt.Println("tt version:", info.TTVersion)
	fmt.Println("go version:", info.GoVersion)
	fmt.Println("os/arch:   ", info.OSArch)
}

func init() {
	RootCmd.AddCommand(versionCmd)
}

func getTTVersion() *versionInfo {
	var versionInfo = &versionInfo{
		TTVersion: TTVERSION,
		GoVersion: runtime.Version(),
		OSArch:    getOSArch(),
	}
	return versionInfo
}

func getOSArch() string {
	return runtime.GOOS + "/" + runtime.GOARCH
}
