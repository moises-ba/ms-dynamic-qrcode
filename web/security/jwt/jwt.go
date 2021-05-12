package jwt

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/moises-ba/ms-dynamic-qrcode/utils"
	"github.com/moises-ba/ms-dynamic-qrcode/web/security/jwt/model"
)

const (
	JWT_PASSWORD_ENV = "JWT_PASSWORD_QRCODE"
	KEY_CLOAK_URI    = "KEY_CLOAK_URI_QRCODE"
)

var jwtConfig model.JwtConfig = model.JwtConfig{
	JWTPassword: os.Getenv(JWT_PASSWORD_ENV),
	KeyCloakURI: utils.GetEnv(KEY_CLOAK_URI, "http://localhost:8080/auth/realms/principal/protocol/openid-connect/certs"),
	ContentType: "application/json",
}

func Config(config model.JwtConfig) {
	jwtConfig = config
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")

	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}

	return ""
}

func ExtractCert() string {
	certKeyCloak := getKeyCloakCerts()
	cert := "-----BEGIN CERTIFICATE-----\n"
	cert += certKeyCloak.Keys[0].X5C[0]
	cert += "\n-----END CERTIFICATE-----"
	return cert
}

/**
valida o token via chave publica do keycloak
**/
func validateKeyCloak(token *jwt.Token) (interface{}, error) {
	rsaPublicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(ExtractCert()))
	if err != nil {
		log.Fatal(err)
	}

	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	return rsaPublicKey, nil
}

func getKeyCloakCerts() *model.KeyCloakCert {

	keyCloakURI := jwtConfig.KeyCloakURI
	if keyCloakURI == "" {
		log.Fatal(errors.New("nao foi possvel obter certificado do keycloak"))
	}

	response, err := http.Get(keyCloakURI)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	keyCert := new(model.KeyCloakCert)

	json.Unmarshal(responseData, keyCert)

	return keyCert
}

func convert(token jwt.Token) *model.PrincipalUserDetail {

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {

		var groups = make([]string, 0, 10)
		if tokenGroups := claims["groups"]; tokenGroups != nil {
			for _, group := range tokenGroups.([]interface{}) {
				groups = append(groups, group.(string))
			}
		}

		var userName, login, uuid string

		if givenName := claims["given_name"]; givenName != nil {
			userName = givenName.(string)
		}

		if preferedUserName := claims["preferred_username"]; preferedUserName != nil {
			login = preferedUserName.(string)
		}

		if sub := claims["sub"]; sub != nil {
			uuid = sub.(string)
		}

		return &model.PrincipalUserDetail{
			UserName: userName,
			Login:    login,
			UUID:     uuid,
			Roles:    groups,
		}
	}

	return nil
}

func ValidateJWTToken(r *http.Request) (*model.PrincipalUserDetail, error) {

	if jwtToken := strings.TrimSpace(ExtractToken(r)); jwtToken != "" {

		funcjWTValidation := validateKeyCloak
		//fazer if caso seja via seha

		token, err := jwt.Parse(jwtToken, funcjWTValidation)

		if err != nil {
			log.Error(err)
			return nil, err
		}

		user := convert(*token)

		return user, nil
	}

	return nil, errors.New("JWT Token nao fornecido")
}

func Authorize(next func(c *gin.Context), roles ...string) func(c *gin.Context) {

	return func(c *gin.Context) {

		if user, err := ValidateJWTToken(c.Request); err == nil {

			if user.HasRole(roles) {
				paramsKeys := map[string]interface{}{
					utils.UserParamName: user,
				}

				c.Keys = paramsKeys
				next(c)
				return
			} else {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": true, "message": "Nao possui perfil de acesso"})
			}

		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": true, "message": err.Error()})
		}

	}

}
