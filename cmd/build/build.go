package build

import (
	"io"
	"strings"
)

type Command struct {
	//运行方法
	Run func(c *Command, args []string) int
	//使用方法
	UsageLine string
	//输出
	output io.Writer
	//
	Short string
	Help  string
}

var AvailableCommands []*Command

//获取命令名称
func (c *Command) Name() string {
	name := c.UsageLine
	i := strings.Index(name, " ")
	if i >= 0 {
		name = name[:i]
	}
	return name
}

//设置输出
func (c *Command) SetOutput(output io.Writer) {
	c.output = output
}

//是否可运行
func (c *Command) Runnable() bool {
	return c.Run != nil
}
