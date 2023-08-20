package config

import "go-jwt/config"

func Get() config.Config {
	return configurations
}
