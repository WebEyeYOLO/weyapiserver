package main

import (
	"github.com/astaxie/beego"
	"github.com/spf13/pflag"
	"os"
	"weyapiserver/option"
	_ "weyapiserver/routers"
	"weyapiserver/version"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		version.ShowVersion("weyapiserver")
	}

	option.Conf.AddFlags(pflag.CommandLine)
	pflag.Parse()

	beego.Run()
}
