package server

import (
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/peer"
	_ "github.com/davyxu/cellnet/peer/tcp"
	"github.com/davyxu/cellnet/proc"
	_ "github.com/davyxu/cellnet/proc/tcp"
	"github.com/cool112/gogame/server/handle"
)

//开启默认监听
func StartDefaultTcpServer(name, addr string, queue cellnet.EventQueue, qg cellnet.EventQueueGroup) (err error) {
	p := peer.NewGenericPeer("tcp.Acceptor", name, addr, queue, qg)
	proc.BindProcessorHandler(p, "tcp.ltv", handle.HandleAcceptor)
	handle.SetPeer(p)
	p.Start() //socket listen
	return nil
}

//开启默认监听
func StartTcpServer(name, addr string, handler handle.AcceptorHandler, queue cellnet.EventQueue, qg cellnet.EventQueueGroup) (err error) {
	p := peer.NewGenericPeer("tcp.Acceptor", name, addr, queue, qg)
	proc.BindProcessorHandler(p, "tcp.ltv", handler.HandleAcceptor)
	handle.SetPeerWithName(name, p)
	p.Start() //socket listen
	return nil
}
