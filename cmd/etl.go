package etl

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go/service/s3"
	"go.mongodb.org/mongo-driver/mongo"
)

const linesInsert int = 1000

var wg = sync.WaitGroup{}

func insertFile(scanner *bufio.Scanner, coll *mongo.Collection, config Config) {
	var i = 0
	lines := make([]string, 0)

	for scanner.Scan() {
		newLine := scanner.Text()
		if i > 10 && newLine != "" && newLine != "  </body>" && newLine != "</tmx>" {
			lines = append(lines, scanner.Text())
		}
		if len(lines) == linesInsert*4 {
			go insertSentences(coll, config.DbName, getStructList(lines))
			lines = make([]string, 0)
		}
		i++
	}
	insertSentences(coll, config.DbName, getStructList(lines))
	wg.Done()
}

func loadFiles(files []string, mongoClient *mongo.Client, svc *s3.S3, config Config) {

	for _, file := range files {
		fmt.Println(file)
		coll := mongoClient.Database(config.DbName).Collection(strings.Split(file, "/")[0])

		body := getFileBody(config.Bucket, file, svc)
		defer body.Close()

		scanner := bufio.NewScanner(body)
		scanner.Split(bufio.ScanLines)

		wg.Add(1)
		go insertFile(scanner, coll, config)
		wg.Wait()
	}
}

func Start(config Config) {
	svc := connect(config.Region)
	mongoClient := connectMongo(config.MongoUri)
	defer func() {
		if err := mongoClient.Disconnect(context.TODO()); err != nil {
			log.Fatal(err.Error())
		}
	}()
	files := getTMXFilesNames(config.Bucket, svc)

	loadFiles(files, mongoClient, svc, config)
}
