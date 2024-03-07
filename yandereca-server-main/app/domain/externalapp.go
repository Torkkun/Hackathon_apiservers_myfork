package domain

type UrlResponse struct {
	AuthUrl string `json:"url"`
}

type PostAuthCodeRequest struct {
	Code string `json:"code"`
}

type CalcProgressResponse struct {
	Progress float64 `json:"progress"`
}
