package schema

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Notification ...
type Notification struct {
	ID    bson.ObjectId `json:"id" bson:"_id"`
	Param interface{}   `json:"param" bson:"param"`
	State string        `json:"state" bson:"state"`
	// Estimated Time of Arrival
	ETA       time.Time `json:"eta" bson:"eta"`
	Created   time.Time `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at"`
}

// String returns JSON string with database instance info
func (d Notification) String() string {
	return jsonMarshal(d)
}

// DeliveryResult ...
type DeliveryResult struct {
	ID string `json:"id" bson:"_id"`
	// notification id
	NID string `json:"nid" bson:"nid"`
	//Result    string
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// String returns JSON string with desensitized team info
func (t *DeliveryResult) String() string {
	return jsonMarshal(t)
}
