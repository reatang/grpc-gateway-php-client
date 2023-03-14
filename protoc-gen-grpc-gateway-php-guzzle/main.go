package main

import (
	"flag"
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
)

const (
	exampleTip = `
TODO 先挖个坑
`
)

func main() {
	var example bool
	flag.BoolVar(&example, "example", false, "usage example")
	flag.Parse()
	if example {
		fmt.Printf("%s", exampleTip)
		return
	}

	var flags flag.FlagSet

	var serverName string
	flags.StringVar(&serverName, "serverName", "", "server name for plugin")

	options := protogen.Options{
		ParamFunc: flags.Set,
	}

	options.Run(func(gen *protogen.Plugin) error {
		// TODO 先挖个坑
		return nil
	})
}
