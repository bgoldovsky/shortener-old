package config

import "os"

func GetPort() string {
	p := os.Getenv("PORT")
	if p == "" {
		p = "8080"
	}

	return ":" + p
}

func GetShortcutHost() string {
	h := os.Getenv("SHORTCUT_HOST")
	if h == "" {
		h = "http://localhost:8080"
	}

	return h
}
