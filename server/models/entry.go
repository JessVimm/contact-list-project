// Model for DB
package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Entry struct {
	ID    primitive.ObjectID `bson:"id"`
	Name  *string            `json:"name"`
	Age   *int               `json:"age"`
	City  *string            `json:"city"`
	Phone *string            `json:"phone"`
	Email *string            `json:"email"`
}
