package model

type PrincipalUserDetail struct {
	UserName string
	Login    string
	Roles    []string
	UUID     string
}

func (user *PrincipalUserDetail) HasRole(roles []string) bool {

	if len(roles) == 0 {

		return true
	}

	if user != nil {
		for _, current := range user.Roles {
			for _, currentParam := range roles {
				if currentParam == current {
					return true
				}
			}
		}
	}

	return false
}

type KeyCloakCert struct {
	Keys []struct {
		Kid     string   `json:"kid"`
		Kty     string   `json:"kty"`
		Alg     string   `json:"alg"`
		Use     string   `json:"use"`
		N       string   `json:"n"`
		E       string   `json:"e"`
		X5C     []string `json:"x5c"`
		X5T     string   `json:"x5t"`
		X5TS256 string   `json:"x5t#S256"`
	} `json:"keys"`
}

type JwtConfig struct {
	KeyCloakURI string
	JWTPassword string
	ContentType string
}
