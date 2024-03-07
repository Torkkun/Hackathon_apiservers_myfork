package domain

type PostExamination struct {
	Message  string   `json:"message"`
	Deadline string   `json:"deadline"`
	People   int      `json:"people"`
	MediaID  []string `json:"media_id,omitempty"`
}
type Examination struct {
	Message_id string `json:"message_id"`
	Message    string `json:"message"`
	People     int    `json:"people"`
	Good_num   int    `json:"good_num"`
	Bad_num    int    `json:"bad_num"`
	CreatedAt  string `json:"created_at"`
	Deadline   string `json:"deadline,omitempty"`
	UserId     string `json:"user_id"`
	Username   string `json:"username"`
	State      int    `json:"state"`
}

type ResponseExaminations struct {
	Examinations []Examination
	//Media        []MediaResponse
}
