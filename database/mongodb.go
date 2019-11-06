package database

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/shrikar007/01-mongo-example/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var databaseNotOpenError = errors.New("database is not open, please open database to continue")

const DbName = "dbname"
const CollectionName = "collectionname"
//const Key = "id"
//const Value = "request"

type MongoDB struct {
	client *mongo.Client
	open   bool

	Host string
	Port int
}

func (db *MongoDB) Open() error {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%v:%v", db.Host, db.Port))

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	db.client = client
	db.open = true

	return nil
}

func (db *MongoDB) Close() error {
	db.open = false

	err := db.client.Disconnect(context.Background())
	if err != nil {
		return errors.Wrap(err, "unable to close database. exiting.")
	}

	return nil
}

func (db *MongoDB) Save(req types.Request) error {
	if !db.open {
		return databaseNotOpenError
	}

	collection := db.client.Database(DbName).Collection(CollectionName)

	inserted, err := collection.InsertOne(context.TODO(), req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted : ", inserted.InsertedID)

	return nil
}

func (db *MongoDB) Retrieve() ([]types.Request, error) {


	if !db.open {
		return nil, databaseNotOpenError
	}

	var results []types.Request

	collection := db.client.Database(DbName).Collection(CollectionName)

	cur, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {
		var elem types.Request
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.TODO())

	return results, nil
}
//
//func (db *MongoDB) Delete(Id string) error {
//	if !db.open {
//		return databaseNotOpenError
//	}
//
//	collection := db.client.Database(DbName).Collection(CollectionName)
//
//	_, err := collection.DeleteOne(context.Background(), bson.M{Key: Id})
//	if err != nil {
//		return errors.Wrap(err, "unable to delete session data")
//	}
//
//	return nil
//}
