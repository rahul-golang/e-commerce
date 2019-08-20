package repository

import (
	"context"
	"fmt"
	"gokit/ecommerse/users/pkg/database"
	"gokit/ecommerse/users/pkg/models"

	"time"
)

//UserRepositoryInterface implimets all methods in UserRepository
type UserRepositoryInterface interface {
	Create(ctx context.Context, createReq models.User) (createResp *models.User, err error)
	Get(ctx context.Context, id string) (getResp *models.User, err error)
	Delete(ctx context.Context, id string) (deleteResp *models.DeleteUserResp, err error)
	Update(ctx context.Context, updateReq models.User) (updateResp *models.User, err error)
	All(ctx context.Context) (getAllResp []*models.User, err error)
}

// UserRepository **
type UserRepository struct {
	mysqlInterface database.MysqlInterface
}

//NewUserRepository inject dependancies of DataStore
func NewUserRepository(mysqlInterface database.MysqlInterface) UserRepositoryInterface {
	return &UserRepository{mysqlInterface: mysqlInterface}
}

//Create add new record in datastore
func (userRepository *UserRepository) Create(ctx context.Context, user models.User) (*models.User, error) {

	dbConn := userRepository.mysqlInterface.NewClientConnection()

	dbConn.AutoMigrate(&models.User{})
	createOn := time.Now().In(time.UTC)

	// record create Time
	user.CreatedOn = createOn

	d := dbConn.Create(&user)
	if d.Error != nil {
		return nil, d.Error
	}

	return &user, nil
}

/**

 */
func (userRepository *UserRepository) Get(ctx context.Context, id string) (*models.User, error) {

	dbConn := userRepository.mysqlInterface.NewClientConnection()
	fmt.Println("id", id)
	user := models.User{}
	err := dbConn.Where("id=?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (userRepository *UserRepository) Delete(ctx context.Context, id string) (*models.DeleteUserResp, error) {
	dbConn := userRepository.mysqlInterface.NewClientConnection()

	err := dbConn.Where("id=?", id).Delete(&models.User{}).Error
	if err != nil {
		return nil, err
	}
	return &models.DeleteUserResp{
		Message: "Records Deleted Sucessfully",
		ID:      id,
	}, nil
}
func (userRepository *UserRepository) Update(ctx context.Context, user models.User) (*models.User, error) {

	dbConn := userRepository.mysqlInterface.NewClientConnection()
	err := dbConn.Model(&models.User{}).Where("id=?", user.ID).Update(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (userRepository *UserRepository) All(ctx context.Context) (getAllResp []*models.User, err error) {

	dbConn := userRepository.mysqlInterface.NewClientConnection()

	//users := make([]*models.User, 0)
	fmt.Println("")

	users := []*models.User{}
	err = dbConn.Find(&users).Error
	fmt.Println("users : ", users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
