package userdto

type UserDTO struct {
	Header       string `json:"header"`
	Boby         string `json:"body"`
	ShortBody    string `json:"short_body"`
	CategoryUUID string `json:"category_uuid"`
	Tags         []int  `json:"tags"`
}
