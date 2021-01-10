package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProjectModel struct {
	ID             primitive.ObjectID `bson:"_id, omitempty"`
	Name           string
	OwenerID       string
	OwenerName     string
	State          string
	Progress       uint
	ParticipantsID []string
}
