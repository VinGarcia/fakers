package main

import "time"

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u UserService) ListUsers(adminsOnly bool) ([]User, error) {
	return []User{
		User{
			ID:        1,
			Name:      "FakeName",
			IsAdmin:   true,
			CreatedAt: time.Time{},
		},
	}, nil
}
