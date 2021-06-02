package out

import (
	"context"
	"fmt"

	"app/internal/4_domain/domain"
)

type shipment struct{}

// NewToShipment ...
func NewToShipment() service.ToShipment {
	s := new(shipment)
	return s
}

// HandOver ...
func (s *Shipment) HandOver(ctx context.Context, order *domain.Order) error {
	return nil
}

// Logging ...
func (s *Shipment) Logging(ctx context.Context, order *domain.Order) error {

	fmt.Println(" ============================== ")
	fmt.Printf("%+v\n", order.OrderNumber)
	fmt.Println(" ============================== ")

	return nil
}
