package domain

type UserData struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	GoogleUid    string `json:"google_uid"`
}

type UserDataSuccessMessage struct {
	Message string `json:"message"`
	Result  bool   `json:"result"`
	UserId  string `json:"id"`
}

type UserDataList []UserData
