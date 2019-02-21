package usercrudhandler

import (
	"github.com/satori/go.uuid"
)

type RestCreateReqDTO struct {
	Name         string  `json:"name" bson:"name"`
	Address      string  `json:"address" bson:"address"`
	AddressLine2 string  `json:"addressLine2" bson:"addressLine2"`
	URL          string  `json:"url" bson:"url"`
	Outcode      string  `json:"outcode" bson:"outcode"`
	Postcode     string  `json:"postcode" bson:"postcode"`
	Rating       float32 `json:"rating" bson:"rating"`
	TypeOfFood   string  `json:"typeOfFood" bson:"typeOfFood"`
}

type RestUpdateDTO struct {
	DBID	     uuid.UUID `json:"id" bson:"_id"`	
	Name         string  `json:"name" bson:"name"`
	Address      string  `json:"address" bson:"address"`
	AddressLine2 string  `json:"addressLine2" bson:"addressLine2"`
	URL          string  `json:"url" bson:"url"`
	Outcode      string  `json:"outcode" bson:"outcode"`
	Postcode     string  `json:"postcode" bson:"postcode"`
	Rating       float32 `json:"rating" bson:"rating"`
	TypeOfFood   string  `json:"typeOfFood" bson:"typeOfFood"`
}


type RestDeleteDTO struct {
	DBID	     string `json:"id" bson:"_id"`	
	}

