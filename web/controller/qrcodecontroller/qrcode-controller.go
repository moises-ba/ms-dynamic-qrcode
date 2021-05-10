package qrcodecontroller

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
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
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
	}

}

func (api *qrcodeApi) Generate(c *gin.Context) {
	user := c.Keys[(utils.UserParamName)].(*model.PrincipalUserDetail)

	qrCode := &domain.QRCodeModel{}
	qrCode.User = user.Login

	if err := c.BindJSON(qrCode); err != nil {
		log.Println(err)
		return
	}

	qrcodeResponse, err := api.service.Insert(qrCode)
	if err == nil {
		c.JSON(http.StatusOK, qrcodeResponse)
	} else {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
	}
}

func (api *qrcodeApi) Upload(c *gin.Context) {

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Erro ao receber arquivo: %s", err.Error()))
		return
	}

	filePathStore := "/tmp/"
	filename := filepath.Base(filePathStore + file.Filename)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Arquivo %s salvo.", file.Filename)})
}
