package services

import (
	"context"
	"log"

	"gitlab.com/e-market724/back-end/user_service/genproto/user_service"
	"gitlab.com/e-market724/back-end/user_service/storage"
)

type UserService struct {
	storage storage.StorageRepoI
	user_service.UnimplementedUserServiceServer
}

func NewUserService(storage storage.StorageRepoI) *UserService {

	return &UserService{storage: storage}

}

func (u *UserService) CheckExists(ctx context.Context, req *user_service.Common) (*user_service.CommonResp, error) {

	resp, err := u.storage.GetUserRepo().IsExists(ctx, req)

	if err != nil {

		log.Println(err)
		return nil, err
	}

	return resp, nil
}
func (u *UserService) CreateUser(ctx context.Context, req *user_service.UserCreateReq) (*user_service.User, error) {

	resp, err := u.storage.GetUserRepo().CreateUser(ctx, req)

	if err != nil {

		log.Println(err)
		return nil, err
	}

	return resp, nil
}
func (u *UserService) DeleteUser(ctx context.Context, req *user_service.DeleteReq) (*user_service.Empty, error) {

	resp, err := u.storage.GetUserRepo().DeleteUser(ctx, req)

	if err != nil {

		log.Println(err)
		return nil, err
	}

	return resp, nil
}
func (u *UserService) GetUser(ctx context.Context, req *user_service.GetByIdReq) (*user_service.User, error) {

	resp, err := u.storage.GetUserRepo().GetUserById(ctx, req)

	if err != nil {

		log.Println(err)
		return nil, err
	}

	return resp, nil
}
func (u *UserService) GetUsers(ctx context.Context, req *user_service.GetListReq) (*user_service.UserGetListResp, error) {
	
	resp, err := u.storage.GetUserRepo().GetUsers(ctx, req)

	if err != nil {

		log.Println(err)
		return nil, err
	}

	return resp, nil
}
func (u *UserService) UpdateUser(ctx context.Context, req *user_service.UserUpdateReq) (*user_service.User, error) {

	resp, err := u.storage.GetUserRepo().UpdateUser(ctx, req)

	if err != nil {

		log.Println(err)
		return nil, err
	}

	return resp, nil
}
func (u *UserService) UserLogin(ctx context.Context, req *user_service.UserLogIn) (*user_service.Clamis, error) {

	resp, err := u.storage.GetUserRepo().UserLogin(ctx, req)

	if err != nil {

		log.Println(err)
		return nil, err
	}

	return resp, nil
}
