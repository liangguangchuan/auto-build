package utils

import (
	"bytes"
	"errors"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/fsnotify/fsnotify"
)

var (
	cmd      *exec.Cmd
	state    sync.Mutex
	dir_list []string
	err      error
)

//监控文件夹
func WatchPath(path string) error {
	var eventTime = make(map[string]int64)
	//获取文件夹列表
	if dir_list, err = getDirList(path); err != nil {
		return err
	}
	//如果文件夹列表为0
	if len(dir_list) == 0 {
		return errors.New("get dir list error")
	}
	done := make(chan bool)
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	// Process events
	go func() {
		for {
			select {
			case e := <-watcher.Events:
				isBuild := true
				//获取文件修改时间
				mt, _ := GetFileModTime(e.Name)
				//如果当前数组保存 修改时间 与文件修改时间一致 就说明同时触发两次 文件修改
				//不再次进行build 操作
				if t := eventTime[e.Name]; t == mt {
					isBuild = false
				}
				eventTime[e.Name] = mt
				//如果不是golang 文件 不进行自动编译
				if filepath.Ext(e.Name) != ".go" {
					isBuild = false
				}

				if isBuild {
					log.Println("操作文件:", e.Name)
					go AutoBuild()
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	for _, dir_path := range dir_list[1:] {
		log.Println("start monitor dir:", dir_path)
		if err = watcher.Add(dir_path); err != nil {
			return err
		}
	}
	//阻塞防止退出
	<-done
	return nil
}

//重启
func Restart(appName string) {
	log.Println("kill running process")
	Kill()
	Start(appName)
}

//启动
func Start(appname string) {
	log.Printf("Restarting '%s'...", appname)
	if !strings.Contains(appname, "./") {
		appname = "./" + appname
	}

	cmd = exec.Command(appname)
	//将运行输出 切换到 当前运行终端
	cmd.Stdout = os.Stdout
	go cmd.Run()
	log.Printf("'%s' is running...", appname)
}

//杀死进程
func Kill() {
	defer func() {
		if e := recover(); e != nil {
			log.Printf("Kill recover: %s", e)
		}
	}()
	if cmd != nil && cmd.Process != nil {
		err := cmd.Process.Kill()
		if err != nil {
			log.Printf("Error while killing cmd process: %s", err)
		}
	}

}

//自动编译
func AutoBuild() {
	var (
		err    error
		stderr bytes.Buffer
	)
	state.Lock()
	defer state.Unlock()
	os.Chdir(currpath)

	appName := GetAppName()
	//如果是windows 添加 exe 后缀
	if runtime.GOOS == "windows" {
		appName += ".exe"
	}
	args := []string{"build"}
	args = append(args, "-o", appName)
	build_cmd := exec.Command("go", args...)
	build_cmd.Stderr = &stderr
	err = build_cmd.Run()
	if err != nil {
		log.Printf("Failed to build the application: %s\n", stderr.String())
		return
	}
	log.Println("build ", appName, " success")
	Restart(appName)
}
