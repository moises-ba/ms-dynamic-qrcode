package config

import (
	"github.com/moises-ba/ms-dynamic-qrcode/utils"
)

const (
	MONGO_SERVER_URL = "MONGO_SERVER_URL"
	MONGO_USER       = "MONGO_USER"
	MONGO_PASSWORD   = "MONGO_PASSWORD"
	MONGO_QRCODE_BD  = "MONGO_QRCODE_BD"
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
