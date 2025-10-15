package controller

import (
	"context"
	"go-single-file-upload/config"
	"go-single-file-upload/model"
	"go-single-file-upload/storage"
	"log"
	"net/http"
	"time"

	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)

type FileController struct {
	db *storage.DB
	cld *storage.Cloudinary
	conf *config.Cloudinary
}

func InitFileController(db *storage.DB, cld *storage.Cloudinary, conf *config.Cloudinary) *FileController {
	return &FileController{
		db: db,
		cld: cld,
		conf: conf,
	}
}

func (fc *FileController) UploadFile(c *gin.Context) {
	// single file
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "file is required", "error": err.Error()})
		return
	}

	// open uploaded file
	f, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "cannot open file", "error": err.Error()})
		return
	}
	defer f.Close()

	uploadParams := uploader.UploadParams{
		PublicID:       "",          // let Cloudinary pick or set if you want e.g., "tzuri/" + file.Filename
		Folder:         fc.conf.Folder,
		UseFilename:    api.Bool(true), // store original filename
		UniqueFilename: api.Bool(true),
		Overwrite:      api.Bool(false),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// stream upload
	cld := fc.cld.GetCld()

	uploadResult, err := cld.Upload.Upload(ctx, f, uploadParams)
	if err != nil {
		log.Println("cloudinary upload error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "failed to upload to cloudinary", "error": err.Error()})
		return
	}

	// store metadata
	rec := &model.FileRecord{
		PublicID: uploadResult.PublicID,
		URL:      uploadResult.SecureURL,
		Filename: file.Filename,
		Bytes:    file.Size,
	}

	db := fc.db.GetDB()

	if err := db.Create(&rec).Error; err != nil {
		log.Println("db create error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "failed to save metadata", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    rec,
		"raw":     uploadResult,
	})
}

func (fc *FileController) GetFiles(c *gin.Context) {
	var rows []model.FileRecord
	db := fc.db.GetDB()
	if err := db.Order("created_at desc").Find(&rows).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "db error", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": rows})
}
