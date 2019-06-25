package file

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/gofrs/uuid"

	"fileupload/db"

	"github.com/joho/godotenv"
)

// Test is just a function for testing
func Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": "success", "message": "testing 123", "requrl": c.Request.URL})
	return
}

// Upload file endpoint
func Upload(c *gin.Context) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"result": "error", "error": "Error"})
		return
	}

	file, header, err := c.Request.FormFile("file")

	// check if there's an error receiving file
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"result": "error", "error": "failed to upload file"})
		return
	}

	// server received file
	defer file.Close()

	// generate uuid
	uid, err := uuid.NewV4()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"result": "error", "error": "failed to generate file id"})
		return
	}

	// generate string uuid as file id
	fileID := uid.String()

	// store fileid in redis
	err = db.SetFile(&header.Filename, &fileID)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"result": "error", "error": "failed to store record into redis"})
		return
	}

	// upload file to s3
	// get file size
	size := header.Size
	buffer := make([]byte, size)
	file.Read(buffer)

	//err = S3Upload(file, header, &fileID)
	err = UploadS3(buffer, &fileID, &size, http.DetectContentType(buffer))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"result": "error", "error": "failed to upload to S3"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "success", "message": "uploaded", "url": fmt.Sprintf("%s/api/file/d/%s", os.Getenv("HOSTNAME"), fileID), "key": fileID})
	return
}

// Download a file fetched from S3
func Download(c *gin.Context) {
	fileid := c.Param("fileid")

	// download and stream file from S3 to responsewriter
	err := S3Download(&fileid, c)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"result": "error", "error": "failed to fetch file from S3"})
		return
	}
	// successfully sent file data
	// now delete record from redis and from s3
	/*
		errchn := make(chan error)
		go func() { errchn <- db.DelEntry(&fileid) }()
		go func(){ errchn <- S3Delete(&fileid) }()
	*/
	go db.DelEntry(&fileid)
	go S3Delete(&fileid)

	return
}
