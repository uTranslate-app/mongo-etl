package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/uTranslate-app/uTranslate-api/api/v1/handlers"
	"github.com/uTranslate-app/uTranslate-api/configs"
)

func ServeRouter() {
	r := chi.NewRouter()

	r.Get("/", handlers.Etl)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.Cfg.Port), r)
}
