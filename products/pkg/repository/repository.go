package repository

import (
	"context"
	"fmt"
	"gokit/ecommerse/products/pkg/database"
	"gokit/ecommerse/products/pkg/models"

	"time"
)

//ProductRepositoryInterface implimets all methods in ProductRepository
type ProductRepositoryInterface interface {
	Create(ctx context.Context, createReq models.Product) (createResp *models.Product, err error)
	Get(ctx context.Context, id string) (getResp *models.Product, err error)
	Delete(ctx context.Context, id string) (deleteResp *models.DeleteProductResp, err error)
	Update(ctx context.Context, updateReq models.Product) (updateResp *models.Product, err error)
	All(ctx context.Context) (getAllResp []*models.Product, err error)
}

// ProductRepository **
type ProductRepository struct {
	mysqlInterface database.MysqlInterface
}

//NewProductRepository inject dependancies of DataStore
func NewProductRepository(mysqlInterface database.MysqlInterface) ProductRepositoryInterface {
	return &ProductRepository{mysqlInterface: mysqlInterface}
}

//Create add new record in datastore
func (productRepository *ProductRepository) Create(ctx context.Context, product models.Product) (*models.Product, error) {

	dbConn := productRepository.mysqlInterface.NewClientConnection()

	dbConn.AutoMigrate(&models.Product{})
	createOn := time.Now().In(time.UTC)

	// record create Time
	product.CreatedAt = createOn

	d := dbConn.Create(&product)
	if d.Error != nil {
		return nil, d.Error
	}

	return &product, nil
}

/**

 */
func (productRepository *ProductRepository) Get(ctx context.Context, id string) (*models.Product, error) {

	dbConn := productRepository.mysqlInterface.NewClientConnection()
	fmt.Println("id", id)
	product := models.Product{}
	err := dbConn.Where("id=?", id).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (productRepository *ProductRepository) Delete(ctx context.Context, id string) (*models.DeleteProductResp, error) {
	dbConn := productRepository.mysqlInterface.NewClientConnection()

	err := dbConn.Where("id=?", id).Delete(&models.Product{}).Error
	if err != nil {
		return nil, err
	}
	return &models.DeleteProductResp{
		Message: "Records Deleted Sucessfully",
		ID:      id,
	}, nil
}
func (productRepository *ProductRepository) Update(ctx context.Context, product models.Product) (*models.Product, error) {

	dbConn := productRepository.mysqlInterface.NewClientConnection()
	fmt.Println(product)
	err := dbConn.Model(&models.Product{}).Where("id=?", product.ID).Update(&product).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (productRepository *ProductRepository) All(ctx context.Context) (getAllResp []*models.Product, err error) {

	dbConn := productRepository.mysqlInterface.NewClientConnection()

	//products := make([]*models.Product, 0)
	fmt.Println("")

	products := []*models.Product{}
	err = dbConn.Find(&products).Error
	fmt.Println("products : ", products)
	if err != nil {
		return nil, err
	}
	return products, nil
}
