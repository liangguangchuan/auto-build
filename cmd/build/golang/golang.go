package golang

import (
	"log"

	"github.com/liangguangchuan/auto-build/cmd/build"
	"github.com/liangguangchuan/auto-build/utils"
)

var CmdRun = &build.Command{
	UsageLine: "golang [appname]",
	Short:     "运行应用程序监听当前目录",
	Run:       RunApp,
	Help:      "切换监控目录,运行auto-build golang 即可",
}

var (
	exit chan bool
)

func init() {
	build.AvailableCommands = append(build.AvailableCommands, CmdRun)
	exit = make(chan bool)
}

//运行应用
func RunApp(c *build.Command, args []string) int {
	appName := utils.GetAppName()
	runPath := utils.GetRunPath()
	log.Printf("start run appname '%s'\n", appName)
	utils.AutoBuild()
	utils.WatchPath(runPath)
	return 0
}
