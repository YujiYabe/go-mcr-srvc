package main

import (
	"context"
	"fmt"
	"order/delivery"
	"time"

	"google.golang.org/grpc"
)

func main() {
	dial, err := grpc.Dial("localhost:3456", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}
	defer dial.Close()

	conn := delivery.NewDeliveryServiceClient(dial)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var request = &delivery.DeliveryRequest{
		Order: &delivery.Order{
			Combo:     []*delivery.Combo{},
			Hamburger: []*delivery.Hamburger{},
			SideMenu:  []*delivery.SideMenu{},
			Drink:     []*delivery.Drink{},
			// Hamburger: []*delivery.Hamburger{
			// 	{
			// 		Top:    1,
			// 		Cheese: 1,
			// 	},
			// },
		},
	}

	res, err := conn.DeliveryRPC(ctx, request)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}

	fmt.Printf("%+v\n", res.String())
	return
}
