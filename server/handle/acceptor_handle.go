package handle

import (
	"github.com/davyxu/cellnet"
	golog "github.com/cool112/gogame/utils/log"
)

var log = golog.New("jm_handle")

type AcceptorHandler interface {
	HandleAcceptor(e cellnet.Event)
	GetPeer() cellnet.GenericPeer
}

type DefaultAcceptorHandler struct {
	Peer string
}

func (self *DefaultAcceptorHandler) HandleAcceptor(ev cellnet.Event) {
	switch msg := ev.Message().(type) {
	// 有新的连接
	case *cellnet.SessionAccepted:
		log.Debugln("server accepted")
	// 有连接断开
	case *cellnet.SessionClosed:
		log.Debugln("session closed: ", ev.Session().ID())
	// 收到某个连接的ChatREQ消息
	default:
		log.Debugf("recv: %v", msg)
		HandleBusiness(ev)
	}
}

func (self *DefaultAcceptorHandler) GetPeer() cellnet.GenericPeer {
	return GetServerPeer(self.Peer)
}

var defHandler AcceptorHandler
var peers map[string]cellnet.GenericPeer

func init(){
	peers = make(map[string]cellnet.GenericPeer)
}
func GetServerPeer(name string) cellnet.GenericPeer {
	return peers[name]
}
func init() {
	defHandler = &DefaultAcceptorHandler{}
}

func SetAcceptorHandler(newHandler AcceptorHandler) {
	defHandler = newHandler
}

func SetPeer(newP cellnet.GenericPeer) {
	peers[""] = newP
}

func SetPeerWithName(name string, newP cellnet.GenericPeer){
	peers[name] = newP
}

func HandleAcceptor(ev cellnet.Event) {
	defHandler.HandleAcceptor(ev)
}

func StopAllPeer(){
	for _,peer:=range peers{
		peer.Stop()
	}
}
