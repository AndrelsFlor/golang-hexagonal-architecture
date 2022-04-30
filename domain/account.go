package domain

import (
	"rest_api/dto"
	"rest_api/errs"
)

type Account struct {
	AccountId   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

type AccountRepository interface {
	Save(*Account) (*Account, *errs.AppError)
}

func (a Account) CanWithdraw(amount float64) bool {
	return a.Amount < amount
}

func (a *Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{a.AccountId}
}
