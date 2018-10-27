package common

import (
	"github.com/davyxu/cellnet"
	"github.com/cool112/gogame/proto"
	"log"
	"github.com/cool112/gogame/server/service/logcellect"
)

func (pk *LogCommon) Do(session cellnet.Session) {
	log.Printf("log: %v", *pk)
	err:=logcellect.PersistLog(pk)
	if err == nil{
		session.Send(&Ack{0, int32(pk.GetPid())})
	}else{
		log.Printf("%#v", err)
	session.Send(&Ack{1, int32(pk.GetPid())})
	}
}

func (pk *LogCommon) GetPid() int32 {
	return proto.LOG_COMMON
}
