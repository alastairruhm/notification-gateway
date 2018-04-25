package model

import (
	"github.com/teambition/gear"
	"gopkg.in/mgo.v2"
)

func dbError(err error) error {
	if err == nil {
		return nil
	}
	if err == mgo.ErrNotFound {
		return gear.ErrNotFound.WithMsg(err.Error())
	}
	return gear.ErrInternalServerError.From(err)
}

// Models ....
type Models struct {
	Notification *Notification
}

// Init ...
func (m *Models) Init() *Models {
	m.Notification = new(Notification)
	return m
}
