
package main

import (
	"flag"
	"fmt"
	"github.com/owenliang/crontab/worker"
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
	// worker  -config ./worker.json
	// worker -h
	flag.StringVar(&confFile, "config", "./worker.json", "指定worker.json")
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
	if err = worker.InitConfig(confFile);err!=nil{
		goto ERR
	}

	// 初始化job
	if err = worker.InitJobMgr();err!=nil{
		goto ERR
	}

	for {
		time.Sleep(1*time.Second)
	}

ERR:
	fmt.Println(err)



}