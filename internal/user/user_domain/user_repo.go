package user_domain

type UserRepository interface {
	GetUser(id uint64) (*User, error)
	GetUsers(userFilter *UserFilter) ([]User, error)
	// SignIn(email, password string) (*models.User, string, error)                                 //LOGIN
	// SignUp(user models.User, userCredential models.UserCredential) (*models.User, string, error) //REG
	CreateUser(user User) (*User, error)
	UpdateUser(user User) (*User, error)
	CreateUserCredential(credential UserCredential) (*UserCredential, error)
	GetUserCredential(email string) (*UserCredential, error)
	UpdateUserCredential(credential UserCredential) (*UserCredential, error)
	DeleteUser(id uint64) error
}
