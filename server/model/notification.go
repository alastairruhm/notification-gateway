package model

import (
	"time"

	"github.com/alastairruhm/notification-gateway/server/schema"
	"github.com/alastairruhm/notification-gateway/server/service/mongodb"
	"gopkg.in/mgo.v2/bson"
)

const (
	notificationCollectionName string = "notification"
)

// Database ...
type Notification struct {
	conn *mongodb.Collection
}

func newNotificationCollection() *mongodb.Collection {
	return mongodb.NewCollectionSession(notificationCollectionName)
}

// Create create database record
func (d *Database) Create(notification *schema.Notification) (*schema.Notification, error) {
	d.conn = newNotificationCollection()
	defer d.conn.Close()

	// set default mongodb ID  and created date
	notification.ID = bson.NewObjectId()

	// instance.Token = bson.NewObjectId().Hex()
	//notification.Token = util.GenerateToken()
	notification.CreatedAt = time.Now()
	// Insert database record to mongodb
	err := d.conn.Session.Insert(&notification)
	if err != nil {
		return nil, dbError(err)
	}
	return notification, nil
}
