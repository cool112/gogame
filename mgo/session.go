package mgo

import (
	_mog "github.com/globalsign/mgo"
	"errors"
)

var sessions map[string]*_mog.Session
const(
	DEFAULT_SES_NAME = "default"
)
func init() {
 sessions = make(map[string]*_mog.Session)
}

//创建session
func CreateSession(url,name string) (err error) {
	ses, err := _mog.Dial(url)
	if err != nil{
		return
	}
	if len(name) == 0{
		name = DEFAULT_SES_NAME
	}
	
	if _, ok:=sessions[name];ok{
		return errors.New("Duplicated session exsited! name:"+name)
	}else{
		sessions[name] = ses
	}
	return
}
//创建默认session
func CreateDefaultSession(url string) (err error) {
	return CreateSession(url, "")
}

func DefaultSession() *_mog.Session{
	ses,ok := sessions[DEFAULT_SES_NAME]
	if ok{
		return ses.Copy()
	}else{
		return nil
	}
}

func Session(name string) *_mog.Session{
	ses,ok := sessions[name]
	if ok{
		return ses.Copy()
	}else{
		return nil
	}
}

func CloseAll(){
	for _,ses := range sessions{
		ses.Close()
	}
}

func Close(name string){
	ses := sessions[name]
	if ses != nil{
		ses.Close()
	}
}