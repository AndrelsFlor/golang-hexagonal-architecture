//esse arquivo é uma "porta". Porta define o protocolo que precisa ser seguido por qqr trecho de código
//que queira se conectar com "customer" através da interface CustomerRepository
//OBS: em go vc não precisa explicitamente implementar uma interface. Só de implementar todos os métodos
// descritos na mesma, vc já a está implementando

package domain

import (
	"rest_api/dto"
	"rest_api/errs"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateofBirth string `db:"date_of_birth"`
	Status      string
}

func (c Customer) statusAsText() string {
	statusAstext := "active"

	if c.Status == "0" {
		statusAstext = "inactive"
	}

	return statusAstext
}

func (c Customer) ToDto() dto.CustomerResponse {

	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateofBirth: c.DateofBirth,
		Status:      c.statusAsText(),
	}
}

type CustomerRepository interface {
	FindAll(status string) (*[]Customer, *errs.AppError)
	// ById Usa ponteiro pq queremos enviar nil caso nao haja customer
	ById(string) (*Customer, *errs.AppError)
}
