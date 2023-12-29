package dto

import (
	"strings"

	"github.com/jmechavez/gotest2023/errorCust"
)

type NewAccountRequest struct {
	PlayerId    string  `json:"player_id"`
	AccountType string  `json:"account_type"`
	Pin         string  `json:"pin"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *errorCust.AppError {
	if r.Amount < 5000 {
		return errorCust.NewValidationError("To open new player account you need to depostit atleast php5,000.00 ")
	}
	if strings.ToLower(r.AccountType) != "saving" && r.AccountType != "checking" {
		return errorCust.NewValidationError("Account type should be checking or saving")
	}
	return nil
}
