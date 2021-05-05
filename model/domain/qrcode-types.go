package domain

import "encoding/json"

//import "go.mongodb.org/mongo-driver/bson/primitive"

type QRCodeContent interface {
	ToContentQRCode() (string, error)
}

type StringQRCodeField string

func (c StringQRCodeField) ToContentQRCode() (string, error) {
	return string(c), nil
}

type QRCodeFilter struct {
	User string
}

type QRCodeModel struct {
	Uuid           string             `bson:"_id,omitempty" json:"uuid,omitempty"`
	User           string             `bson:"user,omitempty" json:"user,omitempty"`
	QrCodeInBase64 string             `bson:"qrCodeInBase64,omitempty" json:"qrCodeInBase64,omitempty"`
	Content        string             `bson:"content,omitempty" json:"content,omitempty"`
	Dynamic        bool               `bson:"dynamic,omitempty" json:"dynamic,omitempty"`
	FilePath       string             `bson:"filePath,omitempty" json:"filePath,omitempty"`
	FileBase64     string             `bson:"fileBase64,omitempty" json:"fileBase64,omitempty"`
	Type           string             `bson:"type,omitempty" json:"type,omitempty"`
	IsImage        bool               `bson:"isImage,omitempty" json:"isImage,omitempty"`
	CustomFields   []*CustomField     `bson:"customFields,omitempty" json:"customFields,omitempty"`
	Url            *StringQRCodeField `bson:"url,omitempty" json:"url,omitempty"`
	Text           *StringQRCodeField `bson:"text,omitempty" json:"text,omitempty"`
	Vcard          *VCardField        `bson:"vcard,omitempty" json:"vcard,omitempty"`
	Email          *EmailField        `bson:"email,omitempty" json:"email,omitempty"`
	Wifi           *WIFIField         `bson:"wifi,omitempty" json:"wifi,omitempty"`
	Bitcoin        *BitCoinField      `bson:"bitcoin,omitempty" json:"bitcoin,omitempty"`
	Twitter        *TwitterField      `bson:"twitter,omitempty" json:"twitter,omitempty"`
	Facebook       *FacebookField     `bson:"facebook,omitempty" json:"facebook,omitempty"`
	Pdf            *PDFField          `bson:"pdf,omitempty" json:"pdf,omitempty"`
	Mp3            *MP3Field          `bson:"mp3,omitempty" json:"mp3,omitempty"`
	Appstores      *AppStoresField    `bson:"appstores,omitempty" json:"appstores,omitempty"`
	Photos         *PhotosField       `bson:"photos,omitempty" json:"photos,omitempty"`
}

func (q *QRCodeModel) IsFile() bool {
	return q.Type == "pdf" || q.Type == "mp3" || q.Type == "img"
}

type CustomField struct {
	Key   string `bson:"key,omitempty" json:"key,omitempty"`
	Value string `bson:"value,omitempty" json:"value,omitempty"`
}

func (c *CustomField) ToContentQRCode() (string, error) {
	if jsonStr, err := json.Marshal(c); err == nil {
		return string(jsonStr), nil
	} else {
		return "", err
	}
}

type VCardField struct {
	Name            string `bson:"name,omitempty" json:"name,omitempty"`
	LastName        string `bson:"lastName,omitempty" json:"lastName,omitempty"`
	Cellphone       string `bson:"cellphone,omitempty" json:"cellphone,omitempty"`
	Phone           string `bson:"phone,omitempty" json:"phone,omitempty"`
	Fax             string `bson:"fax,omitempty" json:"fax,omitempty"`
	Email           string `bson:"email,omitempty" json:"email,omitempty"`
	CorporationName string `bson:"corporationName,omitempty" json:"corporationName,omitempty"`
	Ocupation       string `bson:"ocupation,omitempty" json:"ocupation,omitempty"`
	Street          string `bson:"street,omitempty" json:"street,omitempty"`
	City            string `bson:"city,omitempty" json:"city,omitempty"`
	PostalCode      string `bson:"postalCode,omitempty" json:"postalCode,omitempty"`
	State           string `bson:"state,omitempty" json:"state,omitempty"`
	Country         string `bson:"country,omitempty" json:"country,omitempty"`
	Website         string `bson:"website,omitempty" json:"website,omitempty"`
}

func (c *VCardField) ToContentQRCode() (string, error) {
	return "IMPLEMENTAR", nil
}

type EmailField struct {
	Email   string `bson:"email,omitempty" json:"email,omitempty"`
	Subject string `bson:"subject,omitempty" json:"subject,omitempty"`
	Message string `bson:"message,omitempty" json:"message,omitempty"`
}

func (c *EmailField) ToContentQRCode() (string, error) {
	return "IMPLEMENTAR", nil
}

type SMSField struct {
	Number  string `bson:"number,omitempty" json:"number,omitempty"`
	Message string `bson:"message,omitempty" json:"message,omitempty"`
}

func (c *SMSField) ToContentQRCode() (string, error) {
	return "IMPLEMENTAR", nil
}

type WIFIField struct {
	NetName  string `bson:"netName,omitempty" json:"netName,omitempty"`
	Visible  bool   `bson:"visible,omitempty" json:"visible,omitempty"`
	Password string `bson:"password,omitempty" json:"password,omitempty"` //WPA ou WEB (radio button)
}

func (c *WIFIField) ToContentQRCode() (string, error) {
	return "IMPLEMENTAR", nil
}

type BitCoinField struct {
	Coin        string `bson:"coin,omitempty" json:"coin,omitempty"` //Bitcoin, Bitcoin Cash, Ether, LiteCoin ou Dash
	Value       string `bson:"value,omitempty" json:"value,omitempty"`
	Destinatary string `bson:"destinatary,omitempty" json:"destinatary,omitempty"`
	Message     string `bson:"message,omitempty" json:"message,omitempty"`
}

func (c *BitCoinField) ToContentQRCode() (string, error) {
	return "IMPLEMENTAR", nil
}

type TwitterField struct {
	Option              string `bson:"option,omitempty" json:"option,omitempty"` //1 vincular ao perfil ou 2 - postar um twitter (radio button)
	Value               string `bson:"value,omitempty" json:"value,omitempty"`
	TextToPost          string `bson:"textToPost,omitempty" json:"textToPost,omitempty"`
	UsernameVinculation string `bson:"usernameVinculation,omitempty" json:"usernameVinculation,omitempty"`
}

func (c *TwitterField) ToContentQRCode() (string, error) {
	return "IMPLEMENTAR", nil
}

type PDFField struct {
	FilePath string `bson:"filePath,omitempty" json:"filePath,omitempty"`
}

func (c *PDFField) ToContentQRCode() (string, error) {
	return "IMPLEMENTAR", nil
}

type MP3Field struct {
	FilePath string `bson:"filePath,omitempty" json:"filePath,omitempty"`
}

func (c *MP3Field) ToContentQRCode() (string, error) {
	return "IMPLEMENTAR", nil
}

type FacebookField struct {
}

func (c *FacebookField) ToContentQRCode() (string, error) {
	return "IMPLEMENTAR", nil
}

type AppStoresField struct {
}

func (c *AppStoresField) ToContentQRCode() (string, error) {
	return "IMPLEMENTAR", nil
}

type PhotosField struct {
}

func (c *PhotosField) ToContentQRCode() (string, error) {
	return "IMPLEMENTAR", nil
}
