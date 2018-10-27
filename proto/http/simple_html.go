package http

import (
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/codec"
	_ "github.com/davyxu/cellnet/codec/httpform"
	_ "github.com/davyxu/cellnet/codec/json"
	"github.com/davyxu/cellnet/peer/http"
	"github.com/cool112/gogame/proto"
	"reflect"
)

type TestHtml struct {
}

func init() {
	cellnet.RegisterHttpMeta(&cellnet.HttpMeta{
		Path:          "/test/html",
		Method:        "GET",
		RequestCodec:  codec.GetCodec("httpform"),
		RequestType:   reflect.TypeOf((*TestHtml)(nil)).Elem(),
		ResponseCodec: codec.GetCodec("httpform"),//响应是客户端解析需要,通常是不用配置
		ResponseType:  reflect.TypeOf((*TestHtml)(nil)).Elem(),
	})
}

func (self *TestHtml) Do(session cellnet.Session) {
	session.Send(&http.HTMLRespond{200, "test", struct {
		Name string
		Age  int
	}{"seg", 30}})
}

func (self *TestHtml) GetPid() int32 {
	return proto.HTTP_TEST
}
