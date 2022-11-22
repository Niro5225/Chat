package models

type ChatFilter struct {
	IDs     []uint64
	Search  *string // LIKE name or description
	UserIDs []uint64
}
