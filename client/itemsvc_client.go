package client

import "github.com/mercadolibre/golang-restclient/rest"

type IItemSvcClient interface {
	Item() IItemClient
}

type ItemSvcClient struct {
	HttpClient *rest.RequestBuilder
}

func (c ItemSvcClient) Item() IItemClient {
	return &ItemClient{c.HttpClient}
}
