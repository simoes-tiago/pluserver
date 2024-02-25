package service

import (
	"log"
	"pluserver/domain"
)

func (s *Service) CreateAccount(user int64) *domain.Account {
	result := s.db.Create(&domain.Account{
		UserID: user,
	})
	if result.Error != nil {
		log.Printf("error creating account %v", result.Error)
		return nil
	}
	return s.GetAccount(user)
}

func (s *Service) DeleteAccount(id int64) error {
	result := s.db.
		Delete(&domain.Account{}, id)

	if result.Error != nil {
		log.Printf("error deleting account %v", result.Error)
		return result.Error
	}

	return nil
}

func (s *Service) GetAccount(id int64) *domain.Account {
	var result domain.Account
	s.db.
		Where("user_id = ?", &id).
		First(&result, id)

	return &result
}

func (s *Service) UpdateAccount(account *domain.Account) error {
	result := s.db.Updates(account)

	return result.Error
}
