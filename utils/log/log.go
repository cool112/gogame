//author:seg
//date:2018-10-24
//日志封装 提供分级文件输出流
package log

import (
	"github.com/davyxu/golog"
	"io"
	"os"
	"runtime"
)

var log = golog.New("logutil")
var infoFile, errFile io.Writer

type LvOutput struct {
	logger *golog.Logger
}

func (self *LvOutput) Write(p []byte) (n int, err error) {
	curLv := self.logger.CurLevel()
	if curLv > golog.Level_Info {
		return errFile.Write(p)
	} else {
		return infoFile.Write(p)
	}
}

//初始化golog
func init() {
	mode := os.O_RDWR | os.O_CREATE | os.O_APPEND

	fileHandle, err := os.OpenFile("./log/info.out", mode, 0666)
	if err != nil {
		panic(err)
	}
	runtime.SetFinalizer(fileHandle, closeFile)
	infoFile = io.MultiWriter(os.Stdout, fileHandle)
	fileHandle, err = os.OpenFile("./log/error.out", mode, 0666)
	if err != nil {
		panic(err)
	}
	runtime.SetFinalizer(fileHandle, closeFile)
	errFile = io.MultiWriter(os.Stdout, fileHandle)
}

func New(name string) *golog.Logger {
	log := golog.New(name)
	log.SetOutput(&LvOutput{log})
	return log
}

func closeFile(fd *os.File) {
	if fd == nil {
		return
	}
	err := fd.Close()
	if err != nil {
		log.Errorf("log file close fail!%#v\n", err)
	}
}
