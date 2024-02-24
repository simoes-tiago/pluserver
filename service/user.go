package service

import (
	"errors"
	"log"
	"pluserver/domain"
)

func (s *Service) CreateUser(user domain.User) error {
	result := s.db.Create(&user)
	if result.Error != nil {
		log.Println("error creating user:", result.Error)
		return result.Error
	}

	u := s.GetUser(user.Username)

	account := s.CreateAccount(u.ID)
	if account == nil {
		log.Println("error creating account, deleting user:")
		s.DeleteUser(u.Username)
		return errors.New("error creating account")
	}
	u.Account = account
	result = s.db.Updates(u)
	if result.Error != nil {
		log.Println("error updating user:", result.Error)
		s.DeleteUser(u.Username)
		return result.Error
	}
	return result.Error
}

func (s *Service) DeleteUser(user string) error {
	result := s.db.
		Delete(&domain.User{}, "username = ?", &user)

	return result.Error
}

func (s *Service) UpdateUser(username string, user domain.User) error {
	u := s.GetUser(username)
	if u.Username == "" {
		return errors.New("not found")
	}
	user.ID = u.ID
	result := s.db.Updates(&user)

	return result.Error
}
func (s *Service) GetUser(user string) domain.User {
	var result domain.User
	s.db.
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
