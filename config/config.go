package config

import "github.com/kataras/iris/config"

// Config Project Config Init
func Config() config.Iris {
	return config.Iris{
		IsDevelopment: true,
		Gzip:          true,
		Charset:       "UTF-8",
	}
}
