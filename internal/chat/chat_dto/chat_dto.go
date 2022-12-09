package chatdto

type ChatDTO struct {
	Id          uint64 `json:"chat_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
