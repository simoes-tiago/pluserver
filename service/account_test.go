package service

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	mockDB "pluserver/db/mock"
	"pluserver/domain"
	"testing"
)

func TestService_CreateAccount_OK(t *testing.T) {
	mockDb, mock := mockDB.NewMockDB()

	service := NewService(mockDb)

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT ").
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id"}).
			AddRow(1, 1))
	mock.ExpectCommit()
	mock.ExpectQuery("SELECT ").
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id"}).
			AddRow(1, 1))

	account := service.CreateAccount(nil, 1)

	assert.Equal(t, *account, domain.Account{
		ID:     1,
		UserID: 1,
	})
	assert.Nil(t, mock.ExpectationsWereMet())
	assert.NotNil(t, account)
}

func TestService_CreateAccount_NOK(t *testing.T) {
	mockDb, mock := mockDB.NewMockDB()

	service := NewService(mockDb)

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT ").
		WillReturnError(sqlmock.ErrCancelled)

	account := service.CreateAccount(nil, 1)
	assert.Nil(t, mock.ExpectationsWereMet())
	assert.Nil(t, account)
}

func TestService_DeleteAccount_OK(t *testing.T) {
	mockDb, mock := mockDB.NewMockDB()

	service := NewService(mockDb)

	mock.ExpectBegin()
	mock.ExpectExec("DELETE ").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := service.DeleteAccount(1)

	assert.Nil(t, mock.ExpectationsWereMet())
	assert.Nil(t, err)
}

func TestService_DeleteAccount_NOK(t *testing.T) {
	mockDb, mock := mockDB.NewMockDB()

	service := NewService(mockDb)

	mock.ExpectBegin()
	mock.ExpectExec("DELETE ").
		WillReturnError(sqlmock.ErrCancelled)
	err := service.DeleteAccount(1)

	assert.Nil(t, mock.ExpectationsWereMet())
	assert.NotNil(t, err)
}
func TestService_GetAccount_OK(t *testing.T) {
	mockDb, mock := mockDB.NewMockDB()

	service := NewService(mockDb)

	mock.ExpectQuery("SELECT ").
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id"}).
			AddRow(1, 1))

	account := service.GetAccount(nil, 1)

	assert.Nil(t, mock.ExpectationsWereMet())
	assert.Equal(t, *account, domain.Account{
		ID:     1,
		UserID: 1,
	})
	assert.NotNil(t, account)
}

func TestService_GetAccount_NOK(t *testing.T) {
	mockDb, mock := mockDB.NewMockDB()

	service := NewService(mockDb)

	mock.ExpectQuery("SELECT ").
		WillReturnError(sqlmock.ErrCancelled)

	account := service.GetAccount(nil, 1)

	assert.Nil(t, mock.ExpectationsWereMet())
	assert.Nil(t, account)
}

func TestService_UpdateAccount_OK(t *testing.T) {
	mockDb, mock := mockDB.NewMockDB()

	service := NewService(mockDb)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE ").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := service.UpdateAccount(nil,
		&domain.Account{
			ID:     1,
			UserID: 1,
		})

	assert.Nil(t, mock.ExpectationsWereMet())
	assert.Nil(t, err)
}

func TestService_UpdateAccount_NOK(t *testing.T) {
	mockDb, mock := mockDB.NewMockDB()

	service := NewService(mockDb)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE ").
		WillReturnError(sqlmock.ErrCancelled)

	err := service.UpdateAccount(nil,
		&domain.Account{
			ID:     1,
			UserID: 1,
		})

	assert.Nil(t, mock.ExpectationsWereMet())
	assert.NotNil(t, err)
}
