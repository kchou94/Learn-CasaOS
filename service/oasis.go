package service

type OasisService interface{}

type oasisService struct {
}

func NewOasisService() OasisService {
	return &oasisService{}
}
