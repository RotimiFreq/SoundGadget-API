package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Mixers struct{
	Id  primitive.ObjectID `json:"id,omitempty"`
	Brand string `json:"brand,omitempty" validate: "required"` 
	Name string  `json:"name,omitempty" validate: "required"`
	Price int    `json:"price,omitempty" validate: "required"`
	MixerDetails  string `json:"mixerdetails,omitempty" validate: "required"` 

}


