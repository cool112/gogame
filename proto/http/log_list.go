package http

import (
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/codec"
	_ "github.com/davyxu/cellnet/codec/httpform"
	"github.com/davyxu/cellnet/peer/http"
	"github.com/cool112/gogame/proto"
	"github.com/cool112/gogame/proto/common"
	"github.com/cool112/gogame/server/service/logcellect"
	"reflect"
)

type ListLog struct {
}

func init() {
	cellnet.RegisterHttpMeta(&cellnet.HttpMeta{
		Path:          "/log/list",
		Method:        "GET",
		RequestCodec:  codec.GetCodec("httpform"),
		RequestType:   reflect.TypeOf((*ListLog)(nil)).Elem(),
		ResponseCodec: codec.GetCodec("httpform"),
		ResponseType:  reflect.TypeOf((*ListLog)(nil)).Elem(),
	})
}

func (self *ListLog) Do(session cellnet.Session) {
	var logs []common.LogCommon
	query, ses := logcellect.GetAll()
	if ses != nil {
		defer ses.Close()//必须查询完才能关闭连接
		query.All(&logs)
	}
	session.Send(&http.HTMLRespond{200, "loglist", logs})
}

func (self *ListLog) GetPid() int32 {
	return proto.HTTP_LOG_LIST
}
