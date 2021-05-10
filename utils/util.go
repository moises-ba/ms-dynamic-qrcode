package utils

import (
	"os"
	"reflect"
	"strings"
)

const (
	UserParamName string = "principal_user"
)

func GetEnv(envName, defaultValue string) string {
	if val := os.Getenv(envName); len(val) > 0 {
		return val
	} else {
		return defaultValue
	}

}

func GetValueByReflection(obj interface{}, field string) interface{} {
	fieldName := strings.Title(strings.ToLower(field))
	valueQRCode := reflect.ValueOf(obj)
	return valueQRCode.Elem().FieldByName(fieldName).Interface()
}

func IsEmpty(txt string) bool {
	return strings.TrimSpace(txt) == ""
}

func NvlString(value, anotherElse string) string {
	if !IsEmpty(value) {
		return value
	} else {
		return anotherElse
	}
}
