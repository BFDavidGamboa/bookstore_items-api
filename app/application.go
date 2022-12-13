package app

import (
	"fmt"
	"net/http"

	"github.com/BFDavidGamboa/bookstore_items-api/client/elasticsearch"
	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	elasticsearch.Init()
	mapUrls()

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8080",
	}

	fmt.Println("Starting server on port 8080")
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
