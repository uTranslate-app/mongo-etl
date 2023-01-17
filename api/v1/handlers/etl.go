package handlers

import (
	"net/http"

	"github.com/uTranslate-app/uTranslate-api/internal/services"
)

func Etl(w http.ResponseWriter, r *http.Request) {
	services.StartEtl()

	w.Write([]byte("hello"))
}
