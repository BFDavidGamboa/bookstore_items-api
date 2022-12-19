package items

import (
	"errors"

	"github.com/BFDavidGamboa/bookstore_items-api/client/elasticsearch"
	"github.com/BFDavidGamboa/bookstore_utils-go/rest_errors"
)

const (
	//elasticsearch index
	indexItems = "items"
)

func (i *Item) Save() rest_errors.RestErr {

	result, err := elasticsearch.Client.Index(indexItems, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save items", errors.New("data error "))
	}
	i.Id = result.Id
	return nil
}
