package services

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/uTranslate-app/uTranslate-api/db"
	"github.com/uTranslate-app/uTranslate-api/internal/extract"
)

func insertFile(scanner *bufio.Scanner, file string) {

	collName := strings.Split(file, "/")[0]

	var i = 0
	lines := make([]string, 0)
	linesInsert := 10000

	for scanner.Scan() {
		newLine := scanner.Text()
		if i > 10 && newLine != "" && newLine != "  </body>" && newLine != "</tmx>" {
			lines = append(lines, scanner.Text())
		}
		if len(lines) == linesInsert*4 {
			db.InsertSentences(collName, db.GetStructList(lines))
			lines = make([]string, 0)
		}
		i++
	}
	db.InsertSentences(collName, db.GetStructList(lines))
}

func StartEtl() {
	files := extract.GetTMXFilesNames()

	for _, file := range files {
		fmt.Println(file)

		body := extract.GetFileBody(file)
		defer body.Close()

		scanner := bufio.NewScanner(body)
		scanner.Split(bufio.ScanLines)

		insertFile(scanner, file)
	}
}
