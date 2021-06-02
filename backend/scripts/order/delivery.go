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
			Combos:    []*delivery.Combo{},
			SideMenus: []*delivery.SideMenu{},
			Drinks:    []*delivery.Drink{},
			Hamburgers: []*delivery.Hamburger{
				{
					Top:     1,
					Cheese:  1,
					Lettuce: 1,
					Tomato:  1,
				},
			},
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
