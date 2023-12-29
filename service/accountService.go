package service

import (
	"time"

	"github.com/jmechavez/gotest2023/domain"
	"github.com/jmechavez/gotest2023/dto"
	"github.com/jmechavez/gotest2023/errorCust"
)

type AccountService interface {
	NewAccount(request dto.NewAccountRequest) (*dto.NewAccountResponse, *errorCust.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

// NewAccount creates a new account based on the provided request.
// It uses the DefaultAccountService instance and interacts with the underlying repository.
func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errorCust.AppError) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}
	// Create a new Account instance with the provided request data.
	a := domain.Account{
		PlayerId:    req.PlayerId,
		RegOpenDate: time.Now().Format("2006-01-02 04:05"),
		AccountType: req.AccountType,
		Pin:         req.Pin,
		Amount:      req.Amount,
	}

	// Save the new account using the repository.
	newAccount, err := s.repo.Save(a)
	if err != nil {
		// Return an error if there's an issue during the account creation.
		return nil, err
	}

	// Convert the newly created account to a response DTO.
	response := newAccount.ToNewAccountResponseDto()

	// Return the response DTO and nil error, indicating successful account creation.
	return &response, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}
