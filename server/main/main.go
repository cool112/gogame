package main

import (
	"github.com/davyxu/cellnet"
	golog "github.com/cool112/gogame/utils/log"

	"github.com/magiconair/properties"
	"github.com/cool112/gogame/mgo"
	_ "github.com/cool112/gogame/proto"
	_ "github.com/cool112/gogame/proto/common"
	_ "github.com/cool112/gogame/proto/http"
	"github.com/cool112/gogame/server/handle"
	"os"
	"os/signal"
	"github.com/cool112/gogame/server"
)

var log = golog.New("jmserver")

func main() {
	//读取配置1
	config := properties.MustLoadFile("./conf/config", properties.UTF8)

	//连接数据库
	mgoUrl := config.MustGet("mgoUrl")
	err := mgo.CreateDefaultSession(mgoUrl)
	if err != nil {
		log.Errorln(err)
	}
	//公共事件队列组
	queueGroup := cellnet.NewEventGroup(0)
	//tcp监听,使用默认handler
	addr := config.MustGet("port")
	server.StartDefaultTcpServer("jmserver", addr, nil, queueGroup)

	//http监听,自定义handler
	httpAddr := config.MustGet("httpPort")
	conf := &server.HttpConf{
		"jmhttpserver",
		httpAddr,
		"./tmpl",
		"./res",
		nil, nil,
		[]string{".tmpl"},
		&handle.DefaultAcceptorHandler{"jmhttpserver"},
	}
	server.StartHttpServer(conf)
	// 事件队列开始循环
	queueGroup.StartLoop()

	// 中断等待
	gracefullyShutdown(func() {
		queueGroup.StopLoop()
		queueGroup.Wait() //等待所有任务处理完成
		handle.StopAllPeer()//关闭所有连接
		mgo.CloseAll()
	})
}
//优雅关闭
func gracefullyShutdown(hook func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	log.Infoln("sys signal is:" + sig.String())
	hook()
}
