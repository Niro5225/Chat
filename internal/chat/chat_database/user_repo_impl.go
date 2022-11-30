package chat_database

import (
	"chat-app/internal/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
)

const tokenKey = "ndkasd#nasjnda#kndkj"

type UserRepoImpl struct {
	db *sqlx.DB
}

func NewUserRepoImpl(db *sqlx.DB) *UserRepoImpl {
	return &UserRepoImpl{db: db}
}

func (r *UserRepoImpl) CreateUser(user models.User) (*models.User, error) {
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

func (r *UserRepoImpl) GetUser(id uint64) (*models.User, error) {
	var u models.User
	if err := r.db.QueryRow(
		"SELECT first_name,last_name,email FROM users WHERE id = $1", id,
	).Scan(&u.FirstName, &u.LastName, &u.Email); err != nil {
		return nil, err
	}

	u.ID = id

	return &u, nil
}

func (r *UserRepoImpl) GetUsers(userFilter *models.UserFilter) ([]models.User, error) {
	users := []models.User{}
	if userFilter != nil {
		if len(userFilter.IDs) != 0 {
			for _, id := range userFilter.IDs {
				var u models.User
				if err := r.db.QueryRow(
					"SELECT first_name,last_name,email FROM users WHERE id = $1", id,
				).Scan(&u.FirstName, &u.LastName, &u.Email); err != nil {
					return nil, err
				}
				users = append(users, u)
			}
		} else if userFilter.Email != nil {
			var u models.User
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
				var u models.User
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
			var u models.User
			err = rows.StructScan(&u)
			users = append(users, u)
		}
	}
	return users, nil
}

type NewTokenClaims struct {
	jwt.StandardClaims
	UserId uint64
}

func GenerateToken(userId uint64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, NewTokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix()},
		userId})

	return token.SignedString([]byte(tokenKey))
}

func (r *UserRepoImpl) SignIn(email, password string) (*models.User, string, error) {
	uc, err := r.GetUserCredential(email)
	if err != nil {
		return nil, "", err
	}

	err = uc.CheckPasswords(password)
	if err != nil {
		return nil, "", err
	}
	user, err := r.GetUser(uc.ID)
	if err != nil {
		return nil, "", err
	}
	token, err := GenerateToken(user.ID)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (r *UserRepoImpl) SignUp(user models.User, userCredential models.UserCredential) (*models.User, string, error) {
	// u, err := r.CreateUser(user)
	// if err != nil {
	// 	return nil, "", err
	// }
	// _, err = r.CreateUserCredential(userCredential)
	// if err != nil {
	// 	return nil, "", err
	// }
	token, err := GenerateToken(user.ID)
	if err != nil {
		return nil, "", err
	}

	return &user, token, nil
}

func (r *UserRepoImpl) UpdateUser(user models.User) (*models.User, error) {
	row := r.db.QueryRow("UPDATE users SET first_name = $2, last_name = $3, email=$4,updated_at=$5 WHERE id = $1",
		user.ID, user.FirstName, user.LastName, user.Email, time.Now())
	if row.Err() != nil {
		return nil, row.Err()
	}
	return &user, nil
}

func (r *UserRepoImpl) CreateUserCredential(credential models.UserCredential) (*models.UserCredential, error) {
	err := r.db.QueryRow("INSERT INTO user_credential (id,email,password) VALUES ($1, $2,$3) RETURNING id",
		credential.ID, credential.Email, credential.Password).Scan(&credential.ID)

	if err != nil {
		return nil, err
	}

	return &credential, nil
}

func (r *UserRepoImpl) GetUserCredential(email string) (*models.UserCredential, error) {
	var uc models.UserCredential
	if err := r.db.QueryRow(
		"SELECT id,email,password FROM user_credential WHERE email = $1", email,
	).Scan(&uc.ID, &uc.Email, &uc.Password); err != nil {
		return nil, err
	}

	return &uc, nil
}

func (r *UserRepoImpl) UpdateUserCredential(credential models.UserCredential) (*models.UserCredential, error) {
	row := r.db.QueryRow("UPDATE user_credential SET email=$2,password=$3 WHERE id = $1",
		credential.ID, credential.Email, credential.Password)

	if row.Err() != nil {
		return nil, row.Err()
	}

	return &credential, nil
}
