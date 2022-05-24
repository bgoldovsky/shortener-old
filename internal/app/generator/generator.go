package generator

import (
	"math/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type generator struct{}

func NewGenerator() *generator {
	return &generator{}
}

func (g *generator) Shortcut() string {
	return generate(5)
	//return fmt.Sprintf("%s.ets", generate(10))
}

func generate(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
