package main

import "github.com/namjagbrawa/jvm_go/base"

func main() {
	cmd := base.ParseCmd()

	if cmd.VersionFlag {
		println("version 0.0.1")
	} else if cmd.HelpFlag || cmd.Class == "" {
		base.PrintUsage()
	} else {
		base.NewJVM(cmd).Start()
	}
}
