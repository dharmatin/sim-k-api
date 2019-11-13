package db

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MysqlMock struct {
	mock.Mock
}

var mockedConn MysqlMock

func TestNewMysqlClientWithBadConnectionString(t *testing.T) {
	_, _, err := sqlmock.NewWithDSN("expected connection")
	if err != nil {
		panic("Got an unexpected error.")
	}
	db := Setup("sqlmock", "unexpected connection")
	defer db.DB().Close()
	if assert.NotNil(t, db.DB().Ping()) {
		assert.Equal(t, errors.New("sql: database is closed"), db.DB().Ping())
	}
}

func TestNewMySqlConnection(t *testing.T) {
	_, _, err := sqlmock.NewWithDSN("mock connection")
	if err != nil {
		panic("Got an unexpected error.")
	}

	db := Setup("sqlmock", "mock connection")
	defer db.DB().Close()
	assert.Nil(t, db.DB().Ping())
}
