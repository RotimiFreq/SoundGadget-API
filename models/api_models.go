package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	Id          primitive.ObjectID `json:"id,omitempty"`
	ProductType string             `json:"producttype,omitempty"`
	Brand       string             `json:"brand,omitempty"`
	Name        string             `json:"name,omitempty"`
	Price       int                `json:"price ,omitempty"`

	//mixer details
	MixerDetails string `json:"mixerdetails,omitempty"`

	//Microphone
	PickupPattern string `json:"pickuppattern,omitempty"`
	MicType       string `json:"mictype,omitempty"`
	Wireless      bool   `json:"wireless,omitempty"`

	//Processors
	ProcessorType    string `json:"string,omitempty"`
	ProcessorDetails string `json:"processortype,omitempty"`

	//speaker details
	SpeakerPeakpower    string `json:"speakerpeakpower,omitempty"`
	SpeakerProgrampower string `json:"speakerprogrampower,omitempty"`
	SpeakerType         string `json:"speakertype,omitempty"`
	SpeakerAbout        string `json:"speakerabout,omitempty"`

	//amplifier details
	AmplifierType        string `json:"amplifiertype,omitempty"`
	AmplifierPowerRating string `json:"amplifierpowerrating,omitempty"`
	AmplifierAbout       string `json:"amplifierabout,omitempty"`
}
