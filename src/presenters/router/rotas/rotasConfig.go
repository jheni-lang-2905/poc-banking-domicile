package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Rotas struct {
	URI    string
	Metodo string
	Funcao func(http http.ResponseWriter, r *http.Request)
}

func RotasConfig(r *mux.Router) *mux.Router {
	Rotas := BankingDomicileRotas

	for _, rota := range Rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}
	return r

}
