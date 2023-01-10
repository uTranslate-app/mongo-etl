package main

import (
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func connect() *s3.S3 {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("sa-east-1")},
	)
	if err != nil {
		ExitErrorf("Unable to connect, %v", err)
	}

	return s3.New(sess)
}

func getTMXFilesNames(bucket string, svc *s3.S3) []string {
	var TMXFilesNames []string

	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(bucket)})
	if err != nil {
		ExitErrorf("Unable to list items in bucket %q, %v", bucket, err)
	}
	for _, item := range resp.Contents {
		if *item.Size != 0 {
			TMXFilesNames = append(TMXFilesNames, *item.Key)
		}
	}
	return TMXFilesNames
}

func getFileBody(bucket string, TMXFile string, svc *s3.S3) io.ReadCloser {
	requestInput := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(TMXFile),
	}
	result, err := svc.GetObject(requestInput)
	if err != nil {
		fmt.Println(err)
	}

	return result.Body
}
