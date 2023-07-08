package config


var authConfig *AuthConfig

type AuthIdentity string

const (
	USERNAME AuthIdentity = "username"
	EMAIL AuthIdentity = "email"
	PHONE_NUMBER AuthIdentity = "phone_number"
)

type AuthConfig struct {
	Identity AuthIdentity
}

func Auth() *AuthConfig  {
	return authConfig
}

func InitAuthConfig () {
	ac := new(AuthConfig)
	ac.Identity = USERNAME
}
