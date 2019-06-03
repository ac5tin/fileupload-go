package file

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gofrs/uuid"

    "fileupload/db"
)

// Test is just a function for testing
func Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": "success", "message": "testing 123"})
	return
}

// Upload file endpoint
func Upload(c *gin.Context) {
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
    err = db.SetFile(&header.Filename,&fileID)
    if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"result": "error", "error": "failed to store record into redis"})
		return
    }

    // upload file to s3
	err = S3Upload(file, header, &fileID)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"result": "error", "error": "failed to upload to S3"})
		return
	}


	c.JSON(http.StatusOK, gin.H{"result": "success", "message": "uploaded"})
	return
}
