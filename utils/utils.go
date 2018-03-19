package utils

import (
	"log"
	"os"
	"path/filepath"
	"text/template"
	"fmt"
)

//模板解析
func Tmpl(text string, data interface{}) {
	t := template.New("Usage")
	template.Must(t.Parse(text))

	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err.Error())
	}
}

//获取环境变量 gopath
func GetGOPATH() string {
	go_path := os.Getenv("GOPATH")
	if go_path == "" {
		return go_path
	}
	return filepath.SplitList(go_path)[0]
}

func PrintErrorAndExit(message, errorTemplate string) {
	Tmpl(fmt.Sprintf(errorTemplate, message), nil)
	os.Exit(2)
}