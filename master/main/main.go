package main

import (
	"flag"
	"fmt"
	"github.com/owenliang/crontab/master"
	"runtime"
	"time"
)


var (
	// 全局变量
	confFile string // 配置文件路径
)

func initEnv(){
	// 线程数量和CPU相同
	runtime.GOMAXPROCS(runtime.NumCPU())

}


// 解析命令行参数
func initArgs(){
	// master -config ./master.json
	// master -h
	flag.StringVar(&confFile, "config", "./master.json", "指定master.json")
	flag.Parse()

}


func main(){
	var (
		err error

	)


	// 初始化命令行参数
	initArgs()


	// 初始化线程
	initEnv()

	// 加载配置
	if err = master.InitConfig(confFile);err!=nil{
		goto ERR
	}

	// 日志管理器
	if err = master.InitLogMgr();err!=nil{
		goto ERR
	}

	// 任务管理器
	if err = master.InitJobMgr();err!=nil{
		goto ERR
	}

	// 启动Api HTTP服务
	if err=master.InitApiServer();err!=nil{
		goto ERR

	}

	for {
		time.Sleep(1*time.Second)
	}

	ERR:
		fmt.Println(err)



}

