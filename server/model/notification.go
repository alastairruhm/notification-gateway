package model

import (
	"github.com/alastairruhm/notification-gateway/server/schema"
	"github.com/alastairruhm/notification-gateway/server/service/mongodb"
	"gopkg.in/mgo.v2/bson"
)

const (
	notificationCollectionName string = "notification"
)

// Notification ...
type Notification struct {
	conn *mongodb.Collection
}

func NewNotificationCollection() *mongodb.Collection {
	return mongodb.NewCollectionSession(notificationCollectionName)
}

// Save create database record
func (n *Notification) Save(notification *schema.Notification) (*schema.Notification, error) {
	n.conn = NewNotificationCollection()
	defer n.conn.Close()

	// set default mongodb ID  and created date
	notification.ID = bson.NewObjectId()

	// instance.Token = bson.NewObjectId().Hex()
	//notification.Token = util.GenerateToken()
	//notification.CreatedAt = time.Now()
	// Insert database record to mongodb
	err := n.conn.Session.Insert(&notification)
	if err != nil {
		return nil, dbError(err)
	}
	return notification, nil
}

func (n *Notification) FindByID(id string) (*schema.Notification, error) {
	n.conn = NewNotificationCollection()
	defer n.conn.Close()

	var notification schema.Notification

	err := n.conn.Session.FindId(bson.ObjectIdHex(id)).One(&notification)

	if err != nil {
		return nil, dbError(err)
	}
	return &notification, nil
}
