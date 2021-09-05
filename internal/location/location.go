package location

import (
	"github.com/core-go/mongo/geo"
	"time"
)

type Location struct {
	Id          string        `json:"id,omitempty" gorm:"column:id;primary_key" bson:"_id,omitempty" dynamodbav:"id,omitempty" firestore:"-"`
	Type        string        `json:"type,omitempty" gorm:"column:type" bson:"type,omitempty" dynamodbav:"type,omitempty" firestore:"type,omitempty" validate:"required,max=40"`
	Name        string        `json:"name,omitempty" gorm:"column:name" bson:"name,omitempty" dynamodbav:"name,omitempty" firestore:"name,omitempty" validate:"required,max=255"`
	Description string        `json:"description,omitempty" gorm:"column:description" bson:"description,omitempty" dynamodbav:"description,omitempty" firestore:"description,omitempty"`
	CustomUrl   string        `json:"customUrl,omitempty" gorm:"column:urlId" bson:"customUrl,omitempty" dynamodbav:"customUrl,omitempty" firestore:"customUrl,omitempty"`
	Longitude   *float64      `json:"longitude,omitempty" gorm:"column:longitude" bson:"-" dynamodbav:"longitude,omitempty" firestore:"longitude,omitempty"`
	Latitude    *float64      `json:"latitude,omitempty" gorm:"column:latitude" bson:"-" dynamodbav:"latitude,omitempty" firestore:"latitude,omitempty"`
	Geo         *geo.JSON     `json:"-" bson:"geo,omitempty" gorm:"-" dynamodbav:"-" firestore:"-"`
	Info        *LocationInfo `json:"info,omitempty" bson:"-" gorm:"column:info" dynamodbav:"info,omitempty" firestore:"info,omitempty"`
	CreatedBy   string        `json:"createdBy,omitempty" gorm:"column:createdby" bson:"createdBy,omitempty" dynamodbav:"createdBy,omitempty" firestore:"createdBy,omitempty"`
	CreatedAt   *time.Time    `json:"createdAt,omitempty" gorm:"column:createdat" bson:"createdAt,omitempty" dynamodbav:"createdAt,omitempty" firestore:"-"`
	UpdatedBy   string        `json:"updatedBy,omitempty" gorm:"column:updatedby" bson:"updatedBy,omitempty" dynamodbav:"updatedBy,omitempty" firestore:"updatedBy,omitempty"`
	UpdatedAt   *time.Time    `json:"updatedAt,omitempty" gorm:"column:updatedat" bson:"updatedAt,omitempty" dynamodbav:"updatedAt,omitempty" firestore:"-"`
	Version     int           `json:"version,omitempty" gorm:"column:version" bson:"version,omitempty" dynamodbav:"version,omitempty" firestore:"version,omitempty"`
}
