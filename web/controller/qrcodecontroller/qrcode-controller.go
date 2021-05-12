package qrcodecontroller

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/moises-ba/ms-dynamic-qrcode/config"
	"github.com/moises-ba/ms-dynamic-qrcode/model/domain"
	"github.com/moises-ba/ms-dynamic-qrcode/model/domain/usecase/qrcodeservice"
	"github.com/moises-ba/ms-dynamic-qrcode/utils"
	"github.com/moises-ba/ms-dynamic-qrcode/web/security/jwt/model"
)

type qrcodeApi struct {
	service qrcodeservice.Service
}

func NewController(pService qrcodeservice.Service) QRCodeAPI {
	return &qrcodeApi{
		service: pService,
	}
}

func (api *qrcodeApi) List(c *gin.Context) {

	user := c.Keys[(utils.UserParamName)].(*model.PrincipalUserDetail)

	if qrCodes, err := api.service.FindQRCodes(&domain.QRCodeFilter{User: user.Login}); err == nil {
		c.JSON(http.StatusOK, qrCodes)
	} else {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
	}

}

func (api *qrcodeApi) Generate(c *gin.Context) {
	user := c.Keys[(utils.UserParamName)].(*model.PrincipalUserDetail)

	qrCode := &domain.QRCodeModel{}
	qrCode.User = user.Login

	if err := c.BindJSON(qrCode); err != nil {
		log.Error(err)
		return
	}

	qrcodeResponse, err := api.service.Insert(qrCode)
	if err == nil {
		c.JSON(http.StatusOK, qrcodeResponse)
	} else {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
	}
}

func (api *qrcodeApi) Upload(c *gin.Context) {
	user := c.Keys[(utils.UserParamName)].(*model.PrincipalUserDetail)
	// Source
	file, err := c.FormFile("file")
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Erro ao receber arquivo: %s", err.Error()))
		return
	}

	fullFilePathDest := config.GetQRCodeVolumeStorePath() + user.Login + "/"

	err = os.MkdirAll(fullFilePathDest, 0755)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Nao foi possivel criar arquivo: %s", err.Error()))
	}

	fullFilePathDest += file.Filename

	if err := c.SaveUploadedFile(file, fullFilePathDest); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"filePath": fullFilePathDest})
}

func (api *qrcodeApi) Delete(c *gin.Context) {
	user := c.Keys[(utils.UserParamName)].(*model.PrincipalUserDetail)

	qrCode := &domain.QRCodeFilter{}
	qrCode.User = user.Login
	qrCode.Uuid = c.Param("qrcodeid")

	err := api.service.Delete(qrCode)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{})
	} else {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
	}
}
