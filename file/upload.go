package file

import (
	"bytes"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"

	"net/http"
)

// UploadS3 uploads file to S3
func UploadS3(buffer []byte, filename *string, size *int64, ctype string) error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	s, err := session.NewSession(&aws.Config{Region: aws.String(os.Getenv("S3_REGION"))})
	if err != nil {
		return err
	}

	_, err = s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(os.Getenv("S3_BUCKET")),
		Key:           aws.String(*filename),
		Body:          bytes.NewReader(buffer),
		ACL:           aws.String("public-read"),
		ContentLength: aws.Int64(*size),
		ContentType:   aws.String(http.DetectContentType(buffer)),
	})

	return nil
}
