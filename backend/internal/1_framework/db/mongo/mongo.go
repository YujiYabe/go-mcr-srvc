package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"backend/internal/2_adapter/gateway"
	"backend/pkg"
)

var isEnable = false

type (
	// Mongo ...
	Mongo struct {
		Conn *mongo.Client
	}

	// Stock ...
	Stock struct {
		ID    int
		Name  string
		Stock int
	}
)

// NewToMongo ...
func NewToMongo() gateway.ToMongo {
	conn, err := open(30)
	ctx := context.Background()
	if err != nil {
		pkg.Logging(ctx, err)
		panic(err)
	}

	s := new(Mongo)
	s.Conn = conn
	return s
}

func open(count uint) (
	client *mongo.Client,
	err error,
) {
	if !isEnable {
		return client, nil
	}

	uri := "mongodb://user:user@mongo:27017"
	ctx := context.Background()

	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		pkg.Logging(ctx, err)
		panic(err)
	}

	return client, nil
}

// UpdateBans ...
func (receiver *Mongo) UpdateBans(ctx context.Context, items map[string]int) error {
	bans := receiver.Conn.Database(pkg.MongoDatabase).Collection("bans")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for item, num := range items {

		filter := bson.M{"name": item}
		stock := &Stock{}

		err := bans.FindOne(ctx, filter).Decode(stock)
		if err != nil {
			pkg.Logging(ctx, err)
			return err
		}

		change := bson.M{"$set": bson.M{"stock": stock.Stock - num}}
		_, err = bans.UpdateOne(ctx, filter, change)
		if err != nil {
			pkg.Logging(ctx, err)
			return err
		}

		// 作業時間を擬似的に再現
		time.Sleep(1 * time.Second)
	}

	return nil
}
