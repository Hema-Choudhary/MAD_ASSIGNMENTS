package usercrudhandler

import (
	"github.com/satori/go.uuid"
)


type RestGetRespDTO struct {
	DBID         uuid.UUID  `json:"id" bson:"_id"`
	Name         string  `json:"name" bson:"name"`
	Address      string  `json:"address" bson:"address"`
	AddressLine2 string  `json:"addressLine2" bson:"addressLine2"`
	URL          string  `json:"url" bson:"url"`
	Outcode      string  `json:"outcode" bson:"outcode"`
	Postcode     string  `json:"postcode" bson:"postcode"`
	Rating       float32 `json:"rating" bson:"rating"`
	TypeOfFood   string  `json:"typeOfFood" bson:"typeOfFood"`
}

type RestGetListRespDTO struct {
	Rests []RestGetRespDTO `json:"rests"`
	Count int              `json:"count"`
}

type RestCreateRespDTO struct {
	ID uuid.UUID `json:"id"`
}

type RestUpdateRespDTO struct {
	ID uuid.UUID `json:"id"`
}

type RestDeleteRespDTO struct {
	ID string `json:"id"`
}

