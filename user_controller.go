package goasample

import (
	"context"
	usercontroller "goa-sample/gen/user_controller"
	"log"
)

// user_controller service example implementation.
// The example methods log the requests and return zero values.
type userControllersrvc struct {
	logger *log.Logger
}

// NewUserController returns the user_controller service implementation.
func NewUserController(logger *log.Logger) usercontroller.Service {
	return &userControllersrvc{logger}
}

// ユーザ一覧の検索
func (s *userControllersrvc) GetUsers(ctx context.Context) (res []*usercontroller.User, err error) {
	s.logger.Print("userController.GetUsers")
	//TODO: ロジックを実装
	return
}

// ユーザ検索
func (s *userControllersrvc) GetUser(ctx context.Context, p *usercontroller.GetUserPayload) (res *usercontroller.User, err error) {
	res = &usercontroller.User{}
	s.logger.Print("userController.GetUser")
	//TODO: ロジックを実装
	return
}

// ユーザ更新
func (s *userControllersrvc) UpdateUser(ctx context.Context, p *usercontroller.User) (err error) {
	s.logger.Print("userController.UpdateUser")
	//TODO: ロジックを実装
	return
}

// ユーザ登録
func (s *userControllersrvc) CreateUser(ctx context.Context, p *usercontroller.User) (err error) {
	s.logger.Print("userController.CreateUser")
	//TODO: ロジックを実装
	return
}

// ユーザ削除
func (s *userControllersrvc) DeleteUser(ctx context.Context, p *usercontroller.DeleteUserPayload) (err error) {
	s.logger.Print("userController.DeleteUser")
	//TODO: ロジックを実装
	return
}
