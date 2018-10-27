package handle

import (
	"github.com/davyxu/cellnet"
	"github.com/cool112/gogame/proto"
)

type BusinessHandler interface {
	HandleBusiness(ev cellnet.Event)
}

type DefaultBusinessHandler struct {
}

func (handler *DefaultBusinessHandler) HandleBusiness(ev cellnet.Event) {
	msg, ok := ev.Message().(proto.Executable)
	if !ok {
		if msg != nil {
			log.Warnf("maybe incomplete proto: %#v", ev.Message())
		}
		return
	}

	msg.Do(ev.Session())
}

var defBzHandler BusinessHandler

func init() {
	defBzHandler = &DefaultBusinessHandler{}
}

func HandleBusiness(ev cellnet.Event) {
	defBzHandler.HandleBusiness(ev)
}

func SetBzHandler(newHandler BusinessHandler) {
	defBzHandler = newHandler
}
