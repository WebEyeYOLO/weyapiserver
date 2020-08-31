package version

import (
	"fmt"
	"os"
)

var version string

//ShowVersion 显示版本
func ShowVersion(module string) {
	if version != "" {
		fmt.Printf("Rainbond %s %s\n", module, version)
	} else {
		fmt.Printf("Rainbond %s %s\n", module, os.Getenv("RELEASE_DESC"))
	}
	os.Exit(0)
}

//GetVersion GetVersion
func GetVersion() string {
	return version
}
