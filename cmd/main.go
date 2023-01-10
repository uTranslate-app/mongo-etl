package main

import (
	"bufio"
	"fmt"
)

const (
	linesInsert int    = 5
	bucket      string = "utranslate-app"
	region      string = "sa-east-1"
)

func main() {
	svc := connect()
	files := getTMXFilesNames(bucket, svc)

	// This is our buffer now
	var lines []string
	var i int = 0

	for _, file := range files {
		fmt.Println(file)
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
				lines = make([]string, 0)
			}
			i++
		}
	}
}
