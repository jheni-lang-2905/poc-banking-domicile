package rotas

import (
	"net/http"

	"github.com/poc-banking-domicile/src/presenters/controllers"
)

var BankingDomicileRotas = Rotas[]{
	{
		URI:    "/validator/{pvNumber}",
		Metodo: http.MethodGet,
		Funcao: controllers.ValidateToken,
	},
}