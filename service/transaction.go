package service

import "pluserver/domain"

func (s *Service) CreateTransaction(transaction domain.Transaction) {
	s.db.Create(transaction)
}

func (s *Service) DeleteTransaction(transaction domain.Transaction) {
	s.db.Create(transaction)
}
