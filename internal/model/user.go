package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	Id       primitive.ObjectID `bson:"_id"`
	Uuid     string     	`bson:"uuid"`
	Username string     	`bson:"username"`
	Email    string     	`bson:"email"`
	CreateAt time.Time  `bson:"create_at"`
	UpdateAt time.Time 	`bson:"updated_at"`
}