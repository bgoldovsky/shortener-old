package config

import "os"

func GetPort() string {
	p := os.Getenv("PORT")
	if p == "" {
		p = "8080"
	}

	return ":" + p
}
