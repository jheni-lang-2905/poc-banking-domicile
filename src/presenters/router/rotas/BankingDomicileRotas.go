package rotas

import (
	"net/http"

	usecase "github.com/poc-banking-domicile/src/useCase"
)

var BankingDomicileRotas = []Rotas{
	{
		URI:    "/validator/{pvNumber}",
		Metodo: http.MethodGet,
		Funcao: usecase.ValidateTokenUseCase,
	},
}
