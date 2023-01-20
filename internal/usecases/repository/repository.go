package repository

type Repository interface {
	InsertSentences(file string, docs []interface{})
}
