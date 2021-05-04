package mongorepo

import "github.com/moises-ba/ms-dynamic-qrcode/model/domain"

type Reader interface {
	FindQRCodes(filter *domain.QRCodeFilter) ([]*domain.QRCodeModel, error)
}

type Writer interface {
	Insert(qrcode *domain.QRCodeModel) error
}

type Repository interface {
	Reader
	Writer
}
