package qrcodeservice

import (
	"github.com/moises-ba/ms-dynamic-qrcode/model/domain"
)

type Service interface {
	FindQRCodes(filter *domain.QRCodeFilter) ([]*domain.QRCodeModel, error)
	Insert(qrcode *domain.QRCodeModel) error
}
