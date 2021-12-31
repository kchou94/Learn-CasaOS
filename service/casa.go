package service

type CasaService interface{}

type casaService struct {
}

func NewOasisService() CasaService {
	return &casaService{}
}
