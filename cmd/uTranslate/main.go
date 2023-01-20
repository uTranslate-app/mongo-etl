package main

import (
	"github.com/uTranslate-app/uTranslate-api/api/v1/router"
	"github.com/uTranslate-app/uTranslate-api/configs"
	"github.com/uTranslate-app/uTranslate-api/internal/extract"
	"github.com/uTranslate-app/uTranslate-api/internal/gateways/mongo"
	"github.com/uTranslate-app/uTranslate-api/internal/usecases/etl"
)

func main() {
	configs.LoadConfig(".")

	var s3extractor = new(extract.ExtractS3)
	s3extractor.Bucket = configs.Cfg.Bucket
	s3extractor.Region = configs.Cfg.Region

	var mongoRepo = new(mongo.MongoDb)
	mongoRepo.Uri = configs.Cfg.MongoUri

	var loader = new(etl.Loader)
	loader.Extractor = s3extractor
	loader.Rep = mongoRepo

	router.ServeRouter(loader)
}
