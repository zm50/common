package dbcli

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoCli *mongo.Client

func InitMongo() error {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return errors.WithMessage(err, "Failed to connect to MongoDB")
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return errors.WithMessage(err, "Failed to ping MongoDB")
	}

	mongoCli = client

	return nil
}

func Mdb() *mongo.Client {
	return mongoCli
}
