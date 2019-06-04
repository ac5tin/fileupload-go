package file

import (
	"bytes"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

// S3Download returns numbytes from S3 bucket
func S3Download(fileID *string, c *gin.Context) error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	s, err := session.NewSession(&aws.Config{Region: aws.String(os.Getenv("S3_REGION"))})
	if err != nil {
		return err
	}

	buff := &aws.WriteAtBuffer{}

	downloader := s3manager.NewDownloader(s)
	_, err = downloader.Download(buff, &s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET")),
		Key:    aws.String(*fileID),
	})
	if err != nil {
		return err
	}
	// write buffer to responsewriter
	_, err = io.Copy(c.Writer, bytes.NewReader(buff.Bytes()))
	if err != nil {
		return err
	}

	return nil

}
