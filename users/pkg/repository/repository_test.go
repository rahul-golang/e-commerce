package repository

import (
	"context"
	"gokit/ecommerse/users/pkg/database"
	"gokit/ecommerse/users/pkg/models"
	"reflect"
	"testing"
)

func TestNewUserRepository(t *testing.T) {
	type args struct {
		mysqlInterface database.MysqlInterface
	}
	tests := []struct {
		name string
		args args
		want UserRepositoryInterface
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserRepository(tt.args.mysqlInterface); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_Create(t *testing.T) {
	type args struct {
		ctx  context.Context
		user models.User
	}
	tests := []struct {
		name           string
		userRepository *UserRepository
		args           args
		want           *models.User
		wantErr        bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.userRepository.Create(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepository.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_Get(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name           string
		userRepository *UserRepository
		args           args
		want           *models.User
		wantErr        bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.userRepository.Get(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepository.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_Delete(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name           string
		userRepository *UserRepository
		args           args
		want           *models.DeleteUserResp
		wantErr        bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.userRepository.Delete(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepository.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_Update(t *testing.T) {
	type args struct {
		ctx  context.Context
		user models.User
	}
	tests := []struct {
		name           string
		userRepository *UserRepository
		args           args
		want           *models.User
		wantErr        bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.userRepository.Update(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepository.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_All(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name           string
		userRepository *UserRepository
		args           args
		wantGetAllResp []*models.User
		wantErr        bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotGetAllResp, err := tt.userRepository.All(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.All() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotGetAllResp, tt.wantGetAllResp) {
				t.Errorf("UserRepository.All() = %v, want %v", gotGetAllResp, tt.wantGetAllResp)
			}
		})
	}
}
