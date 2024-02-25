package service

import (
	"errors"
	"gorm.io/gorm"
	"pluserver/domain"
)

func (s *Service) CreateTransaction(transaction domain.Transaction) error {
	return s.db.Transaction(func(tx *gorm.DB) error {

		var destAccount, origAccount *domain.Account
		if transaction.DestinationAccountID == nil {
			return errors.New("missing destination-account-id")
		}

		destAccount = s.GetAccount(tx, *transaction.DestinationAccountID)
		if destAccount == nil {
			return errors.New("invalid destination-account-id")
		}
		switch transaction.Type {
		case domain.Withdraw:
			if destAccount.Balance-transaction.Amount >= 0 && transaction.Amount > 0 {
				destAccount.Balance -= transaction.Amount
			} else {
				return errors.New("invalid amount")
			}

		case domain.Deposit:
			if transaction.Amount > 0 {
				destAccount.Balance += transaction.Amount
			} else {
				return errors.New("invalid amount")
			}
		case domain.Transfer:
			if transaction.OriginAccountID == nil {
				return errors.New("missing destination-account-id")
			}
			origAccount = s.GetAccount(tx, *transaction.DestinationAccountID)
			if origAccount == nil {
				return errors.New("invalid origin-account-id")
			}
			if origAccount.Balance-transaction.Amount >= 0 && transaction.Amount > 0 {
				origAccount.Balance -= transaction.Amount
				destAccount.Balance += transaction.Amount
				err := s.UpdateAccount(tx, origAccount)
				if err != nil {
					return err
				}
			} else {
				return errors.New("invalid amount")
			}
		}
		result := tx.Create(&transaction)
		if result.Error != nil {
			return result.Error
		}
		err := s.UpdateAccount(tx, destAccount)
		if err != nil {
			return err
		}

		return result.Error
	})
}

func (s *Service) DeleteTransaction(id uint) error {
	result := s.db.
		Delete(&domain.Transaction{}, id)

	return result.Error
}

func (s *Service) UpdateTransaction(id uint, transaction domain.Transaction) error {
	t := s.GetTransaction(id)
	if t.ID == 0 {
		return errors.New("not found")
	}
	transaction.ID = t.ID
	result := s.db.Updates(&transaction)

	return result.Error
}
func (s *Service) GetTransaction(id uint) domain.Transaction {
	var result domain.Transaction
	s.db.
		Preload("Account").
		First(&result, id)

	return result
}

func (s *Service) GetAllTransactions() []domain.Transaction {
	var result []domain.Transaction
	s.db.Find(&result)

	return result
}
