package controllers

import (
	"net/http"

	"github.com/poc-banking-domicile/src/infra/config"
)

func ValidateToken(w http.ResponseWriter, r *http.Request) {
	config.VariablesLoading()

}
