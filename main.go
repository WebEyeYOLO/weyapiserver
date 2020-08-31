package main

import (
	"github.com/astaxie/beego"
	"os"
	_ "weyapiserver/routers"
	"weyapiserver/version"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		version.ShowVersion("weyapiserver")
	}
	beego.Run()
}
