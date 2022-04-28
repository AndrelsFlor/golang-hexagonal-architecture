package service

import (
	"rest_api/domain"
	"rest_api/dto"
	"rest_api/errs"
)

type CustomerService interface {
	GetAllCustomer(status string) (*[]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) (*[]dto.CustomerResponse, *errs.AppError) {
	customers, err := s.repo.FindAll(status)
	var customerDtos []dto.CustomerResponse

	if err != nil {
		return nil, err
	}

	for _, customer := range *customers {
		customerDto := customer.ToDto()
		customerDtos = append(customerDtos, customerDto)
	}

	return &customerDtos, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}

	response := c.ToDto()

	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
