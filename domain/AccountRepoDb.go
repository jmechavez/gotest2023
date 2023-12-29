package domain

import (
	"strconv"

	"github.com/jmechavez/gotest2023/errorCust"
	"github.com/jmechavez/gotest2023/logger"
	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDB struct {
	client *sqlx.DB
}

func (d AccountRepositoryDB) Save(a Account) (*Account, *errorCust.AppError) {
	sqlInsert := "INSERT INTO accounts (player_id, regopen_date, account_type, pin, status, amount) values (?,?,?,?,?,?)"

	result, err := d.client.Exec(sqlInsert, a.PlayerId, a.RegOpenDate, a.AccountType, a.Pin, a.Status, a.Amount)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errorCust.NewUnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new account: " + err.Error())
		return nil, errorCust.NewUnexpectedError("Unexpected error from database")
	}

	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

// func (d AccountRepositoryDB) Save(a Account) (*Account, *errorCust.AppError) {
// 	sqlInsert := "INSERT INTO accounts (player_id, regopen_date, account_type, pin, status, amount) values (?,?,?,?,?,?)"

// 	result, err := d.client.Exec(sqlInsert, a.PlayerId, a.RegOpenDate, a.AccountType, a.Pin, a.Status, a.Amount)
// 	if err != nil {
// 		logger.Error("Error while creating new account: " + err.Error())
// 		return nil, errorCust.NewUnexpectedError("Unexpected error from database")
// 	}
// 	id, err := result.LastInsertId()
// 	if err != nil {
// 		logger.Error("Error while getting last insert id for new account: " + err.Error())
// 		return nil, errorCust.NewUnexpectedError("Unexpected error from database")
// 	}
// 	a.AccountId = strconv.FormatInt(id, 10)
// 	return &a, nil
// }

func NewAccountRepositoryDB(dbClient *sqlx.DB) AccountRepositoryDB {
	return AccountRepositoryDB{dbClient}
}
