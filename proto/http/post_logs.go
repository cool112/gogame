package http

import (
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/codec"
	_ "github.com/davyxu/cellnet/codec/httpjson"
	"github.com/davyxu/cellnet/peer/http"
	"github.com/cool112/gogame/proto"
	"github.com/cool112/gogame/server/service/logcellect"
	"reflect"
)

type HttpLogs struct{
	Logs []HttpLog
}

func init() {
	cellnet.RegisterHttpMeta(&cellnet.HttpMeta{
		Path:          "/log/batchlog",
		Method:        "POST",
		RequestCodec:  codec.GetCodec("httpjson"),
		RequestType:   reflect.TypeOf((*HttpLogs)(nil)).Elem(),
		ResponseCodec: codec.GetCodec("httpjson"),
		ResponseType:  reflect.TypeOf((*HttpLogs)(nil)).Elem(),
	})
}

func (self *HttpLogs) Do(session cellnet.Session) {
	log.Infof("log: %v", *self)
	var args =make([]interface{},len(self.Logs))
	for i,log:=range self.Logs{
		args[i] = &log
	}
	err := logcellect.PersistLog(args...)
	if err == nil {
		session.Send(&http.MessageRespond{0, &HttpAck{0, ""}, "json"})
	} else {
		log.Errorf("%#v", err)
		session.Send(&http.MessageRespond{0, &HttpAck{1, err.Error()}, "json"})
	}
}

func (self *HttpLogs) GetPid() int32 {
	return proto.HTTP_LOGS
}