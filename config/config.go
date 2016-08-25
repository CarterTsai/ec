package config

import (
	"time"

	"github.com/kataras/iris/config"
)

// Init Project Config Init
func Init() config.Iris {

	return config.Iris{
		IsDevelopment: true,
		Gzip:          true,
		Charset:       "UTF-8",
		Sessions: config.Sessions{
			Cookie:                      "ecsessionid",
			CookieLength:                config.DefaultCookieLength,
			DecodeCookie:                false,
			Expires:                     time.Duration(2) * time.Hour,
			DisableSubdomainPersistence: false,
		},
	}
}
