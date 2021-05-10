package qrcodeservice

import (
	"github.com/moises-ba/ms-dynamic-qrcode/model/domain"
)

type Service interface {
	FindQRCodes(filter *domain.QRCodeFilter) ([]*domain.QRCodeResponse, error)
	Insert(qrcode *domain.QRCodeModel) (*domain.QRCodeResponse, error)
}
