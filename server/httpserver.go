package server

import (
"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/peer"
	"github.com/davyxu/cellnet/proc"

	_ "github.com/davyxu/cellnet/peer/http"
	_ "github.com/davyxu/cellnet/proc/http"
	"github.com/cool112/gogame/server/handle"
	"errors"
)
type HttpConf struct{
	Name,Addr,TmplDir,RootDir string
	Queue cellnet.EventQueue
	QueueGroup cellnet.EventQueueGroup
	TmplSuff []string
	Handler handle.AcceptorHandler
}
//开始httpserver
func StartHttpServer(conf *HttpConf) (err error){
	if conf == nil{
		return errors.New("http conf is nil")
	}
	httpP := peer.NewGenericPeer("http.Acceptor", conf.Name, conf.Addr, conf.Queue, conf.QueueGroup).(cellnet.HTTPAcceptor) //http 短连接不支持eventqueue,因为逻辑是同步的
	httpP.SetTemplateExtensions(conf.TmplSuff)
	httpP.SetTemplateDir(conf.TmplDir)
	httpP.SetFileServe(".", conf.RootDir)
	proc.BindProcessorHandler(httpP, "http", conf.Handler.HandleAcceptor)
	handle.SetPeerWithName(conf.Name, httpP)
	httpP.Start() //http listen
	return 
}