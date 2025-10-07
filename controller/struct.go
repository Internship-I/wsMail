package controller

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	ID             	primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ConsigmentNote 	string             `bson:"connote,omitempty" json:"connote,omitempty"`
	SenderName     	string             `bson:"sender_name,omitempty" json:"sender_name,omitempty"`
	SenderPhone    	string             `bson:"sender_phone,omitempty" json:"sender_phone,omitempty"`
	ReceiverName   	string             `bson:"receiver_name,omitempty" json:"receiver_name,omitempty"`
	AddressReceiver string 			   `bson:"address_receiver,omitempty" json:"address_receiver,omitempty"`
	ReceiverPhone   string             `bson:"receiver_phone,omitempty" json:"receiver_phone,omitempty"`
	ItemContent    	string             `bson:"item_content,omitempty" json:"item_content,omitempty"`
	DeliveryStatus 	string             `bson:"delivery_status,omitempty" json:"delivery_status,omitempty"`
	CODValue       	float64            `bson:"cod_value,omitempty" json:"cod_value,omitempty"`
	CreatedAt      	primitive.DateTime `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt      	primitive.DateTime `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

type User struct {
	ID                     primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	FullName               string             `bson:"name,omitempty" json:"name,omitempty"`
	PhoneNumber            string             `bson:"phone,omitempty" json:"phone,omitempty"`
	Username               string             `bson:"username,omitempty" json:"username,omitempty"`
	Password               string             `bson:"password,omitempty" json:"password,omitempty"`
	Role                   string             `bson:"role,omitempty" json:"role,omitempty"` // "admin" or "kurir"
}

type ReqTransactionTransaction struct {
	ID             	primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ConsigmentNote 	string             `bson:"connote,omitempty" json:"connote,omitempty"`
	SenderName     	string             `bson:"sender_name,omitempty" json:"sender_name,omitempty"`
	SenderPhone    	string             `bson:"sender_phone,omitempty" json:"sender_phone,omitempty"`
	ReceiverName   	string             `bson:"receiver_name,omitempty" json:"receiver_name,omitempty"`
	AddressReceiver string 			   `bson:"address_receiver,omitempty" json:"address_receiver,omitempty"`
	ReceiverPhone   string             `bson:"receiver_phone,omitempty" json:"receiver_phone,omitempty"`
	ItemContent    	string             `bson:"item_content,omitempty" json:"item_content,omitempty"`
	DeliveryStatus 	string             `bson:"delivery_status,omitempty" json:"delivery_status,omitempty"`
	CODValue       	float64            `bson:"cod_value,omitempty" json:"cod_value,omitempty"`
	CreatedAt      	primitive.DateTime `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt      	primitive.DateTime `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

type ReqUser struct {
	ID                     primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	FullName               string             `bson:"name,omitempty" json:"name,omitempty"`
	PhoneNumber            string             `bson:"phone,omitempty" json:"phone,omitempty"`
	Username               string             `bson:"username,omitempty" json:"username,omitempty"`
	Password               string             `bson:"password,omitempty" json:"password,omitempty"`
	Role                   string             `bson:"role,omitempty" json:"role,omitempty"` // "admin" or "kurir"
}