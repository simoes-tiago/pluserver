package service

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	mockDB "pluserver/db/mock"
	"pluserver/domain"
	"testing"
)

func TestService_CreateTransaction_OK(t *testing.T) {
	mockDb, mock := mockDB.NewMockDB()

	service := NewService(mockDb)
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT ").
		WillReturnRows(sqlmock.NewRows([]string{"id", "balance"}).
			AddRow(1, 10))
	mock.ExpectQuery("SELECT ").
		WillReturnRows(sqlmock.NewRows([]string{"id", "balance"}).
			AddRow(2, 10))
	mock.ExpectExec("UPDATE ").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectQuery("INSERT ").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
			AddRow(1))
	mock.ExpectExec("UPDATE ").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := service.CreateTransaction(
		domain.Transaction{
			OriginAccountID:      &[]int64{1}[0],
			DestinationAccountID: &[]int64{2}[0],
			Amount:               5,
			Type:                 domain.Transfer,
		})

	assert.Nil(t, mock.ExpectationsWereMet())
	assert.Nil(t, err)
}

func TestService_CreateTransaction_NOK(t *testing.T) {
	mockDb, mock := mockDB.NewMockDB()

	service := NewService(mockDb)

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT ").
		WillReturnRows(sqlmock.NewRows([]string{"id", "balance"}).
			AddRow(1, 10))
	mock.ExpectQuery("SELECT ").
		WillReturnRows(sqlmock.NewRows([]string{"id", "balance"}).
			AddRow(2, 10))
	mock.ExpectExec("UPDATE ").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectQuery("INSERT ").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
			AddRow(1))
	mock.ExpectExec("UPDATE ").
		WillReturnError(sqlmock.ErrCancelled)
	mock.ExpectRollback()

	err := service.CreateTransaction(
		domain.Transaction{
			OriginAccountID:      &[]int64{1}[0],
			DestinationAccountID: &[]int64{2}[0],
			Amount:               5,
			Type:                 domain.Transfer,
		})
	assert.Nil(t, mock.ExpectationsWereMet())
	assert.NotNil(t, err)
}

func TestService_CreateWithdraw_OK(t *testing.T) {
	mockDb, mock := mockDB.NewMockDB()

	service := NewService(mockDb)
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT ").
		WillReturnRows(sqlmock.NewRows([]string{"id", "balance"}).
			AddRow(1, 10))
	mock.ExpectQuery("INSERT ").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
			AddRow(1))
	mock.ExpectExec("UPDATE ").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := service.CreateTransaction(
		domain.Transaction{
			DestinationAccountID: &[]int64{2}[0],
			Amount:               5,
			Type:                 domain.Withdraw,
		})

	assert.Nil(t, mock.ExpectationsWereMet())
	assert.Nil(t, err)
}

func TestService_CreateWithdraw_NOK(t *testing.T) {
	mockDb, mock := mockDB.NewMockDB()

	service := NewService(mockDb)

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT ").
		WillReturnRows(sqlmock.NewRows([]string{"id", "balance"}).
			AddRow(1, 10))

	err := service.CreateTransaction(
		domain.Transaction{
			OriginAccountID:      &[]int64{1}[0],
			DestinationAccountID: &[]int64{2}[0],
			Amount:               15,
			Type:                 domain.Withdraw,
		})

	assert.Nil(t, mock.ExpectationsWereMet())
	assert.NotNil(t, err)
}

func TestService_CreateDeposit_OK(t *testing.T) {
	mockDb, mock := mockDB.NewMockDB()

	service := NewService(mockDb)

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT ").
		WillReturnRows(sqlmock.NewRows([]string{"id", "balance"}).
			AddRow(1, 10))
	mock.ExpectQuery("INSERT ").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
			AddRow(1))
	mock.ExpectExec("UPDATE ").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := service.CreateTransaction(
		domain.Transaction{
			OriginAccountID:      &[]int64{1}[0],
			DestinationAccountID: &[]int64{2}[0],
			Amount:               5,
			Type:                 domain.Deposit,
		})

	assert.Nil(t, mock.ExpectationsWereMet())
	assert.Nil(t, err)
}

func TestService_CreateDeposit_NOK(t *testing.T) {
	mockDb, mock := mockDB.NewMockDB()

	service := NewService(mockDb)

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT ").
		WillReturnRows(sqlmock.NewRows([]string{"id", "balance"}).
			AddRow(1, 10))
	mock.ExpectQuery("INSERT ").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
			AddRow(1))
	mock.ExpectExec("UPDATE ").
		WillReturnError(sqlmock.ErrCancelled)
	mock.ExpectRollback()

	err := service.CreateTransaction(
		domain.Transaction{
			OriginAccountID:      &[]int64{1}[0],
			DestinationAccountID: &[]int64{2}[0],
			Amount:               5,
			Type:                 domain.Deposit,
		})

	assert.Nil(t, mock.ExpectationsWereMet())
	assert.NotNil(t, err)
}
