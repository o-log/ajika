package gormhelper

import (
	"errors"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

const MysqlErRowIsReferended2 = 1451 // "Cannot delete or update a parent row: a foreign key constraint fails": https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html

// This function is needed because GORM doesn't translate mysql 1451 error ("Cannot delete or update a parent row: a foreign key constraint fails") to "ErrForeignKeyViolated"
// (https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html)
// only 1452 error ("Cannot add or update a child row: a foreign key constraint fails") is translated:
// https://github.com/go-gorm/mysql/blob/b87f024d0e3d00c24b87e8e2b3d2f19a9f1eea31/error_translator.go#L12
// So this function checks both: ErrForeignKeyViolated detected by GORM and also mysql 1451 error
func IsErrForeignKeyViolated(err error) bool {
	if errors.Is(err, gorm.ErrForeignKeyViolated) {
		return true
	}

	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) { // supports wrapped errors (unwraps and checks all): https://go.dev/blog/go1.13-errors
		if mysqlErr.Number == MysqlErRowIsReferended2 {
			return true
		}
	}

	return false
}
