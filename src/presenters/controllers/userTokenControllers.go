package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/poc-banking-domicile/src/infra/config"
	"github.com/poc-banking-domicile/src/presenters/router"
)

func ValidateToken() {
	config.VariablesLoading()

	r := router.RouterConfig()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.ApiPort), r))
}
