package apis

import "net/http"

type Client http.Client

func NewHTTPCLient() *http.Client {
	return &http.Client{}
}
