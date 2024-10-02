package config

import (
	"os"
)

type ChEnv struct {
	Env string
}

func REnv() *ChEnv {
	env := os.Getenv("KEY")
	if env == "" {
		panic("Не удалось прочитать переменные окружения")
	}
	return &ChEnv{
		Env: env,
	}
}
