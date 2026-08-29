package api

import (
	"fmt"
	"net/http"

	"github.com/yehezkiel1086/go-concurrent-request-aggregator/internal/adapter/config"
)

type API struct {
	client *http.Client
	base   string
}

func New(conf *config.API) *API {
	client := &http.Client{}
	base := conf.Base

	return &API{
		client,
		base,
	}
}

func (a *API) Get(path string) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", a.base, path)
	return a.client.Get(url)
}
