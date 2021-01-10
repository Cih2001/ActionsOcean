package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProjectModel struct {
	ID              primitive.ObjectID `json:"id" bson:"_id, omitempty"`
	Name            string             `json:"name"`
	OwenerID        string             `json:"owener_id"`
	State           string             `json:"state"`
	Progress        uint               `json:"progress"`
	ParticipantsIDs []string           `json:"participants_ids"`
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

func GetAllProjects() ([]ProjectModel, error) {
	result := []ProjectModel{}
	collection := mongoClient.Database(dbname).Collection(ProjectsCollectionName)
	cur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return result, err
	}

	err = cur.All(context.Background(), &result)
	return result, err
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
