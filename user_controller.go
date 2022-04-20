package goasample

import (
	"goa-sample/src/database"
	"goa-sample/src/domain/model"
	"context"
	usercontroller "goa-sample/gen/user_controller"
	"github.com/jinzhu/copier"
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
	// ユーザー一覧を検索
	users := []model.User{}
	database.DB.Find(&users)

	// 返還用オブジェクト作成
	result_users := []*usercontroller.User{}
	// 検索結果を返還用オブジェクトへマッピング
	copier.Copy(&result_users, &users)

	return result_users, nil
}

// ユーザ検索
func (s *userControllersrvc) GetUser(ctx context.Context, p *usercontroller.GetUserPayload) (res *usercontroller.User, err error) {
	s.logger.Print("userController.GetUser")

	user := model.User{Id: p.ID}
	// ユーザー検索
	database.DB.Take(&user)

	// 返還用オブジェクト作成
	res = &usercontroller.User{}
	// 検索結果を返還用オブジェクトへマッピング
	copier.Copy(&res, &user)

	return res, nil
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
