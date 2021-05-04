package qrcodecontroller

import (
	"github.com/gin-gonic/gin"
)

type QRCodeAPI interface {
	List(c *gin.Context)
	Generate(c *gin.Context)
}
