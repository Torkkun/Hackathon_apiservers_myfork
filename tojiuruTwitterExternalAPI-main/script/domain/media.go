package domain

type Media struct {
	Media_id string `json:"media_id"`
}

type MediaResponse struct {
	MessageID string   `json:"message_id"`
	MediaURL  []string `json:"media_path"`
}
