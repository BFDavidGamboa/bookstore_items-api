package app

import (
	"net/http"

	"github.com/BFDavidGamboa/bookstore_items-api/controllers"
)

func mapUrls() {

	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)

	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
	router.HandleFunc("/items/{id}", controllers.ItemsController.Get).Methods(http.MethodGet)
}
