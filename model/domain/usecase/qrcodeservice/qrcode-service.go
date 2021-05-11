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

func (s *service) FindQRCodes(filter *domain.QRCodeFilter) ([]*domain.QRCodeResponse, error) {

	qrCodes, err := s.repository.FindQRCodes(filter)
	if err != nil {
		return nil, err
	}

	total := len(qrCodes)
	if total > 0 {
		qrCodesResponse := make([]*domain.QRCodeResponse, total)
		var response *domain.QRCodeResponse
		for i, v := range qrCodes {
			response = new(domain.QRCodeResponse)
			response.QRCodeModel = *v
			response.IsImage = v.IsImage()
			response.Dynamic = v.IsDynamic()
			qrCodesResponse[i] = response
		}

		return qrCodesResponse, nil
	}

	return make([]*domain.QRCodeResponse, 0), nil
}
func (s *service) Insert(qrcode *domain.QRCodeModel) (*domain.QRCodeResponse, error) {

	qrCodeContentIf := utils.GetValueByReflection(qrcode, qrcode.Type).(domain.QRCodeContent)

	contentString, err := (qrCodeContentIf).ToContentQRCode()
	if err != nil {
		return nil, err
	}

	if qrcode.IsDynamic() {
		contentString += qrcode.Uuid
	}

	qrcode.Content = contentString

	qrcodeBytes, err := qrcodegenerator.GenerateQRCode(qrcode.Content)
	if err != nil {
		return nil, err
	}

	qrcode.QrCodeInBase64 = base64.StdEncoding.EncodeToString(qrcodeBytes)

	errMongo := s.repository.Insert(qrcode)
	if errMongo != nil {
		return nil, errMongo
	}

	return &domain.QRCodeResponse{QRCodeModel: *qrcode,
		IsImage: qrcode.IsImage(),
		Dynamic: qrcode.IsDynamic()}, nil
}

func (s *service) Delete(pFilter *domain.QRCodeFilter) error {
	return s.repository.Delete(pFilter)
}
