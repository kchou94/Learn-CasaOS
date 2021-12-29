package service

type SystemService interface{}

type systemService struct{}

func NewSystemService() SystemService {
	return &systemService{}
}
