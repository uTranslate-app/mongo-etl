package repository

type Repository interface {
	InsertSentences(file string, docs []interface{})
	GetMongoLangs(lang string) []string
}
