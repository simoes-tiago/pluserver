package service

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"pluserver/domain"
)

func (s *Service) CreateUser(user domain.User) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Create(&user)
		if result.Error != nil {
			log.Println("error creating user:", result.Error)
			return result.Error
		}

		u := s.GetUser(tx, user.Username)

		account := s.CreateAccount(tx, u.ID)
		if account == nil {
			return errors.New("error creating account")
		}
		u.Account = account
		result = tx.Updates(u)
		if result.Error != nil {
			return result.Error
		}
		return result.Error
	})
}

func (s *Service) DeleteUser(externalTx *gorm.DB, user string) error {
	var tx *gorm.DB
	if externalTx == nil {
		tx = s.db
	} else {
		tx = externalTx
	}

	result := tx.
		Delete(&domain.User{}, "username = ?", &user)

	return result.Error
}

func (s *Service) UpdateUser(externalTx *gorm.DB, username string, user domain.User) error {
	var tx *gorm.DB
	if externalTx == nil {
		tx = s.db
	} else {
		tx = externalTx
	}

	u := s.GetUser(tx, username)
	if u.Username == "" {
		return errors.New("not found")
	}
	user.ID = u.ID
	result := tx.Updates(&user)

	return result.Error
}
func (s *Service) GetUser(externalTx *gorm.DB, user string) domain.User {
	var tx *gorm.DB
	if externalTx == nil {
		tx = s.db
	} else {
		tx = externalTx
	}

	var result domain.User
	tx.
		Where("username = ?", &user).
		Preload("Account").
		First(&result)

	return result
}

func (s *Service) GetAllUser() []domain.User {
	var result []domain.User
	s.db.Find(&result)

	return result
}
