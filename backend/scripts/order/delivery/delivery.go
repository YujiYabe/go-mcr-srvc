package main

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"

	"delivery/pb"
)

func main() {
	fmt.Println("1==============================")
	dial, err := grpc.Dial("localhost:3456", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}
	defer dial.Close()
	fmt.Println("2==============================")
	conn := pb.NewDeliveryServiceClient(dial)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	fmt.Println("3==============================")
	var request = &pb.DeliveryRequest{
		Order: &pb.Order{
			Combos:    []*pb.Combo{},
			SideMenus: []*pb.SideMenu{},
			Drinks:    []*pb.Drink{},
			Hamburgers: []*pb.Hamburger{
				{
					Top:     1,
					Cheese:  1,
					Lettuce: 1,
					Tomato:  1,
				},
			},
		},
	}
	fmt.Println("4==============================")
	res, err := conn.DeliveryRPC(ctx, request)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}
	fmt.Println("5==============================")
	fmt.Printf("%+v\n", res.String())
	return
}
