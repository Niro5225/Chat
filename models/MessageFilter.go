package models

type MessageFilter struct {
	IDs     []uint64
	Search  *string // LIKE text
	ChatIDs []uint64
	UserIDs []uint64
}
