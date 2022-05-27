package framework

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	mongo2 "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoConnection struct {
	Name         string
	URL          string
	client       *mongo.Client
	connected    bool
	remoteDbName string
}

type MongoConnectionDescription struct {
	Name        string
	EnvVarName  string
	Description string
	CanFail     bool
}

func (m *MongoConnection) init() error {
	opts := options.Client().ApplyURI(m.URL)
	m.deduceRemoteDbName()
	fmt.Println("REMOTE DB NAME: ", m.remoteDbName)
	client, err := mongo.NewClient(opts)
	if err != nil {
		panic(err)
	}
	m.client = client
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return err
	}
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println(err)
		return errors.New("[MONGO DB CONNECTION] Unable to connect to mongo database: " + m.Name)
	}
	fmt.Println("Connected! " + m.Name)
	m.connected = true
	return nil
}

func (m *MongoConnection) deduceRemoteDbName() {
	urlTokens := strings.Split(m.URL, "?")
	pathSegments := strings.Split(urlTokens[0], "/")
	m.remoteDbName = pathSegments[len(pathSegments)-1]
}

func (m *MongoConnection) DB() *mongo2.Database {
	return m.client.Database(m.remoteDbName)
}

func (m *MongoConnection) Connected() bool {
	return m.connected
}
