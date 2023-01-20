package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/uTranslate-app/uTranslate-api/configs"
	"github.com/uTranslate-app/uTranslate-api/internal/usecases/etl"
)

func ServeRouter(loader *etl.Loader) {
	var c = new(controller)
	c.loader = *loader

	r := chi.NewRouter()
	r.Get("/", c.Etl)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.Cfg.Port), r)
}

type controller struct {
	loader etl.Loader
}

func (c controller) Etl(w http.ResponseWriter, r *http.Request) {
	c.loader.LoadLines()

	w.Write([]byte("hello"))
}
