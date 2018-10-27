package http

import (
	"fmt"
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/codec"
	_ "github.com/davyxu/cellnet/codec/httpform"
	_ "github.com/davyxu/cellnet/codec/json"
	"github.com/davyxu/cellnet/peer/http"
	"github.com/cool112/gogame/proto"
	"reflect"
)

type TestParams struct {
	Msg string
	Id  int
}

func init() {
	cellnet.RegisterHttpMeta(&cellnet.HttpMeta{
		Path:          "/test/json",
		Method:        "GET",
		RequestCodec:  codec.GetCodec("httpform"),
		RequestType:   reflect.TypeOf((*TestParams)(nil)).Elem(),
		ResponseCodec: codec.GetCodec("httpform"),
		ResponseType:  reflect.TypeOf((*TestParams)(nil)).Elem(),
	})
}

func (self *TestParams) Do(session cellnet.Session) {
	session.Send(&http.MessageRespond{0, &TestParams{"suc", self.Id}, "json"})
}

func (self *TestParams) GetPid() int32 {
	return proto.HTTP_TEST
}

func (self *TestParams) String() string {
	return fmt.Sprintf("%s-%d", self.Msg, self.Id)
}
