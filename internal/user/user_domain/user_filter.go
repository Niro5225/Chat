package user_domain

type UserFilter struct {
	IDs    []uint64
	Email  *string
	Search *string // LIKE first_name or last_name
	Limit  *int
	Offset *int
}
