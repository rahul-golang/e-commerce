package repository

import (
	"context"
	"fmt"
	"gokit/ecommerse/products/pkg/database"
	"gokit/ecommerse/products/pkg/models"

	"time"

	"github.com/jinzhu/gorm"
)

//ProductRepositoryInterface implimets all methods in ProductRepository
type ProductRepositoryInterface interface {
	Create(ctx context.Context, createReq models.Product) (createResp *models.Product, err error)
	Get(ctx context.Context, id string) (getResp *models.Product, err error)
	Delete(ctx context.Context, id string) (deleteResp *models.DeleteProductResp, err error)
	Update(ctx context.Context, updateReq models.Product) (updateResp *models.Product, err error)
	All(ctx context.Context) (getAllResp []*models.Product, err error)
	UpdateProductStock(context.Context, models.UpdateStockReq) (*models.UpdateStockResp, error)
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
	defer dbConn.Close()

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
	defer dbConn.Close()
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
	defer dbConn.Close()

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
	defer dbConn.Close()
	fmt.Println(product)
	err := dbConn.Model(&models.Product{}).Where("id=?", product.ID).Update(&product).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (productRepository *ProductRepository) All(ctx context.Context) (getAllResp []*models.Product, err error) {

	dbConn := productRepository.mysqlInterface.NewClientConnection()
	defer dbConn.Close()

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

//UpdateProductStock UPDATES Products Stock
func (productRepository *ProductRepository) UpdateProductStock(ctx context.Context, updateStockReq models.UpdateStockReq) (*models.UpdateStockResp, error) {

	products := models.Product{}
	products.ID = uint(updateStockReq.ProductID)
	dbConn := productRepository.mysqlInterface.NewClientConnection()
	defer dbConn.Close()

	if updateStockReq.RemoveStock != 0 {
		err := dbConn.Model(&products).Where("!products_quantity < ?", updateStockReq.RemoveStock).Update("products_quantity", gorm.Expr("products_quantity - ?", updateStockReq.RemoveStock)).Error
		if err != nil {
			return nil, err
		}
	} else {

		err := dbConn.Model(&products).Where("id = ?", updateStockReq.ProductID).Update("products_quantity", gorm.Expr("products_quantity + ?", updateStockReq.AddStock)).Error
		if err != nil {
			return nil, err
		}

	}
	return &models.UpdateStockResp{
		Message: "Updated Successfully",
	}, nil
}
