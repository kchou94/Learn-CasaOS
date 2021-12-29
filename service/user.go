package service

type UserService interface{}

type user struct{}

func NewUserService() UserService {
	return &user{}
}
