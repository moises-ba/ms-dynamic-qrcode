package qrcodegenerator

import qrcode "github.com/skip2/go-qrcode"

func GenerateQRCodeWithSize(content string, size int) ([]byte, error) {

	png, err := qrcode.Encode(content, qrcode.Medium, size)

	return png, err
}

func GenerateQRCode(content string) ([]byte, error) {

	return GenerateQRCodeWithSize(content, 256)
}
