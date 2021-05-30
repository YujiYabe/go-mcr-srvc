package shelf

import (
	"context"

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
	uri := "mongodb://localhost:27017"

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

// GetVegetables ...
func (s *Shelf) GetVegetables(ctx context.Context, items map[string]int) error {
	// for item, num := range items {
	// 	res := s.Conn.
	// 		Table("vegetables").
	// 		Where("name IN (?)", item).
	// 		UpdateColumn("stock", gorm.Expr("stock - ?", num))

	// 	if res.Error != nil {
	// 		return res.Error
	// 	}
	// }

	return nil
}
