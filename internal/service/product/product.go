package product

import "errors"

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Product {
	return allProducts
}

func (s *Service) Get(num int) (*Product, error) {
	if num < 1 || num > len(allProducts) {
		return nil, errors.New("Wrong")
	}
	return &allProducts[num-1], nil

}
