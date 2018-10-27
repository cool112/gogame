package http

import (
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/codec"
	_ "github.com/davyxu/cellnet/codec/httpjson"
	_ "github.com/davyxu/cellnet/codec/json"
	"github.com/davyxu/cellnet/peer/http"
	"github.com/cool112/gogame/proto"
	"github.com/cool112/gogame/proto/common"
	"github.com/cool112/gogame/server/service/logcellect"
	"reflect"
)

type HttpLog struct {
	common.LogCommon `bson:",inline"`
}

func init() {
	cellnet.RegisterHttpMeta(&cellnet.HttpMeta{
		Path:          "/log/onelog",
		Method:        "POST",
		RequestCodec:  codec.GetCodec("httpjson"),
		RequestType:   reflect.TypeOf((*HttpLog)(nil)).Elem(),
		ResponseCodec: codec.GetCodec("httpjson"),
		ResponseType:  reflect.TypeOf((*HttpLog)(nil)).Elem(),
	})
}

func (self *HttpLog) Do(session cellnet.Session) {
	log.Infof("log: %v", *self)
	err := logcellect.PersistLog(self)
	if err == nil {
		session.Send(&http.MessageRespond{0, &HttpAck{0, ""}, "json"})
	} else {
		log.Errorf("%#v", err)
		session.Send(&http.MessageRespond{0, &HttpAck{1, err.Error()}, "json"})
	}
}

func (self *HttpLog) GetPid() int32 {
	return proto.HTTP_LOG
}
