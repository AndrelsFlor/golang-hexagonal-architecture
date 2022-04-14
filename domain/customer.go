//esse arquivo é uma "porta". Porta define o protocolo que precisa ser seguido por qqr trecho de código
//que queira se conectar com "customer" através da interface CustomerRepository
//OBS: em go vc não precisa explicitamente implementar uma interface. Só de implementar todos os métodos
// descritos na mesma, vc já a está implementando

package domain

import "rest_api/errs"

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateofBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	// ById Usa ponteiro pq queremos enviar nil caso nao haja customer
	ById(string) (*Customer, *errs.AppError)
}
