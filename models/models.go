package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`
	First_Name *string `json:"first_name"`
	Last_Name *string `json:"last_name"`
	Password *string `json:"password"`
	Phone *string `json:"phone"`
	Token *string `json:"token"`
	Refresh_Token *string `json:"refresh_token"`
	Created_At *time.Time `json:"created_at"`
	Updated_At *time.Time `json:"updated_at"`
	User_ID *string
	UserCart []ProductUser
	Address_Details []Address
	Order_Statis []Order
}