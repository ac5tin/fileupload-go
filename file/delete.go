package file

import (
	"os"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
)

// S3Delete will delete a file (given key) from the s3 bucket
func S3Delete (fileID *string)error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	s, err := session.NewSession(&aws.Config{Region: aws.String(os.Getenv("S3_REGION"))})
	if err != nil {
		return err
	}

	_,err = s3.New(s).DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET")),
		Key:    aws.String(*fileID),
	})

	if err != nil {
		return err
	}

	return nil
}