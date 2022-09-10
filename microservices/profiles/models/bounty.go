package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Bounty struct {
	ObjectID primitive.ObjectID
	Name     string
	Project  string
	Reward   int
}
