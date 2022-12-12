package services

import (
	"github.com/BFDavidGamboa/bookstore_items-api/domain/items"
	"github.com/BFDavidGamboa/bookstore_utils-go/rest_errors"
)

type itemsServiceInterface interface {
	Create(item items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)
}

type itemsService struct{}

func (s *itemsService) Create(item items.Item) (*items.Item, rest_errors.RestErr) {
	return nil, rest_errors.NewRestError()
}

func (s *itemsService) Get(id string) (*items.Item, rest_errors.RestErr) {
	return nil, rest_errors.NewRestError()
}
