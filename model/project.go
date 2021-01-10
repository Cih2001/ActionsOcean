package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProjectModel struct {
	ID              primitive.ObjectID `bson:"_id, omitempty"`
	Name            string
	OwenerID        string
	State           string
	Progress        uint
	ParticipantsIDs []string
}

const (
	ProjectsCollectionName = "projects" // TODO: can be in a config file if not fixed
)

// Find finds a project given its id
func (model *ProjectModel) Find(id string) error {
	if id == "" {
		return ErrEmptyID
	}
	collection := mongoClient.Database(dbname).Collection(ProjectsCollectionName)
	objID, _ := primitive.ObjectIDFromHex(id)
	err := collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(model)
	if err != nil {
		return err
	}
	return nil
}

// Create creates a new project and returns its id
func (model *ProjectModel) Create() (primitive.ObjectID, error) {
	collection := mongoClient.Database(dbname).Collection(ProjectsCollectionName)
	model.ID = primitive.NewObjectID()
	result, err := collection.InsertOne(context.Background(), model)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

// Delete deletes a project given its id
func (model *ProjectModel) Delete(id string) error {
	if id == "" {
		return ErrEmptyID
	}
	collection := mongoClient.Database(dbname).Collection(ProjectsCollectionName)
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return err
	}
	return nil
}
