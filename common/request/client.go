package request

import "net/http"

func NewHttpClient(r HttpRequest) http.Client {
	return http.Client{
		Transport: &http.Transport{},
	}
}
