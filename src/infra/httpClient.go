package infra

import "net/http"

func HttpClient() *http.Client {
	return &http.Client{}
}
