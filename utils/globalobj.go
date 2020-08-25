package utils

import (
	"encoding/json"
	"github.com/deanisty/zinx/ziface"
	"io/ioutil"
)

type GlobalObj struct {
	TcpServer ziface.IServer // 全局server对象
	Host string  // 服务器监听ip
	TcpPort int  // 监听端口号
	Name string  // 服务器名称
	Version string // zinx版本号

	MaxPacketSize uint32  // 数据包的最大值
	MaxConn int     // 当前服务器允许的最大连接数
}

// 定义一个全局对象
var GlobalObject *GlobalObj

// 读取用户配置
func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("../conf/zinx.json")
	if err != nil {
		panic(err)
	}
	// 解析配置到对象中
	err = json.Unmarshal(data, g)
	if err != nil {
		panic(err)
	}
}

func init() {
	// 初始化GlobalObject变量 设置一些默认值
	GlobalObject = &GlobalObj{
		Host:          "0.0.0.0",
		TcpPort:       7777,
		Name:          "ZinServerApp",
		Version:       "v0.4",
		MaxPacketSize: 4096,
		MaxConn:       1200,
	}
}
