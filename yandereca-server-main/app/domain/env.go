package domain

type Env struct {
	Value string
}

type Envs []Env

var ERROR_RESPONSE Env = Env{Value: "Error!"}
