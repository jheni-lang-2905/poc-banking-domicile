package usecase

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/poc-banking-domicile/src/infra"
)

func ValidateTokenUseCase(w http.ResponseWriter, r *http.Request) {
	// logica do codigo onde sera feita a regras de negocio
	//passar pv na chamada

	params := mux.Vars(r)

	//request
	//ajustar esse params para o tipo esperado
	respValidatorPv, err := infra.ValidatorPvHttp(params)
	if err != nil {
		fmt.Sprintf("Erro na chamada na validação do pv", err)
		return
	}
	fmt.Sprintf("sucess:", respValidatorPv.Success)
	fmt.Sprintf("message", respValidatorPv.Message)
}
