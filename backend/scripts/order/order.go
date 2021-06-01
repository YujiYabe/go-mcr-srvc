package main

import (
	"context"
	"fmt"
	"time"

	grpc "google.golang.org/grpc"

	"order/delivery"
)

func main() {
	dial, err := grpc.Dial("localhost:3456", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return
	}
	defer dial.Close()

	conn := delivery.NewDeliveryServiceClient(dial)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var request = &delivery.DeliveryRequest{
		Order: &delivery.Order{
			Hamburger: &delivery.Hamburger{
				Top: 1,
			},
		},
	}

	res, err := conn.DeliveryRPC(ctx, request)
	if err != nil {
		return
	}
	fmt.Println(" ============================== ")
	fmt.Printf("%+v\n", res.String())
	fmt.Println(" ============================== ")

	return
}
