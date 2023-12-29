package domain

import (
	"github.com/jmechavez/gotest2023/dto"
	"github.com/jmechavez/gotest2023/errorCust"
)

type Account struct {
	AccountId   string
	PlayerId    string
	RegOpenDate string
	AccountType string
	Pin         string
	Amount      float64
	Status      string
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountId: a.AccountId}
}

type AccountRepository interface {
	Save(Account) (*Account, *errorCust.AppError)
}
