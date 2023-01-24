package router

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/uTranslate-app/uTranslate-api/configs"
	"github.com/uTranslate-app/uTranslate-api/internal/usecases/etl"
	"github.com/uTranslate-app/uTranslate-api/internal/usecases/retriever"
)

func ServeRouter(loader *etl.Loader, ret *retriever.Retriever) {
	var c = new(controller)
	c.loader = *loader
	c.ret = *ret

	r := chi.NewRouter()
	r.Get("/", c.Etl)
	r.Get("/langs", c.getLangs)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.Cfg.Port), r)
}

type controller struct {
	loader etl.Loader
	ret    retriever.Retriever
}

func (c controller) Etl(w http.ResponseWriter, r *http.Request) {
	c.loader.LoadLines()

	w.Write([]byte("hello"))
}

func (c controller) getLangs(w http.ResponseWriter, r *http.Request) {
	langUsed := r.URL.Query().Get("lang")
	availableLangs := c.ret.GetLangs(langUsed)
	test := strings.Join(availableLangs, " ")
	w.Write([]byte(test))
}
