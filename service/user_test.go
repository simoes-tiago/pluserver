package service

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	mockDB "pluserver/db/mock"
	"pluserver/domain"
	"testing"
)

func TestService_CreateUser_OK(t *testing.T) {
	mockDb, mock := mockDB.NewMockDB()

	service := NewService(mockDb)

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT ").
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id"}).
			AddRow(1, 1))
	mock.ExpectQuery("SELECT ").
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id"}).
			AddRow(1, 1))
	mock.ExpectQuery("INSERT ").
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id"}).
			AddRow(1, 1))

	err := service.CreateUser(domain.User{
		Username: "1",
	})
	assert.Nil(t, mock.ExpectationsWereMet())
	assert.NotNil(t, err)

}

func TestService_CreateUser_NOK(t *testing.T) {
	mockDb, mock := mockDB.NewMockDB()

	service := NewService(mockDb)

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT ").
		WillReturnError(sqlmock.ErrCancelled)

	err := service.CreateUser(domain.User{
		Username: "1",
	})
	assert.Nil(t, mock.ExpectationsWereMet())
	assert.NotNil(t, err)
}
