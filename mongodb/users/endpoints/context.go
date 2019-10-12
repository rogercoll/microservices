package endpoints

import (
	"gopkg.in/mgo.v2"
	"github.com/rogercoll/microservices/mongodb/users/config"
)

type Context struct {
	MongoSession *mgo.Session
}

func (c *Context) Close() {
	c.MongoSession.Close()
}

func (c *Context) DbCollection(name string) *mgo.Collection {
	return c.MongoSession.DB(config.AppConfig.Database).C(name)
}

func NewContext() *Context {
	session := config.GetSession().Copy()
	context := &Context{
		MongoSession: session,
	}
	return context
}