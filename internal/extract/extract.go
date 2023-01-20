package extract

import (
	"io"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/uTranslate-app/uTranslate-api/configs"
)

type ExtractS3 struct {
	Bucket string
	Region string
}

func (es3 ExtractS3) Connect() *s3.S3 {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(es3.Region)},
	)
	if err != nil {
		log.Fatalf("Unable to connect, %v", err)
	}

	return s3.New(sess)
}

func (es3 ExtractS3) GetTMXFilesNames() []string {
	var TMXFilesNames []string

	sess := es3.Connect()
	resp, err := sess.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(configs.Cfg.Bucket)})
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

func (es3 ExtractS3) GetFilesBody() map[string]io.ReadCloser {
	FilesBodies := make(map[string]io.ReadCloser)

	sess := es3.Connect()
	filesNames := es3.GetTMXFilesNames()

	for _, TMXFile := range filesNames {
		requestInput := &s3.GetObjectInput{
			Bucket: aws.String(configs.Cfg.Bucket),
			Key:    aws.String(TMXFile),
		}
		result, err := sess.GetObject(requestInput)
		if err != nil {
			log.Fatalf("Error %v", err)
		}
		FilesBodies[TMXFile] = result.Body
	}
	return FilesBodies
}
