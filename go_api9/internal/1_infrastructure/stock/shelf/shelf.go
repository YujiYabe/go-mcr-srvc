package shelf

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"app/internal/2_adapter/service"
)

type (
	// Shelf ...
	Shelf struct {
		Conn *mongo.Client
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
	// uri := "mongodb+srv://<username>:<password>@<cluster-address>/test?w=majority"
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
		change := bson.M{"$set": bson.M{"stock": num}}
		res, err := bans.UpdateOne(ctx, filter, change)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("mongo==============================")
		debugTarget := res
		fmt.Printf("%#v\n", debugTarget)
		fmt.Println("==============================")

		// res := s.Conn.
		// 	Table("vegetables").
		// 	Where("name IN (?)", item).
		// 	UpdateColumn("stock", gorm.Expr("stock - ?", num))

		// if res.Error != nil {
		// 	return res.Error
		// }
	}

	return nil
}
