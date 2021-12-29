package service

type ZiMaService interface{}

type zima struct {
}

func NewZiMaService() ZiMaService {
	return &zima{}
}
