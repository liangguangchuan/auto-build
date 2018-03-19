package cmd

import (
	"github.com/liangguangchuan/auto-build/cmd/build"
	_ "github.com/liangguangchuan/auto-build/cmd/build/golang"
	"github.com/liangguangchuan/auto-build/utils"
)

var tmpl = ` welcome start xiaochuan auto-build
 AVAILABLE COMMANDS
{{range .}}{{if .Runnable}}
    {{.Name | printf "%-11s" }} {{.Short}}{{end}}{{end}}

Use {{"auto-build help [command]" }} for more information about a command.
`

var helpTemplate = `
USAGE
	{{.Help}}
`

var ErrorTemplate = `auto-build: %s.
Use auto-build help for more information.
`

func Usage() {
	utils.Tmpl(tmpl, build.AvailableCommands)
}

func Help(args []string) {
	if len(args) == 0 {
		Usage()
	}
	if len(args) != 1 {
		utils.PrintErrorAndExit("Too many arguments", ErrorTemplate)
	}

	arg := args[0]

	for _, cmd := range build.AvailableCommands {
		if cmd.Name() == arg {
			utils.Tmpl(helpTemplate, cmd)
			return
		}
	}
	utils.PrintErrorAndExit("Unknown help topic", ErrorTemplate)
}
