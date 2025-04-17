package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"gitlab.com/e-market724/back-end/user_service/genproto/user_service"
	"gitlab.com/e-market724/back-end/user_service/storage/repoi"
)

type userRepo struct {
	db *pgx.Conn
}

func NewUserRepo(db *pgx.Conn) repoi.UserRepoI {

	return &userRepo{db: db}
}

func (u *userRepo) CreateUser(ctx context.Context, req *user_service.UserCreateReq) (*user_service.User, error) {


	return nil, nil
}
func (u *userRepo) GetUserById(ctx context.Context, req *user_service.GetByIdReq) (*user_service.User, error) {

	return nil, nil
}
func (u *userRepo) GetUsers(ctx context.Context, req *user_service.GetListReq) (*user_service.UserGetListResp, error) {

	return nil, nil
}
func (u *userRepo) UpdateUser(ctx context.Context, req *user_service.UserUpdateReq) (*user_service.User, error) {

	return nil, nil
}
func (u *userRepo) DeleteUser(ctx context.Context, req *user_service.DeleteReq) (*user_service.Empty, error) {

	return nil, nil
}
func (u *userRepo) IsExists(ctx context.Context, req *user_service.Common) (*user_service.CommonResp, error) {

	return nil, nil
}
func (u *userRepo) UserLogin(ctx context.Context, req *user_service.UserLogIn) (*user_service.Clamis, error) {

	return nil, nil
}
