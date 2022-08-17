package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/luccascleann/projeto_api/configs"
	"github.com/luccascleann/projeto_api/handlers"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Post("/user", handlers.Create)
	r.Post("/user/createWithArray", handlers.CreateWithArray)
	r.Put("/user/{username}", handlers.Update)
	r.Delete("/user/{username}", handlers.Delete)
	r.Get("/user", handlers.List)
	r.Get("/user/{username}", handlers.Get)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
