package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	r "github.com/redis/go-redis/v9"
	"github.com/yehezkiel1086/go-photo-api-redis/model"
	"github.com/yehezkiel1086/go-photo-api-redis/storage/redis"
)

type PhotoController struct {
	redis *redis.Redis
}

func InitPhotoController(ctx context.Context, redisStorage *redis.Redis) *PhotoController {
	return &PhotoController{
		redis: redisStorage,
	}
}

func (pc *PhotoController) GetPhotos(c *gin.Context) {
	pc.getAndCachePhotos(c, "photos", "https://jsonplaceholder.typicode.com/photos")
}

func (pc *PhotoController) GetPhotosByAlbumID(c *gin.Context) {
	albumId := c.Param("albumId")
	if albumId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "albumId is required",
		})
		return
	}

	cachedKey := fmt.Sprintf("photos:%s", albumId)
	apiUrl := fmt.Sprintf("https://jsonplaceholder.typicode.com/photos?albumId=%s", albumId)

	pc.getAndCachePhotos(c, cachedKey, apiUrl)
}

func (pc *PhotoController) getAndCachePhotos(c *gin.Context, cacheKey, apiUrl string) {
	client := pc.redis.GetDB()
	ctx := c.Request.Context()

	var photos []model.Photo

	cachedPhotos, err := client.Get(ctx, cacheKey).Result()
	if err == nil {
		log.Println("✅ photos fetched from cache")
		if err := json.Unmarshal([]byte(cachedPhotos), &photos); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to unmarshal cached photos: " + err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, photos)
		return
	}
	
	if err != r.Nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "redis error: " + err.Error(),
		})
		return
	}

	log.Println("⚠️ photos cache miss, fetching from API")
	resp, err := http.Get(apiUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch from external API: " + err.Error(),
		})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to read API response body: " + err.Error(),
		})
		return
	}

	if err := json.Unmarshal(body, &photos); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to unmarshal API response: " + err.Error(),
		})
		return
	}

	if err := client.Set(ctx, cacheKey, body, 0).Err(); err != nil {
		log.Printf("failed to set cache: %v\n", err)
	}

	c.JSON(http.StatusOK, photos)
}