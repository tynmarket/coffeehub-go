package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
	"tynmarket/coffeehub-go/controller/api"
	pb "tynmarket/coffeehub-go/proto"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	arg := flag.Arg(0)

	fmt.Printf("\narg: %s\n\n", arg)

	if arg == "gateway" {
		fmt.Println("Run gRPC gateway")
		runGrpcGateway()
	} else if arg == "grpc" {
		fmt.Println("Run gRPC server")
		runGrpcServer()
	} else if arg == "client" {
		fmt.Println("Run gRPC client")
		runClient()
	} else {
		fmt.Print("Run gin server\n\n")
		runGinServer()
	}
}

const (
	port     = ":8080"
	grpcPort = ":50051"
)

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

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", grpcPort, "gRPC server endpoint")
)

func runGrpcGateway() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterCoffeeProtoHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)

	if err != nil {
		log.Fatalf("failed to register endpoint: %v", err)
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func runGrpcServer() {
	lis, err := net.Listen("tcp", grpcPort)

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
	conn, err := grpc.Dial(grpcPort, grpc.WithInsecure(), grpc.WithBlock())
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
