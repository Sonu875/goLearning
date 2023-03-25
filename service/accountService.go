package service

import (
	"time"

	appError "github.com/Sonu875/goLearning/Errors"
	"github.com/Sonu875/goLearning/domain"
	"github.com/Sonu875/goLearning/dto"
)

type AccountService interface {
	NewAccount(dto.AccountRequest) (*dto.AccountResponse, *appError.AppError)
	NewTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *appError.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepo
}

func (s DefaultAccountService) NewAccount(req dto.AccountRequest) (*dto.AccountResponse, *appError.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	reqDomain := domain.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}
	response, err := s.repo.Save(reqDomain)
	if err != nil {
		return nil, err
	}
	reponseDto := dto.AccountResponse{
		AccountId: response.AccountId,
	}
	return &reponseDto, nil
}

func (s DefaultAccountService) NewTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *appError.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	account, err := s.repo.FindAccountById(req.AccountId)
	if err != nil {
		return nil, err
	}
	var balance float64

	if req.TransactionType == "withdrawal" {
		if account.Amount < req.Amount {
			err := appError.NewValidation("Amount is bigger than balance transaction failed")
			return nil, err
		} else {
			balance = account.Amount - req.Amount
		}

	} else {
		balance = account.Amount + req.Amount
	}

	transDomain := domain.Transaction{
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
	}

	txs, err := s.repo.MakeTransaction(transDomain, balance)
	if err != nil {
		return nil, err
	}

	response := txs.ResponseDto(balance)
	return &response, nil
}

func NewAccountService(repo domain.AccountRepo) DefaultAccountService {
	return DefaultAccountService{repo: repo}
}
