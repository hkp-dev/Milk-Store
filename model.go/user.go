package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	FullName    string             `bson:"full_name" json:"full_name"`
	PhoneNumber string             `bson:"phone" json:"phone_number"`
	OTP         string             `bson:"otp" json:"otp"`
	CreatedAt   time.Time          `bson:"create_time" json:"create_time"`
	Gender      string             `bson:"gender" json:"gender"`
}
