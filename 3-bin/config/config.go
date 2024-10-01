package config

import (
	"os"
)

type chEnv struct {
	Env string
}

func rEnv() *chEnv {
	env := os.Getenv("KEY")
	if env == "" {
		panic("Не удалось прочитать переменные окружения")
	}
	return &chEnv{
		Env: env,
	}
}
