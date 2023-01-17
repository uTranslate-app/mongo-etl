package extract

import (
	"io"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/uTranslate-app/uTranslate-api/configs"
)

func connect(region string) *s3.S3 {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		log.Fatalf("Unable to connect, %v", err)
	}

	return s3.New(sess)
}

func GetTMXFilesNames() []string {
	var TMXFilesNames []string

	svc := connect(configs.Cfg.Region)
	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(configs.Cfg.Bucket)})
	if err != nil {
		log.Fatalf("Unable to list items in bucket %q, %v", configs.Cfg.Bucket, err)
	}

	for _, item := range resp.Contents {
		if *item.Size != 0 {
			TMXFilesNames = append(TMXFilesNames, *item.Key)
		}
	}
	return TMXFilesNames
}

func GetFileBody(TMXFile string) io.ReadCloser {
	svc := connect(configs.Cfg.Region)

	requestInput := &s3.GetObjectInput{
		Bucket: aws.String(configs.Cfg.Bucket),
		Key:    aws.String(TMXFile),
	}
	result, err := svc.GetObject(requestInput)
	if err != nil {
		log.Fatalf("Error %v", err)
	}

	return result.Body
}
