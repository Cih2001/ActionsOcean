package model

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client
var dbname string = "actionsocean"                      // TODO: should be in a config file
var defaultdbaddress string = "mongodb://mongodb:27017" // TODO: should be in a config file

// connectToDB tries to connect to the MongoDB.
func connectToDB(dbaddress string, username string, password string) (*mongo.Client, error) {
	if dbaddress == "" {
		dbaddress = defaultdbaddress
	}

	clientOptions := options.Client()

	// set the credentials
	cred := options.Credential{
		Username:      username,
		Password:      password,
		AuthMechanism: "SCRAM-SHA-1", // TODO: should be in a config file
	}
	clientOptions.SetAuth(cred)

	// set the connection URI
	clientOptions.ApplyURI(dbaddress)

	attempts := 0
	maxAttemps := 3            // TODO: should be in a config file
	timeout := 2 * time.Second // TODO: should be in a config file

	var maxAttempsReached = func() bool {
		attempts++
		time.Sleep(timeout)
		if attempts > maxAttemps {
			return true
		}
		return false
	}

	var client *mongo.Client
	var err error

	for {
		// connect to db
		if client, err = mongo.NewClient(clientOptions); err != nil {
			if maxAttempsReached() {
				return nil, err
			}
			// an error happened, let's try again
			continue
		}

		if err = client.Connect(context.Background()); err != nil {
			if maxAttempsReached() {
				return nil, err
			}
			// an error happened, let's try again
			continue
		}

		// ping the db
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		if err = client.Ping(ctx, readpref.Primary()); err != nil {
			fmt.Println("Could not ping the database")
			if maxAttempsReached() {
				return nil, err
			}
		} else {
			// successfully pinged the database. let's break the try loop
			// and return the client instance.
			break
		}
	}

	return client, nil
}

// InitializeDB initializes DB connection.
func InitializeDB(username string, password string) {
	var err error
	mongoClient, err = connectToDB("", username, password)
	if err != nil {
		panic(err)
	}
}
