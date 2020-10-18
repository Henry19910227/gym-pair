package controller

import (
	"net/http"

	"github.com/Henry19910227/gym-pair/internal/service"
	"github.com/gin-gonic/gin"
)

// UploadController ...
type UploadController struct {
	UploadService service.UploadService
}

// NewUploadController ...
func NewUploadController(router *gin.Engine, s service.UploadService) {
	uploadController := &UploadController{
		UploadService: s,
	}
	v1 := router.Group("/gympair/v1")
	v1.POST("/upload/file", uploadController.Upload)

}

// Upload ...
func (vc *UploadController) Upload(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": err.Error()})
		return
	}
	filename, err := vc.UploadService.UploadImage(file, fileHeader.Filename)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": filename, "msg": "upload success!"})
}
