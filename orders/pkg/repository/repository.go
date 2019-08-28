package repository

import (
	"context"
	"fmt"
	"gokit/ecommerse/orders/pkg/database"
	"gokit/ecommerse/orders/pkg/models"
	"gokit/ecommerse/orders/pkg/service_clients"
	"gokit/ecommerse/products/pkg/grpc/pb"

	"time"
)

//OrderRepositoryInterface implimets all methods in OrderRepository
type OrderRepositoryInterface interface {
	CreateOrders(ctx context.Context, createReq models.Orders) (createResp *models.Orders, err error)
}

// OrderRepository **
type OrderRepository struct {
	mysqlInterface database.MysqlInterface
	clientConn     service_clients.ProductClientConn
}

//NewOrderRepository inject dependancies of DataStore
func NewOrderRepository(mysqlInterface database.MysqlInterface, clientConn service_clients.ProductClientConn) OrderRepositoryInterface {
	return &OrderRepository{mysqlInterface: mysqlInterface, clientConn: clientConn}
}

//Create add new record in datastore
func (orderRepository *OrderRepository) CreateOrders(ctx context.Context, order models.Orders) (*models.Orders, error) {

	dbConn := orderRepository.mysqlInterface.NewClientConnection()
	defer dbConn.Close()
	dbConn.AutoMigrate(&models.Orders{})
	now := time.Now().In(time.UTC)
	order.DatePurchased = now
	order.LastModified = now
	order.OrdersDateFinished = now
	order.ShippingDuration = now
	d := dbConn.Create(&order)
	if d.Error != nil {
		return nil, d.Error
	}
	productClient := orderRepository.clientConn.GetProductsClient()

	// conn, err := grpc.Dial("localhost:8805", grpc.WithInsecure())
	// if err != nil {
	// 	fmt.Println("Error in Dial(): gRPC connection : ", err)
	// }

	//productClient := pb.NewProductsClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	resp, err := productClient.UpdateProductStock(
		ctx,
		&pb.UpdateProductStockRequest{
			AddStock:    0,
			RemoveStock: 1,
			ProductID:   1,
		})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
	return &order, nil
}
