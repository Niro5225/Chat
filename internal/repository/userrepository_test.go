package repository

import (
	"chat-app/models"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	cfg      = New_config()
	db, err  = NewDB(*cfg)
	r        = NewUserR(db)
	TestUser = models.NewUser("test", "test", "testemail")
)

func truncTable() {
	if _, err := db.Exec(fmt.Sprintf("TRUNCATE users CASCADE")); err != nil {
		log.Fatal(err)
	}
}

func TestUserR_CreateUser(t *testing.T) {
	// defer db.Close()

	user, err := r.CreateUser(*TestUser)

	assert.NoError(t, err)
	assert.NotNil(t, user)

	truncTable()
}

func TestUserR_GetUser(t *testing.T) {
	// defer db.Close()

	u, _ := r.CreateUser(*TestUser)

	user, err := r.GetUser(u.ID)

	assert.NoError(t, err)
	assert.NotNil(t, user)

	truncTable()
}

func TestUserR_UpdateUser(t *testing.T) {
	// defer db.Close()

	u, _ := r.CreateUser(*TestUser)

	fmt.Println(u)

	u.FirstName = "firstname"

	fmt.Println(u)

	newUser, err := r.UpdateUser(*u)

	assert.NoError(t, err)
	assert.NotNil(t, newUser)

	truncTable()
}

func TestDeleteUser(t *testing.T) {
	// defer db.Close()
	u, _ := r.CreateUser(*TestUser)

	err := r.DeleteUser(u.ID)

	assert.NoError(t, err)
}

func TestCreateUserCredential(t *testing.T) {
	// defer db.Close()

	u, _ := r.CreateUser(*TestUser)

	userCredential := models.NewUserCredential(u.ID, "testpassword", u.Email)

	userCredential1, err := r.CreateUserCredential(*userCredential)

	fmt.Println(userCredential)

	assert.NoError(t, err)
	assert.NotNil(t, userCredential1)
	// truncTable()

}

func TestUpdateUserCredential(t *testing.T) {
	u, _ := r.CreateUser(*TestUser)

	userCredential := models.NewUserCredential(u.ID, "testpassword", u.Email)

	userCredential, err = r.CreateUserCredential(*userCredential)

	fmt.Println(userCredential)
	truncTable()
}
