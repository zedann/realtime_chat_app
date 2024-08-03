package user

import (
	"context"
	"strconv"
	"time"

	"github.com/zedann/realtime_chat_app/server/util"
)

type service struct {
	Repository
	timeout time.Duration
}

func NewService(repo Repository) Service {

	return &service{
		repo,
		time.Duration(2) * time.Second,
	}
}

func (s *service) CreateUser(ctx context.Context, req *CreateUserReq) (*CreateUserRes, error) {
	c, cancel := context.WithTimeout(ctx, s.timeout)

	defer cancel()

	//TODO: hash password
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	u := &User{
		Username: req.Username,
		Password: hashedPassword,
		Email:    req.Email,
	}

	r, err := s.Repository.CreateUser(c, u)

	if err != nil {
		return nil, err
	}

	return &CreateUserRes{
		ID:       strconv.Itoa(int(r.ID)),
		Username: r.Username,
		Email:    r.Email,
	}, nil

}
