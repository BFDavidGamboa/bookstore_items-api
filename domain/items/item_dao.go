package items

import (
	"encoding/json"
	"errors"
	"strings"

	"fmt"

	"github.com/BFDavidGamboa/bookstore_items-api/client/elasticsearch"
	"github.com/BFDavidGamboa/bookstore_utils-go/rest_errors"
)

const (
	//elasticsearch index
	indexItems = "items"

	// string typeItems also works with "item" value
	typeItems = "_doc"
)

func (i *Item) Save() rest_errors.RestErr {

	result, err := elasticsearch.Client.Index(indexItems, typeItems, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save items", errors.New("data error "))
	}
	i.Id = result.Id
	return nil
}

func (i *Item) Get() rest_errors.RestErr {
	itemId := i.Id
	result, err := elasticsearch.Client.Get(indexItems, typeItems, i.Id)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return rest_errors.NewNotFoundError(fmt.Sprintf("no item found with id %s", i.Id))
		}
		return rest_errors.NewInternalServerError(fmt.Sprintf("error when trying to get items by id %s", i.Id), errors.New("data error "))
	}

	bytes, err := result.Source.MarshalJSON()
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to parse database response", errors.New("database error"))
	}

	if err := json.Unmarshal(bytes, &i); err != nil {
		return rest_errors.NewInternalServerError("error when trying to marshal databas response", errors.New("database error"))
	}
	i.Id = itemId
	return nil
}
