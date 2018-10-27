package logcellect

import (
	baseMgo "github.com/globalsign/mgo"
	"github.com/cool112/gogame/mgo"
	"github.com/davyxu/golog"
)

var log = golog.New("logcellect")
var logDBC = &mgo.DBC{"","test","clientlog"}

//持久化日志
func PersistLog(l ...interface{}) (err error) {
	log.Debugln("start presist log")
	err=logDBC.Insert(l...) 
	return
}

func GetAll() (*baseMgo.Query, *baseMgo.Session){
	return logDBC.Find(nil)
}
