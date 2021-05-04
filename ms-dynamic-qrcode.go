package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moises-ba/ms-dynamic-qrcode/config"
	"github.com/moises-ba/ms-dynamic-qrcode/model/domain/usecase/qrcodeservice"
	"github.com/moises-ba/ms-dynamic-qrcode/model/repository/mongorepo"
	"github.com/moises-ba/ms-dynamic-qrcode/utils"
	"github.com/moises-ba/ms-dynamic-qrcode/web/controller/qrcodecontroller"
	"github.com/moises-ba/ms-dynamic-qrcode/web/security/jwt"
)

func CORS(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusOK)
	}

}

func main() {

	//inicializando a conexão com o mongo
	mongoClient, funcDisconnect, err := mongorepo.Connect()
	if err != nil {
		log.Fatal("Falha ao conectar no mongo.", err)
	}
	defer funcDisconnect()

	//databases
	mongoQRCodeDB := mongoClient.Database(utils.GetEnv(config.MONGO_QRCODE_BD, "qrcodedb"))

	//repositories
	qrcodeRepository := mongorepo.NewRepository(mongoQRCodeDB)

	//services
	qrcodeService := qrcodeservice.NewService(qrcodeRepository)

	//inicializando e resgistrando endpoints
	r := gin.Default()
	r.Use(CORS)

	api := r.Group("/ms-dynamic-qrcode")

	qrCodeController := qrcodecontroller.NewController(qrcodeService)
	qrCodeGroup := api.Group("/qrcode")
	qrCodeGroup.GET("/list", jwt.Authorize(qrCodeController.List, "ADMIN", "USER"))
	qrCodeGroup.POST("/generate", jwt.Authorize(qrCodeController.Generate, "ADMIN", "USER"))

	r.Run(":8081")
}
