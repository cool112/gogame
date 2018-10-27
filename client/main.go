package main

import (
	"encoding/json"
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/peer"
	_ "github.com/davyxu/cellnet/peer/tcp"
	"github.com/davyxu/cellnet/proc"
	_ "github.com/davyxu/cellnet/proc/tcp"
	"github.com/davyxu/golog"
	"github.com/magiconair/properties"
	_ "github.com/cool112/gogame/proto"
	"github.com/cool112/gogame/proto/common"
	protohttp "github.com/cool112/gogame/proto/http"
	"net/http"
	"strings"
	"time"
)

var log = golog.New("jmclient")

func main() {
	config := properties.MustLoadFile("./conf/config", properties.UTF8)
	addr := config.MustGet("port")
	// 创建一个事件处理队列，整个客户端只有这一个队列处理事件，客户端属于单线程模型
	queue := cellnet.NewEventQueue()

	// 创建一个tcp的连接器，名称为client，连接地址为127.0.0.1:8801，将事件投递到queue队列,单线程的处理（收发封包过程是多线程）
	p := peer.NewGenericPeer("tcp.Connector", "jmclient", addr, queue, nil)

	// 设定封包收发处理的模式为tcp的ltv(Length-Type-Value), Length为封包大小，Type为消息ID，Value为消息内容
	// 并使用switch处理收到的消息
	proc.BindProcessorHandler(p, "tcp.ltv", func(ev cellnet.Event) {
		switch msg := ev.Message().(type) {
		case *cellnet.SessionConnected:
			log.Debugln("client connected")
		case *cellnet.SessionClosed:
			log.Debugln("client error")
		case *common.Ack:
			log.Infof("sid%d say: %d", msg.ProId, msg.ErrCode)
		}
	})

	// 开始发起到服务器的连接
	p.Start()

	// 事件队列开始循环
	queue.StartLoop()

	//	// 阻塞的从命令行获取聊天输入
	//	ReadConsole(func(str string) {
	//
	//		p.(interface {
	//			Session() cellnet.Session
	//		}).Session().Send(&proto.ChatREQ{
	//			Content: str,
	//		})
	//
	//	})
	log.Infoln("send")
	p.(interface {
		Session() cellnet.Session
	}).Session().Send(&common.LogCommon{"user", "mac", common.LogCommon_START_APP, "", time.Now().Unix(), "test", "tcp"})
	logone := &protohttp.HttpLog{LogCommon: common.LogCommon{"user2", "mac2", common.LogCommon_UPDATE_START, "", time.Now().Unix(), "app2", "http"}}
	jsonStr, err := json.Marshal(logone)
	if err != nil {
		log.Errorln(err.Error())
	}
	log.Infoln("json:" + string(jsonStr))
	body := strings.NewReader(string(jsonStr))
	http.Post("http://127.0.0.1:8889/log/onelog", "application/json", body)
	queue.Wait()
}
