package todo

const (
	userCreateUrl   = "http://localhost:8080/user/create"
	userReadByIdUrl = "http://localhost:8080/user/read?id="
)

type Token struct {
	AccessToken  string
	RefreshToken string
}

type Signed struct {
	Jwt string
}

type Progress struct {
	Result float64
}
