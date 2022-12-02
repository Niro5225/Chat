package chat_database

import (
	"chat-app/internal/chat/chat_domain"
	"time"

	"github.com/jmoiron/sqlx"
)

type UserRepoImpl struct {
	db *sqlx.DB
}

func NewUserRepoImpl(db *sqlx.DB) *UserRepoImpl {
	return &UserRepoImpl{db: db}
}

func (r *UserRepoImpl) CreateUser(user chat_domain.User) (*chat_domain.User, error) {
	err := r.db.QueryRow(
		"INSERT INTO users (first_name,last_name,email,created_at) VALUES ($1, $2,$3,$4) RETURNING id", user.FirstName, user.LastName, user.Email, user.CreatedAt,
	).Scan(&user.ID)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepoImpl) DeleteUser(id uint64) error {
	row := r.db.QueryRow("DELETE FROM users WHERE id = $1", id)
	if row.Err() != nil {
		return row.Err()
	}
	return nil
}

func (r *UserRepoImpl) GetUser(id uint64) (*chat_domain.User, error) {
	var u chat_domain.User
	if err := r.db.QueryRow(
		"SELECT first_name,last_name,email FROM users WHERE id = $1", id,
	).Scan(&u.FirstName, &u.LastName, &u.Email); err != nil {
		return nil, err
	}

	u.ID = id

	return &u, nil
}

func (r *UserRepoImpl) GetUsers(userFilter *chat_domain.UserFilter) ([]chat_domain.User, error) {
	users := []chat_domain.User{}
	if userFilter != nil {
		if len(userFilter.IDs) != 0 {
			for _, id := range userFilter.IDs {
				var u chat_domain.User
				if err := r.db.QueryRow(
					"SELECT first_name,last_name,email FROM users WHERE id = $1", id,
				).Scan(&u.FirstName, &u.LastName, &u.Email); err != nil {
					return nil, err
				}
				users = append(users, u)
			}
		} else if userFilter.Email != nil {
			var u chat_domain.User
			if err := r.db.QueryRow(
				"SELECT id,first_name,last_name,email FROM users WHERE email = $1", userFilter.Email,
			).Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email); err != nil {
				return nil, err
			}
		} else if userFilter.Search != nil {

			rows, err := r.db.Queryx("SELECT * FROM users WHERE first_name=$1 OR last_name=$1", userFilter.Search)
			if err != nil {
				return nil, err
			}
			for rows.Next() {
				var u chat_domain.User
				err = rows.StructScan(&u)
				users = append(users, u)
			}
		}
	} else {

		rows, err := r.db.Queryx("SELECT * FROM users")
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			var u chat_domain.User
			err = rows.StructScan(&u)
			users = append(users, u)
		}
	}
	return users, nil
}

// func (r *UserRepoImpl) SignIn(email, password string) (*chat_domain.User, string, error) {
// 	uc, err := r.GetUserCredential(email)
// 	if err != nil {
// 		return nil, "", err
// 	}

// 	err = uc.CheckPasswords(password)
// 	if err != nil {
// 		return nil, "", err
// 	}
// 	user, err := r.GetUser(uc.ID)
// 	if err != nil {
// 		return nil, "", err
// 	}
// 	token, err := GenerateToken(user.ID)
// 	if err != nil {
// 		return nil, "", err
// 	}

// 	return user, token, nil
// }

// func (r *UserRepoImpl) SignUp(user chat_domain.User, userCredential chat_domain.UserCredential) (*chat_domain.User, string, error) {
// 	// u, err := r.CreateUser(user)
// 	// if err != nil {
// 	// 	return nil, "", err
// 	// }
// 	// _, err = r.CreateUserCredential(userCredential)
// 	// if err != nil {
// 	// 	return nil, "", err
// 	// }
// 	token, err := GenerateToken(user.ID)
// 	if err != nil {
// 		return nil, "", err
// 	}

// 	return &user, token, nil
// }

func (r *UserRepoImpl) UpdateUser(user chat_domain.User) (*chat_domain.User, error) {
	row := r.db.QueryRow("UPDATE users SET first_name = $2, last_name = $3, email=$4,updated_at=$5 WHERE id = $1",
		user.ID, user.FirstName, user.LastName, user.Email, time.Now())
	if row.Err() != nil {
		return nil, row.Err()
	}
	return &user, nil
}

func (r *UserRepoImpl) CreateUserCredential(credential chat_domain.UserCredential) (*chat_domain.UserCredential, error) {
	err := r.db.QueryRow("INSERT INTO user_credential (id,email,password) VALUES ($1, $2,$3) RETURNING id",
		credential.ID, credential.Email, credential.Password).Scan(&credential.ID)

	if err != nil {
		return nil, err
	}

	return &credential, nil
}

func (r *UserRepoImpl) GetUserCredential(email string) (*chat_domain.UserCredential, error) {
	var uc chat_domain.UserCredential
	if err := r.db.QueryRow(
		"SELECT id,email,password FROM user_credential WHERE email = $1", email,
	).Scan(&uc.ID, &uc.Email, &uc.Password); err != nil {
		return nil, err
	}

	return &uc, nil
}

func (r *UserRepoImpl) UpdateUserCredential(credential chat_domain.UserCredential) (*chat_domain.UserCredential, error) {
	row := r.db.QueryRow("UPDATE user_credential SET email=$2,password=$3 WHERE id = $1",
		credential.ID, credential.Email, credential.Password)

	if row.Err() != nil {
		return nil, row.Err()
	}

	return &credential, nil
}
