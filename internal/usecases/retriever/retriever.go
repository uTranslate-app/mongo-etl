package retriever

import (
	"github.com/uTranslate-app/uTranslate-api/internal/usecases/repository"
)

type Retriever struct {
	Rep repository.Repository
}

func (r Retriever) GetLangs(langUsed string) []string {
	return r.Rep.GetMongoLangs(langUsed)
}
