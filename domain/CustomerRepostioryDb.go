package domain

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"rest_api/errs"
	"rest_api/logger"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {

	return CustomerRepositoryDb{dbClient}
}

func statusToInt(status string) int {
	var statusMap = make(map[string]int)
	statusMap["inactive"] = 0
	statusMap["active"] = 1

	return statusMap[status]
}

func (d CustomerRepositoryDb) FindAll(status string) (*[]Customer, *errs.AppError) {
	var findAllSql string
	var err error
	customers := make([]Customer, 0)

	if len(status) > 0 {
		findAllSql = "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		statusInt := statusToInt(status)
		err = d.client.Select(&customers, findAllSql, statusInt)
	} else {
		findAllSql = "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = d.client.Select(&customers, findAllSql)

	}
	if err != nil {
		logger.Error("Errror while querying customers table" + err.Error())
		return nil, errs.NewUnexpectedError("Error querying customers table")
	}

	return &customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"
	var c Customer
	err := d.client.Get(&c, customerSql, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected db error")
		}
	}

	return &c, nil
}
