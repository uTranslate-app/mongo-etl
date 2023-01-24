package etl

import (
	"bufio"
	"strings"

	"github.com/uTranslate-app/uTranslate-api/internal/entities"
	"github.com/uTranslate-app/uTranslate-api/internal/usecases/extractor"
	"github.com/uTranslate-app/uTranslate-api/internal/usecases/repository"
)

type Loader struct {
	Extractor extractor.Extractor
	Rep       repository.Repository
}

func (l Loader) ToSentStruct(rawPair []string, file string) entities.Pair {
	raw_a := rawPair[1]
	raw_b := rawPair[2]

	lang_a := strings.Split(raw_a, "\"")[1]
	lang_b := strings.Split(raw_b, "\"")[1]

	sentence_a := strings.Split(strings.Split(raw_a, "<seg>")[1], "</seg>")[0]
	sentence_b := strings.Split(strings.Split(raw_b, "<seg>")[1], "</seg>")[0]

	sentType := strings.Split(file, "/")[0]

	return entities.Pair{
		entities.Sent{
			lang_a,
			sentence_a,
		},
		entities.Sent{
			lang_b,
			sentence_b,
		},
		sentType,
	}
}

func (l Loader) GetStructList(sentences []string, file string) []interface{} {
	pairs := make([]interface{}, 0)
	for i := 0; i < len(sentences)-2; i = i + 4 {
		pairs = append(pairs, l.ToSentStruct(sentences[i:i+4], file))
	}
	return pairs
}

func (l Loader) LoadLines() {
	bodies := l.Extractor.GetFilesBody()
	for file, body := range bodies {
		defer body.Close()

		scanner := bufio.NewScanner(body)
		scanner.Split(bufio.ScanLines)

		var i = 0
		lines := make([]string, 0)
		linesInsert := 10000

		for scanner.Scan() {
			newLine := scanner.Text()
			if i > 10 && newLine != "" && newLine != "  </body>" && newLine != "</tmx>" {
				lines = append(lines, scanner.Text())
			}
			if len(lines) == linesInsert*4 {
				l.Rep.InsertSentences(file, l.GetStructList(lines, file))
				lines = make([]string, 0)
			}
			i++
		}
		l.Rep.InsertSentences(file, l.GetStructList(lines, file))
	}
}
