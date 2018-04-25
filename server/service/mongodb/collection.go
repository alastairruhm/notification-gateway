package mongodb

import (
	"gopkg.in/mgo.v2"
)

// Collection ...
type Collection struct {
	db      *Database
	name    string
	Session *mgo.Collection
}

// Connect ...
func (c *Collection) Connect() {
	session := *c.db.session.C(c.name)
	c.Session = &session
}

// NewCollectionSession get new collection
func NewCollectionSession(name string) *Collection {
	var c = Collection{
		db:   newDBSession("guidor"),
		name: name,
	}
	c.Connect()
	return &c
}

// Close ...
func (c *Collection) Close() {
	service.Close(c)
}
