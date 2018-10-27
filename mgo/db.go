package mgo

import (
	_mgo "github.com/globalsign/mgo"
)

type MgoDoc interface {
	MgoMarshal() string
}
type NilSesErr struct{
	name string
}

func (self *NilSesErr) Error() string{
	return "mgo Session not existed!name:" + self.name
}
//db+collection context
type DBC struct {
	Session    string
	Db         string
	Collection string
}

//插入doc
func (self *DBC) Insert(items ...interface{}) error {
	ses:=self.getSession()
	if ses == nil {
		return &NilSesErr{self.Session}
	}
	defer ses.Close()
	err := self.C(ses).Insert(items...)
	if err != nil {
		return err
	}
	return nil
}

func (self *DBC) C(ses *_mgo.Session) *_mgo.Collection{
	return ses.DB(self.Db).C(self.Collection)
}

func (self *DBC) DB(ses *_mgo.Session) *_mgo.Database{
	return ses.DB(self.Db)
}

func (self *DBC) getSession() *_mgo.Session{
	var ses *_mgo.Session
	if len(self.Session) != 0 {
		ses = Session(self.Session)
	} else {
		ses = DefaultSession()
	}
	return ses
}
//查询结果
func (self *DBC) Find(cond interface{}) (*_mgo.Query, *_mgo.Session){
	ses:=self.getSession()
	if ses == nil{
		return nil, nil
	}
	query:=self.C(ses).Find(cond)
	return query, ses
	
}
