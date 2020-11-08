package mysql_utils

import (
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/ninoude/bookstore_users-api/utils/errors"
)

const (
	errorsNoRows = "no rows in result set"
)

func PaseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorsNoRows) {
			return errors.NewNotFoundError("no record matting given id")
		}
		return errors.NewInternalServerError("error paseing database response")
	}
	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("invalid data")
	}
	return errors.NewInternalServerError("error processing request")
}
