package env

type EnvHandler interface {
	ReadEnv(string) (string, error)
}
