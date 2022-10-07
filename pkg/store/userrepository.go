package store

import "chat/models"

type User_repository struct {
	store *Store
}

func (r *User_repository) Create(u *models.User) (*models.User, error) {

	if err := u.Validate(); err != nil {
		return nil, err
	}

	if err := u.Before_create(); err != nil {
		return nil, err
	}

	err := r.store.db.QueryRow(
		"INSERT INTO users (username,password_hash) VALUES ($1, $2) RETURNING id", u.Username, u.Encrypted_Password,
	).Scan(&u.Id)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *User_repository) Find_by_username(username string) (*models.User, error) {
	u := &models.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, username, password_hash FROM users WHERE username = $1", username,
	).Scan(&u.Id, &u.Username, &u.Encrypted_Password); err != nil {
		return nil, err
	}

	return u, nil
}
