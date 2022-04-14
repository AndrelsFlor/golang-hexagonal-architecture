package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"rest_api/domain"
	"rest_api/service"
)

func Start() {
	// mux é uma lib que simplifica o match das rotas
	// Todas as funções dela são extamente iguais às do modulo de http nativo
	router := mux.NewRouter()
	//wiring
	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	//define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	//starts server
	err := http.ListenAndServe("localhost:8000", router)

	if err != nil {
		log.Fatal(err)
	}
}
