package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"time"
)

var (
	currpath string //当前运行目录
)

func init() {
	currpath, _ = os.Getwd()
}

//获取文件夹列表
func getDirList(dirpath string) ([]string, error) {
	var dir_list []string
	f, dir_err := ioutil.ReadDir(dirpath)
	for _, file_info := range f {
		if file_info.IsDir() {
			dir_list = append(dir_list, file_info.Name())
		}
	}
	return dir_list, dir_err
}

//获取文件修改时间
func GetFileModTime(path string) (int64, error) {
	f, err := os.Open(path)
	if err != nil {
		return time.Now().Unix(), err
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return time.Now().Unix(), err
	}

	return fi.ModTime().Unix(), nil
}

//获取应用名称
func GetAppName() string {
	return filepath.Base(currpath)
}

//获取运行目录
func GetRunPath() string {
	return currpath
}
