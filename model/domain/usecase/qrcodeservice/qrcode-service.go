package qrcodeservice

import (
	"encoding/base64"

	"github.com/moises-ba/ms-dynamic-qrcode/model/domain"
	"github.com/moises-ba/ms-dynamic-qrcode/model/domain/qrcodegenerator"
	"github.com/moises-ba/ms-dynamic-qrcode/model/repository/mongorepo"
	"github.com/moises-ba/ms-dynamic-qrcode/utils"
)

type service struct {
	repository mongorepo.Repository
}

func NewService(repo mongorepo.Repository) Service {

	return &service{
		repository: repo,
	}

}

func (s *service) FindQRCodes(filter *domain.QRCodeFilter) ([]*domain.QRCodeModel, error) {
	return s.repository.FindQRCodes(filter)
}
func (s *service) Insert(qrcode *domain.QRCodeModel) error {

	qrCodeContentIf := utils.GetValueByReflection(qrcode, qrcode.Type).(domain.QRCodeContent)

	contentString, err := (qrCodeContentIf).ToContentQRCode()
	if err != nil {
		return err
	}

	qrcode.Content = contentString

	qrcodeBytes, err := qrcodegenerator.GenerateQRCode(qrcode.Content)
	if err != nil {
		return err
	}

	qrcode.QrCodeInBase64 = base64.StdEncoding.EncodeToString(qrcodeBytes)

	return s.repository.Insert(qrcode)
}
