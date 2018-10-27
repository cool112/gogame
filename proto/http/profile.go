package http

import (
	"errors"
	"fmt"
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/codec"
	_ "github.com/davyxu/cellnet/codec/httpform"
	"github.com/davyxu/cellnet/peer/http"
	"github.com/cool112/gogame/proto"
	"github.com/cool112/gogame/server/service/profile"
	"reflect"
)

type Profile struct {
	Type string
}

func init() {
	cellnet.RegisterHttpMeta(&cellnet.HttpMeta{
		Path:          "/debug/profile",
		Method:        "GET",
		RequestCodec:  codec.GetCodec("httpform"),
		RequestType:   reflect.TypeOf((*Profile)(nil)).Elem(),
		ResponseCodec: codec.GetCodec("httpform"),
		ResponseType:  reflect.TypeOf((*Profile)(nil)).Elem(),
	})
}

func (self *Profile) Do(session cellnet.Session) {
	var err error
	var msg = "suc"
	switch self.Type {
	case "dump":
		err = profile.HeapDump()
	case "pprof":
		err = profile.PProf()
		msg = profile.CpuFile() + "\n" + profile.MemFile()
	case "trace":
		err = profile.Trace()
		msg = profile.TraceFile()
		case "stopprof":
		err = profile.StopPProf()
		case "stoptrace":
		err = profile.StopTrace()
	default:
		err = errors.New("unknown type:" + self.Type)
	}
	if err != nil {
		session.Send(&http.HTMLText{200, fmt.Sprintf("%#v", err)})

	} else {
		session.Send(&http.HTMLText{200, msg})
	}
}

func (self *Profile) GetPid() int32 {
	return proto.HTTP_PROF
}
