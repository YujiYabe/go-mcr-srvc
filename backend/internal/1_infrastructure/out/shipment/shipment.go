package out

import (
	"context"
	"fmt"
)

type shipment struct{}

// NewToShipment ...
func NewToShipment() service.ToShipment {
	s := new(shipment)
	return s
}

// HandOver ...
func (s *Shipment) HandOver(ctx context.Context) error {
	return nil
}

// Logging ...
func (s *Shipment) Logging(ctx context.Context) error {

	fmt.Println(" ============================== ")
	fmt.Printf("%+v\n", ctx.Value(key))
	fmt.Println(" ============================== ")

	return nil
}
