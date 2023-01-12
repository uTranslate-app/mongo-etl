package etl

import (
	"bufio"
	"fmt"
	"strings"
)

const (
	linesInsert int    = 1000
	bucket      string = "utranslate-app"
	region      string = "sa-east-1"
	db          string = "uTranslate"
)

func Start() {
	svc := connect()
	mongoClient := connectMongo()
	files := getTMXFilesNames(bucket, svc)

	var lines []string
	var i int

	for _, file := range files {
		fmt.Println(file)
		collName := strings.Split(file, "/")[0]
		coll := mongoClient.Database(db).Collection(collName)
		i = 0

		body := getFileBody(bucket, file, svc)
		defer body.Close()

		scanner := bufio.NewScanner(body)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			newLine := scanner.Text()
			if i > 10 && newLine != "" && newLine != "  </body>" && newLine != "</tmx>" {
				lines = append(lines, scanner.Text())
			}
			if len(lines) == linesInsert*4 {
				insertSentences(coll, db, getStructList(lines))
				lines = make([]string, 0)
			}
			i++
		}
		insertSentences(coll, db, getStructList(lines))
		lines = make([]string, 0)
	}
}
