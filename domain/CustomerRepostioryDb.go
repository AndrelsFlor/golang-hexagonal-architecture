package domain

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"rest_api/errs"
	"time"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:root@tcp(0.0.0.0:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client}
}

func statusToInt(status string) int {
	var statusMap = make(map[string]int)
	statusMap["inactive"] = 0
	statusMap["active"] = 1

	return statusMap[status]
}

func (d CustomerRepositoryDb) FindAll(status string) (*[]Customer, *errs.AppError) {
	var findAllSql string
	var rows *sql.Rows
	var err error

	if len(status) > 0 {
		findAllSql = "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		statusInt := statusToInt(status)
		rows, err = d.client.Query(findAllSql, statusInt)

	} else {
		findAllSql = "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		rows, err = d.client.Query(findAllSql)

	}
	if err != nil {
		return nil, errs.NewUnexpectedError("Error querying customers table")
	}
	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
		if err != nil {
			return nil, errs.NewUnexpectedError("Error  customers customer table ")
		}

		customers = append(customers, c)
	}
	return &customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"
	var c Customer

	row := d.client.QueryRow(customerSql, id)

	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			log.Println("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected db error")
		}
	}

	return &c, nil
}
