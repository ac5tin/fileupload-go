package file

import (
	"bytes"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"

	"mime/multipart"
	"net/http"
)

// S3Upload uploads file to S3
func S3Upload(file multipart.File, fileHeader *multipart.FileHeader, filename *string) error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	// get file size
	size := fileHeader.Size
	buffer := make([]byte, size)
	file.Read(buffer)

	s, err := session.NewSession(&aws.Config{Region: aws.String(os.Getenv("S3_REGION"))})
	if err != nil {
		return err
	}

	_, err = s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(os.Getenv("S3_BUCKET")),
		Key:           aws.String(*filename),
		Body:          bytes.NewReader(buffer),
		ACL:           aws.String("public-read"),
		ContentLength: aws.Int64(int64(size)),
		ContentType:   aws.String(http.DetectContentType(buffer)),
	})

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
