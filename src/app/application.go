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
		Addr:    "127.0.0.1:8081",
	}

	fmt.Printf("Starting server on ip:port = %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
