package config

import (
	"time"

	"github.com/kataras/iris"
)

// Init Project Config Init
func Init() iris.Configuration {
	iris.StaticCacheDuration = time.Duration(0)

	return iris.Configuration{
		IsDevelopment: true,
		Gzip:          true,
		Charset:       "UTF-8",
		Sessions: iris.SessionsConfiguration{
			Cookie:                      "ecsessionid",
			CookieLength:                iris.DefaultCookieLength,
			DecodeCookie:                false,
			Expires:                     time.Duration(2) * time.Hour,
			DisableSubdomainPersistence: false,
		},
	}
}
