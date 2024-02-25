package service

import (
	"gorm.io/gorm"
	"log"
	"pluserver/domain"
)

func (s *Service) CreateAccount(externalTx *gorm.DB, user int64) *domain.Account {
	var tx *gorm.DB
	if externalTx == nil {
		tx = s.db
	} else {
		tx = externalTx
	}

	result := tx.Create(&domain.Account{
		UserID: user,
	})
	if result.Error != nil {
		log.Printf("error creating account %v", result.Error)
		return nil
	}
	return s.GetAccount(tx, user)
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

func (s *Service) GetAccount(externalTx *gorm.DB, id int64) *domain.Account {
	var tx *gorm.DB
	if externalTx == nil {
		tx = s.db
	} else {
		tx = externalTx
	}

	var result domain.Account
	tx.
		Where("user_id = ?", &id).
		First(&result, id)

	if result.ID == 0 {
		return nil
	}
	return &result
}

func (s *Service) UpdateAccount(externalTx *gorm.DB, account *domain.Account) error {
	var tx *gorm.DB
	if externalTx == nil {
		tx = s.db
	} else {
		tx = externalTx
	}

	result := tx.Updates(account)

	return result.Error
}
