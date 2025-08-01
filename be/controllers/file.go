package controllers

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nchikkam/context-findr-be/model"
	utils "github.com/nchikkam/context-findr-be/utils/infrastructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// @Summary		List of all Uploaded files into Storage
// @Description	Uploaded files
// @Produce		json
// @Security 	BearerAuth
// @Success		200
// @Router		/api/v1/uploads [get]
func FileUploads(c *gin.Context) {

	email := c.GetString("userEmail")

	cursor, err := getFileCollection().Find(c, bson.M{"email": email})

	if err != nil {
		log.Fatalf("files collection err : %v", err)
		return
	}
	var queryResult []bson.M
	if err := cursor.All(c, &queryResult); err != nil {
		log.Fatalf("error query mongodb result")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"files": queryResult,
	})
}

// @Summary		Upload A Text File to Storage
// @Description	Upload file
// @ID			file.upload
// @Accept		multipart/form-data
// @Produce		json
// @Param		file formData file true "sample test text file"
// @Security 	BearerAuth
// @Success		200
// @Router		/api/v1/upload [post]
func FileUpload(c *gin.Context) {
	if !strings.Contains(c.Request.Header.Get("Content-Type"), "multipart/form-data") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Content-Type must be multipart/form-data"})
		return
	}

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file field is required"})
		return
	}
	defer file.Close()

	if header.Size == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "uploaded file is empty"})
		return
	}

	filetype, err := guessIncomingFileMimeType(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to identify file type"})
		return
	}

	if _, ok := utils.SupportedMIMEs[filetype]; !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unsupported file type", "type": filetype})
		return
	}

	if _, err := file.Seek(0, io.SeekStart); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to reset file reader"})
		return
	}

	rel_path, abs_path := normalizeFileName(header.Filename)

	if err := c.SaveUploadedFile(header, abs_path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to save file"})
		return
	}

	// persist upload record in db
	email := c.GetString("userEmail")
	newFile := model.File{
		ID:    primitive.NewObjectID(),
		Name:  rel_path,
		Email: strings.ToLower(email),
	}

	getFileCollection().DeleteMany(c, bson.M{"email": email}) // limit user to just 1 knowledge base file for now

	_, err = getFileCollection().InsertOne(c, newFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "file uploaded successfully",
		"filename": rel_path,
	})
}

func guessIncomingFileMimeType(f multipart.File) (string, error) {
	buffer := make([]byte, 512)
	if _, err := f.Read(buffer); err != nil && err != io.EOF {
		return "", err
	}
	return http.DetectContentType(buffer), nil
}

func normalizeFileName(filebase string) (string, string) {
	filename := filepath.Base(filebase)
	filename = filepath.Clean(filename)
	filename = strings.ReplaceAll(filename, " ", "_")

	timestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)

	newFilename := fmt.Sprintf("%s-%s", timestamp, filename)

	// todo: use classifier here later to place file in correct folder
	return newFilename, filepath.Join(utils.Store, newFilename)
}
