package location

import "github.com/core-go/search"

type LocationSM struct {
	*search.SearchModel
	Id          string `json:"id,omitempty" gorm:"column:id;primary_key" bson:"_id,omitempty" dynamodbav:"id,omitempty" firestore:"-"`
	Type        string `json:"type,omitempty" gorm:"column:type" bson:"type,omitempty" dynamodbav:"type,omitempty" firestore:"type,omitempty" validate:"required,max=40"`
	Name        string `json:"name,omitempty" gorm:"column:name" bson:"name,omitempty" dynamodbav:"name,omitempty" firestore:"name,omitempty" validate:"required,max=255"`
	Description string `json:"description,omitempty" gorm:"column:description" bson:"description,omitempty" dynamodbav:"description,omitempty" firestore:"description,omitempty"`
	CustomUrl   string `json:"customUrl,omitempty" gorm:"column:urlId" bson:"customUrl,omitempty" dynamodbav:"customUrl,omitempty" firestore:"customUrl,omitempty"`
}
