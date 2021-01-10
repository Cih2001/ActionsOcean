package model

import (
	"errors"
	"os"
	"testing"
)

func connectToLocalDB() error {
	if mongoClient != nil {
		return nil
	}
	username := os.Getenv("DBUSERNAME")
	password := os.Getenv("DBPASSWORD")
	if username == "" || password == "" {
		return errors.New("we are in build server")
	}

	var err error
	mongoClient, err = connectToDB("mongodb://localhost:27017", username, password)
	return err
}

func TestProjectFind(t *testing.T) {
	if err := connectToLocalDB(); err != nil {
		// we are in build server.
		return
	}

	project := ProjectModel{}
	err := project.Find("1231235")
	if err == nil {
		t.Fatal("a document was find with empty id")
	}
}

func TestProject(t *testing.T) {
	if err := connectToLocalDB(); err != nil {
		// we are in build server.
		return
	}

	project := ProjectModel{
		Name:            "interstellar",
		OwenerID:        "fake id",
		State:           "fake state",
		Progress:        100,
		ParticipantsIDs: []string{"fake id1"},
	}
	id, err := project.Create()
	if err != nil {
		t.Fatal(err)
	}

	tmp := ProjectModel{}
	err = tmp.Find(id.Hex())
	if err != nil {
		t.Fatal(err)
	}

	err = tmp.Delete(tmp.ID.Hex())
	if err != nil {
		t.Fatal(err)
	}

}
