package service_clients

import (
	"fmt"
	"gokit/ecommerse/products/pkg/grpc/pb"

	"google.golang.org/grpc"
)

type ProductClientConn interface {
	GetProductsClient() pb.ProductsClient
}

type Clients struct {
	ProductClient pb.ProductsClient
}

func NewClientConn() ProductClientConn {
	return &Clients{}

}

func (clients Clients) GetProductsClient() pb.ProductsClient {
	conn, err := grpc.Dial("172.18.0.8:8082", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Failed to start gRPC connection: %v", err)
	}
	//defer conn.Close()

	return pb.NewProductsClient(conn)
}
