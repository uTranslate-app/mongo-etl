package etl

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"strings"
)

const linesInsert int = 1000

func Start(config Config) {
	svc := connect(config.Region)
	mongoClient := connectMongo(config.MongoUri)
	defer func() {
		if err := mongoClient.Disconnect(context.TODO()); err != nil {
			log.Fatal(err.Error())
		}
	}()
	files := getTMXFilesNames(config.Bucket, svc)

	var lines []string
	var i int

	for _, file := range files {
		fmt.Println(file)
		collName := strings.Split(file, "/")[0]
		coll := mongoClient.Database(config.DbName).Collection(collName)
		i = 0

		body := getFileBody(config.Bucket, file, svc)
		defer body.Close()

		scanner := bufio.NewScanner(body)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			newLine := scanner.Text()
			if i > 10 && newLine != "" && newLine != "  </body>" && newLine != "</tmx>" {
				lines = append(lines, scanner.Text())
			}
			if len(lines) == linesInsert*4 {
				insertSentences(coll, config.DbName, getStructList(lines))
				lines = make([]string, 0)
			}
			i++
		}
		insertSentences(coll, config.DbName, getStructList(lines))
		lines = make([]string, 0)
	}

}
