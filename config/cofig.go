package config

import (
	"github.com/moises-ba/ms-dynamic-qrcode/utils"
)

const (
	MONGO_SERVER_URL         = "MONGO_SERVER_URL"
	MONGO_USER               = "MONGO_USER"
	MONGO_PASSWORD           = "MONGO_PASSWORD"
	MONGO_QRCODE_BD          = "MONGO_QRCODE_BD"
	QRCODE_BASE_DYNAMIC_URL  = "QRCODE_BASE_DYNAMIC_URL" //url base para os qrcodes dinamicos
	QRCODE_VOLUME_STORE_PATH = "QRCODE_VOLUME_STORE_PATH"
)

func GetMogoServerURL() string {
	return utils.GetEnv(MONGO_SERVER_URL, "mongodb://localhost:27017")
}

func GetMongoUser() string {
	return utils.GetEnv(MONGO_USER, "root")
}

func GetMongoPassWord() string {
	return utils.GetEnv(MONGO_PASSWORD, "example")
}

func GetURLBaseDymamicQRCode() string {
	return utils.GetEnv(QRCODE_BASE_DYNAMIC_URL, "http://localhost:4200/")
}

func GetQRCodeVolumeStorePath() string {
	return utils.GetEnv(QRCODE_VOLUME_STORE_PATH, "/tmp/")
}
