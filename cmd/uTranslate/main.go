package main

import (
	"github.com/uTranslate-app/uTranslate-api/api/v1/router"
	"github.com/uTranslate-app/uTranslate-api/configs"
)

func main() {
	configs.LoadConfig(".")

	router.ServeRouter()
}
