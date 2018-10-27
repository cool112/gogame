package proto

import (
	"github.com/davyxu/cellnet"
)

type Executable interface {
	//执行
	Do(cellnet.Session)
	//协议号
	GetPid() int32
}

//modules
const (
	MOD_ACC = 1000 * (iota + 1)
	MOD_LOG
	MOD_HTTP
)

//acc module
const (
	LOGIN = MOD_ACC + iota + 1
	AUTH
	ACK
)

//log module
const (
	LOG_COMMON = MOD_LOG + iota + 1
)

const (
	HTTP_TEST = MOD_HTTP + iota + 1
	HTTP_THTML
	HTTP_PATCH
	HTTP_PROF
	HTTP_TRACE
	HTTP_LOG_LIST
	HTTP_LOG
	HTTP_LOGS
)
