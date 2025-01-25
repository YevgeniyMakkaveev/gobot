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
		return nil, errors.New("wrong")
	}
	return &allProducts[num-1], nil
}

func (s *Service) Delete(num int) (bool, error) {
	if num < 0 || num > len(allProducts)-1 {
		return false, errors.New("wrong_index")
	}
	allProducts = (append(allProducts[:num], allProducts[num+1:]...))
	return true, nil
}

func (s *Service) AddElement(product Product) (bool, error) {
	allProducts = append(allProducts, product)
	return true, nil
}
