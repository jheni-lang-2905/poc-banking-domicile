package router

import (
	"github.com/gorilla/mux"
	"github.com/poc-banking-domicile/src/presenters/router/rotas"
)

func RouterConfig() *mux.Router {
	router := mux.NewRouter()
	return rotas.RotasConfig(router)
}
