package shelf

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"backend/internal/2_adapter/service"
)

type (
	// Shelf ...
	Shelf struct {
		Conn *mongo.Client
	}

	// Stock ...
	Stock struct {
		ID    int
		Name  string
		Stock int
	}
)

// NewToShelf ...
func NewToShelf() service.ToShelf {
	conn, err := open(30)
	if err != nil {
		panic(err)
	}

	s := new(Shelf)
	s.Conn = conn
	return s
}

func open(count uint) (*mongo.Client, error) {
	uri := "mongodb://user:user@mongo:27017"

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	return client, nil
}

// Dummy ...
func (s *Shelf) Dummy(ctx context.Context) (string, error) {
	return "dummy ok", nil
}

// GetBans ...
func (s *Shelf) GetBans(ctx context.Context, items map[string]int) error {
	bans := s.Conn.Database("app").Collection("bans")

	for item, num := range items {
		filter := bson.M{"name": item}
		stock := &Stock{}

		err := bans.FindOne(ctx, filter).Decode(stock)
		if err != nil {
			return err
		}

		change := bson.M{"$set": bson.M{"stock": stock.Stock - num}}
		_, err = bans.UpdateOne(ctx, filter, change)
		if err != nil {
			return err
		}
	}

	return nil
}
