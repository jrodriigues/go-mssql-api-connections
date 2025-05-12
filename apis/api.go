package apis

import (
	"bytes"
	"errors"
	"net/http"
)

type Api struct {
	Host   string
	ApiKey string
}

func NewApi(host, apiKey string) (*Api, error) {
	if host == "" || apiKey == "" {
		return nil, errors.New("all fields are required")
	}

	api := &Api{
		Host:   host,
		ApiKey: apiKey,
	}

	return api, nil
}

func (api *Api) UrlForEndpoint(endpoint string, params map[string]string) string {
	url := api.Host + "/" + endpoint
	if len(params) > 0 {
		url += "?"
		for k, v := range params {
			url += k + "=" + v + "&"
		}
	}

	return url
}

func (api *Api) NewRequest(method, url string, data []byte, headers map[string]string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	return req, nil
}
