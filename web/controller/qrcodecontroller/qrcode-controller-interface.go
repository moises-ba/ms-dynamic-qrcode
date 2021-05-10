package qrcodecontroller

import (
	"github.com/gin-gonic/gin"
)

type QRCodeAPI interface {
	List(c *gin.Context)
	Generate(c *gin.Context)
	Upload(c *gin.Context)
	Delete(c *gin.Context)
}
