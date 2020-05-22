package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"
	"tynmarket/coffeehub-go/controller/api"
	pb "tynmarket/coffeehub-go/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	arg := flag.Arg(0)

	fmt.Printf("\narg: %s\n\n", arg)

	if arg == "grpc" {
		fmt.Println("Run gRPC server")
		runGrpcServer()
	} else if arg == "client" {
		fmt.Println("Run gRPC client")
		runClient()
	} else {
		runGinServer()
	}
}

func runGinServer() {
	r := gin.Default()

	apiV1 := r.Group("/api")
	apiV1.GET("/coffees", api.Coffees)
	apiV1.GET("/coffees/roast/:roast", api.CoffeesRoast)

	r.Run()
}

type server struct {
	pb.CoffeeServerImpl
}

func runGrpcServer() {
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCoffeeProtoServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func runClient() {
	address := "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewCoffeeProtoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//coffeeID := "1"
	//r, err := c.GetCoffee(ctx, &pb.GetCoffeeRequest{CoffeeId: coffeeID})

	r, err := c.GetCoffees(ctx, &pb.GetCoffeesRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	coffees := r.GetCoffees()

	if coffees != nil {
		coffee := coffees[0]
		name := fmt.Sprintf("%s %s", coffee.Country, coffee.Area)
		log.Printf("name: %s", name)
	}
}
