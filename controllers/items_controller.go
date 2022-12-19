package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/BFDavidGamboa/bookstore_items-api/domain/items"
	"github.com/BFDavidGamboa/bookstore_items-api/services"
	"github.com/BFDavidGamboa/bookstore_items-api/utils/http_utils"
	"github.com/BFDavidGamboa/bookstore_oauth-go/oauth"
	"github.com/BFDavidGamboa/bookstore_utils-go/rest_errors"
	"github.com/gorilla/mux"
)

var (
	ItemsController itemsControllerIntercafe = &itemsController{}
)

type itemsControllerIntercafe interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsController struct{}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {

	if err := oauth.AuthenticateRequest(r); err != nil {

		return
	}

	sellerId := oauth.GetCallerID(r)
	if sellerId == 0 {
		respErr := rest_errors.NewUnauthorizedError("unable to retrieve user information from given access_token")
		http_utils.RespondError(w, respErr)
		return
	}

	requestBoody, err := io.ReadAll(r.Body)
	if err != nil {
		respErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.RespondError(w, respErr)
		return
	}
	defer r.Body.Close()

	var itemRequest items.Item
	if err := json.Unmarshal(requestBoody, &itemRequest); err != nil {
		respErr := rest_errors.NewBadRequestError("invalid item json  body")
		http_utils.RespondError(w, respErr)
		return
	}

	itemRequest.Seller = sellerId

	result, createErr := services.ItemsService.Create(itemRequest)
	if createErr != nil {
		http_utils.RespondError(w, createErr)
		return
	}
	http_utils.RespondJson(w, http.StatusCreated, result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := strings.TrimSpace(vars["id"])

	item, err := services.ItemsService.Get(itemId)
	if err != nil {
		http_utils.RespondError(w, err)
		return
	}
	http_utils.RespondJson(w, http.StatusOK, item)
}
