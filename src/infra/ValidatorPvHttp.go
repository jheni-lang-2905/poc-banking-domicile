package infra

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	domain "github.com/poc-banking-domicile/src/domain/entities"
	"github.com/poc-banking-domicile/src/infra/config"
)

type ValidatorsResponse struct {
	Success bool
	Message string
}

func ValidatorPvHttp(bankAddress domain.BankingDomicileEntities) (ValidatorsResponse, error) {
	bank := bankAddress.BankUuid(bankAddress)

	fmt.Printf("bank", bank)
	url := config.UrlValidator
	var validatorResponse ValidatorsResponse

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Erro ao criar requisição:", err)
		return validatorResponse, err
	}

	req.Header.Set("Cache-control", "no-cache")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("host", config.Host)
	req.Header.Set("x-api-key", config.XApiKey)
	req.Header.Set("x-api-gtw-id", config.XApiGtwId)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao fazer requisição:", err)
		return validatorResponse, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler resposta:", err)
		return validatorResponse, err
	}

	if resp.StatusCode == http.StatusOK {
		if err := json.Unmarshal(body, &validatorResponse); err != nil {
			fmt.Println("Erro ao decodificar JSON:", err)
			return validatorResponse, nil
		}
	}
	return validatorResponse, nil
}
