package qrcodecontroller

import (
	"log"
	"net/http"

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

	if err := api.service.Insert(qrCode); err == nil {
		c.JSON(http.StatusOK, nil)
	} else {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
	}
}
