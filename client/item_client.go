package client

import (
	"encoding/json"
	"fmt"
	"github.com/alganbr/kedai-itemsvc-client/models"
	"github.com/alganbr/kedai-utils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
	"net/http"
)

type IItemClient interface {
	Get(int64) (*models.Item, *errors.Error)
	Create(*models.ItemRq) (*models.Item, *errors.Error)
	Update(int64, *models.ItemRq) (*models.Item, *errors.Error)
	Patch(int64, *models.ItemRq) (*models.Item, *errors.Error)
}

type ItemClient struct {
	httpClient *rest.RequestBuilder
}

func (c *ItemClient) Get(id int64) (*models.Item, *errors.Error) {
	rs := c.httpClient.Get(fmt.Sprintf("/itemsvc/item/%d", id))
	if rs.Err != nil {
		return nil, &errors.Error{
			Code:    http.StatusInternalServerError,
			Message: rs.Err.Error(),
		}
	}
	if rs.StatusCode > 299 {
		var httpErr *errors.Error
		if err := json.Unmarshal(rs.Bytes(), &httpErr); err != nil {
			return nil, &errors.Error{
				Code:    http.StatusInternalServerError,
				Message: "Error when unmarshalling error response",
			}
		}
		return nil, httpErr
	}
	var item *models.Item
	if err := json.Unmarshal(rs.Bytes(), &item); err != nil {
		return nil, &errors.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error when unmarshalling response body",
		}
	}
	return item, nil
}

func (c *ItemClient) Create(rq *models.ItemRq) (*models.Item, *errors.Error) {
	rs := c.httpClient.Post("/itemsvc/item", rq)
	if rs.Response != nil {
		return nil, &errors.Error{
			Code:    http.StatusInternalServerError,
			Message: rs.Err.Error(),
		}
	}
	if rs.StatusCode > 299 {
		var httpErr *errors.Error
		if err := json.Unmarshal(rs.Bytes(), &httpErr); err != nil {
			return nil, &errors.Error{
				Code:    http.StatusInternalServerError,
				Message: "Error when unmarshalling error response",
			}
		}
		return nil, httpErr
	}
	var item *models.Item
	if err := json.Unmarshal(rs.Bytes(), &item); err != nil {
		return nil, &errors.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error when unmarshalling response body",
		}
	}
	return item, nil
}

func (c *ItemClient) Update(id int64, rq *models.ItemRq) (*models.Item, *errors.Error) {
	rs := c.httpClient.Put(fmt.Sprintf("/itemsvc/item/%d", id), rq)
	if rs == nil || rs.Response == nil {
		return nil, &errors.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error getting http response",
		}
	}
	if rs.StatusCode > 299 {
		var httpErr *errors.Error
		if err := json.Unmarshal(rs.Bytes(), &httpErr); err != nil {
			return nil, &errors.Error{
				Code:    http.StatusInternalServerError,
				Message: "Error when unmarshalling error response",
			}
		}
		return nil, httpErr
	}
	var item *models.Item
	if err := json.Unmarshal(rs.Bytes(), &item); err != nil {
		return nil, &errors.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error when unmarshalling response body",
		}
	}
	return item, nil
}

func (c *ItemClient) Patch(id int64, rq *models.ItemRq) (*models.Item, *errors.Error) {
	rs := c.httpClient.Patch(fmt.Sprintf("/itemsvc/item/%d", id), rq)
	if rs.Err != nil {
		return nil, &errors.Error{
			Code:    http.StatusInternalServerError,
			Message: rs.Err.Error(),
		}
	}
	if rs.StatusCode > 299 {
		var httpErr *errors.Error
		if err := json.Unmarshal(rs.Bytes(), &httpErr); err != nil {
			return nil, &errors.Error{
				Code:    http.StatusInternalServerError,
				Message: "Error when unmarshalling error response",
			}
		}
		return nil, httpErr
	}
	var item *models.Item
	if err := json.Unmarshal(rs.Bytes(), &item); err != nil {
		return nil, &errors.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error when unmarshalling response body",
		}
	}
	return item, nil
}
