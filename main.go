package main

import (
	"flag"
	"log"
	"os"

	"github.com/liangguangchuan/auto-build/cmd"
	"github.com/liangguangchuan/auto-build/cmd/build"
)

func main() {

	flag.Usage = cmd.Usage
	flag.Parse()
	log.SetFlags(0)

	args := flag.Args()

	if len(args) < 1 {
		cmd.Usage()
		os.Exit(2)
		return
	}

	if args[0] == "help" {
		cmd.Help(args[1:])
		return
	}

	for _, c := range build.AvailableCommands {
		if c.Name() == args[0] && c.Runnable() {

			os.Exit(c.Run(c, args[1:]))
			return
		}
	}

}
