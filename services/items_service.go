package services

import (
	"net/http"

	"github.com/BFDavidGamboa/bookstore_items-api/domain/items"
	"github.com/BFDavidGamboa/bookstore_utils-go/rest_errors"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(item items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)
}

type itemsService struct{}

func (s *itemsService) Create(item items.Item) (*items.Item, rest_errors.RestErr) {
	return nil, rest_errors.NewRestError(
		"imlement me!",
		http.StatusNotImplemented,
		"not_implemented",
		nil,
	)
}

func (s *itemsService) Get(id string) (*items.Item, rest_errors.RestErr) {
	return nil, rest_errors.NewRestError("implement me!", http.StatusNotImplemented, "not_implemented", nil)
}
