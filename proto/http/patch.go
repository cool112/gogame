package http

import (
	"fmt"
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/codec"
	_ "github.com/davyxu/cellnet/codec/httpform"
	"github.com/davyxu/cellnet/peer/http"
	"github.com/cool112/gogame/proto"
	"github.com/cool112/gogame/server/service/patch"
	"reflect"
)

type Patch struct {
	Path string
}

func init() {
	cellnet.RegisterHttpMeta(&cellnet.HttpMeta{
		Path:         "/debug/patch",
		Method:       "GET",
		RequestCodec: codec.GetCodec("httpform"),
		RequestType:  reflect.TypeOf((*Patch)(nil)).Elem(),
		ResponseCodec: codec.GetCodec("httpform"),
		ResponseType:  reflect.TypeOf((*Patch)(nil)).Elem(),
	})
}

func (self *Patch) Do(session cellnet.Session) {
	err := patch.RunPatch(&self.Path)
	if err != nil {
		session.Send(&http.HTMLText{200, fmt.Sprintf("%#v", err)})

	} else {
		session.Send(&http.HTMLText{200, "suc"})
	}
}

func (self *Patch) GetPid() int32 {
	return proto.HTTP_PATCH
}
