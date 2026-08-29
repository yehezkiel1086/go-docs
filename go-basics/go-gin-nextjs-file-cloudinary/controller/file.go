package controller

import (
	"context"
	"go-gin-nextjs-file-cloudinary/config"
	"go-gin-nextjs-file-cloudinary/config/cloudinary"
	"go-gin-nextjs-file-cloudinary/config/postgres"
	"go-gin-nextjs-file-cloudinary/model"
	"log"
	"net/http"
	"time"

	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"

	"github.com/gin-gonic/gin"
)

type FileController struct {
	cld *cloudinary.Cloudinary
	db *postgres.DB
	conf *config.Cloudinary
}

func InitFileController(conf *config.Cloudinary, cld *cloudinary.Cloudinary, db *postgres.DB) *FileController {
	return &FileController{
		conf: conf,
		cld: cld,
		db: db,
	}
}

func (fc *FileController) UploadFile(c *gin.Context) {
	// limit file size if desired: e.g., 10MB
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 10<<20)

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "file is required", "error": err.Error()})
		return
	}

	f, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "cannot open file", "error": err.Error()})
		return
	}
	defer f.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	uploadParams := uploader.UploadParams{
		Folder:         fc.conf.Folder,
		UseFilename:    api.Bool(true),
		UniqueFilename: api.Bool(true),
		Overwrite:      api.Bool(false),
	}

	res, err := fc.cld.Upload.Upload(ctx, f, uploadParams)
	if err != nil {
		log.Println("cloudinary upload error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "failed to upload to cloudinary", "error": err.Error()})
		return
	}

	rec := &model.FileRecord{
		PublicID: res.PublicID,
		URL:      res.SecureURL,
		Filename: file.Filename,
		Bytes:    file.Size,
	}

	db := fc.db.GetDB()

	if err := db.Create(&rec).Error; err != nil {
		log.Println("db create error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "failed to save metadata", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": rec})
}

func (fc *FileController) GetUploads(c *gin.Context) {
	var rows []model.FileRecord

	db := fc.db.GetDB()

	if err := db.Order("created_at desc").Find(&rows).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "db error", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": rows})
}
