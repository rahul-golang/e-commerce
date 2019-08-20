package repository

import (
	"context"
	"fmt"
	"gokit/ecommerse/carts/pkg/database"
	"gokit/ecommerse/carts/pkg/models"

	"time"
)

//CartRepositoryInterface implimets all methods in CartRepository
type CartRepositoryInterface interface {
	AddToCart(ctx context.Context, addReq models.Cart) (addResp *models.Cart, err error)
	GetFromAll(ctx context.Context) (allRecordResp []*models.Cart, err error)
	UpdateInCart(ctx context.Context, upadteReq models.Cart) (updateResp *models.Cart, err error)
	DeleteFromCart(ctx context.Context, id string) (deleteResp *models.Cart, err error)
	GetFromCart(ctx context.Context, id string) (createResp *models.Cart, err error)
}

// CartRepository **
type CartRepository struct {
	mysqlInterface database.MysqlInterface
}

//NewCartRepository inject dependancies of DataStore
func NewCartRepository(mysqlInterface database.MysqlInterface) CartRepositoryInterface {
	return &CartRepository{mysqlInterface: mysqlInterface}
}

//AddToCart add new record in datastore
func (cartRepository *CartRepository) AddToCart(ctx context.Context, cart models.Cart) (*models.Cart, error) {

	dbConn := cartRepository.mysqlInterface.NewClientConnection()

	dbConn.AutoMigrate(&models.Cart{})
	createOn := time.Now().In(time.UTC)

	// record create Time
	cart.CreatedAt = createOn

	d := dbConn.Create(&cart)
	if d.Error != nil {
		return nil, d.Error
	}

	return &cart, nil
}

/**

 */
func (cartRepository *CartRepository) GetFromCart(ctx context.Context, id string) (*models.Cart, error) {

	dbConn := cartRepository.mysqlInterface.NewClientConnection()
	fmt.Println("id", id)
	cart := models.Cart{}
	err := dbConn.Where("id=?", id).First(&cart).Error
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

func (cartRepository *CartRepository) DeleteFromCart(ctx context.Context, id string) (*models.Cart, error) {
	dbConn := cartRepository.mysqlInterface.NewClientConnection()

	err := dbConn.Where("id=?", id).Delete(&models.Cart{}).Error
	if err != nil {
		return nil, err
	}
	return &models.Cart{}, nil
}
func (cartRepository *CartRepository) UpdateInCart(ctx context.Context, cart models.Cart) (*models.Cart, error) {

	dbConn := cartRepository.mysqlInterface.NewClientConnection()
	fmt.Println(cart)
	err := dbConn.Model(&models.Cart{}).Where("id=?", cart.ID).Update(&cart).Error
	if err != nil {
		return nil, err
	}

	return &cart, nil
}

func (cartRepository *CartRepository) GetFromAll(ctx context.Context) (getAllResp []*models.Cart, err error) {

	dbConn := cartRepository.mysqlInterface.NewClientConnection()

	//carts := make([]*models.Cart, 0)
	fmt.Println("")

	carts := []*models.Cart{}
	err = dbConn.Find(&carts).Error
	fmt.Println("carts : ", carts)
	if err != nil {
		return nil, err
	}
	return carts, nil
}
