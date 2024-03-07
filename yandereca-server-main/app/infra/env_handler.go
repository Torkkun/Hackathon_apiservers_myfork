package infra

import (
	"errors"
	"os"

	"yandereca.tech/yandereca/interface/env"
)

type EnvHandler struct {
}

func NewEnvHandler() env.EnvHandler {
	envHandler := new(EnvHandler)
	return envHandler
}

func (handler *EnvHandler) ReadEnv(statement string) (string, error) {
	envs_as_string := os.Getenv(statement)
	if len(envs_as_string) < 1 {
		//エラー処理
		return "No Value", errors.New("Error! Key is wrong!")
	}
	return envs_as_string, nil
}
