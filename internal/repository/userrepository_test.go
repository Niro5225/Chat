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

func truncTable(table string) {
	if _, err := db.Exec(fmt.Sprintf(fmt.Sprintf("TRUNCATE %s CASCADE", table))); err != nil {
		log.Fatal(err)
	}
}

func TestUserR_CreateUser(t *testing.T) {
	// defer db.Close()

	user, err := r.CreateUser(*TestUser)

	assert.NoError(t, err)
	assert.NotNil(t, user)

	truncTable("users")
}

func TestUserR_GetUser(t *testing.T) {

	u, _ := r.CreateUser(*TestUser)

	user, err := r.GetUser(u.ID)

	assert.NoError(t, err)
	assert.NotNil(t, user)

	truncTable("users")
}

func TestUserR_UpdateUser(t *testing.T) {

	u, _ := r.CreateUser(*TestUser)

	u.FirstName = "firstname"

	newUser, err := r.UpdateUser(*u)

	assert.NoError(t, err)
	assert.NotNil(t, newUser)

	truncTable("users")
}

func TestDeleteUser(t *testing.T) {
	u, _ := r.CreateUser(*TestUser)

	err := r.DeleteUser(u.ID)

	assert.NoError(t, err)
}

func TestCreateUserCredential(t *testing.T) {

	u, _ := r.CreateUser(*TestUser)

	userCredential := models.NewUserCredential(u.ID, "testpassword", u.Email)

	userCredential1, err := r.CreateUserCredential(*userCredential)

	assert.NoError(t, err)
	assert.NotNil(t, userCredential1)
	truncTable("users")

}

func TestUpdateUserCredential(t *testing.T) {
	u, _ := r.CreateUser(*TestUser)

	userCredential := models.NewUserCredential(u.ID, "testpassword", u.Email)

	userCredential, err = r.CreateUserCredential(*userCredential)

	newUser, err := r.GetUser(u.ID)
	if err != nil {
		log.Fatal(err)
	}

	newUser.Email = "1111111"

	newUser, err = r.UpdateUser(*newUser)
	if err != nil {
		log.Fatal(err)
	}

	newCred := models.NewUserCredential(newUser.ID, "testpassword", newUser.Email)
	newCred.Encryption_password()
	newCred, err = r.UpdateUserCredential(*newCred)

	assert.NoError(t, err)
	assert.NotNil(t, newCred)

	truncTable("users")
}

func TestGetUsersByIDs(t *testing.T) {
	u, _ := r.CreateUser(*TestUser)
	u1, _ := r.CreateUser(*models.NewUser("second", "second", "second"))
	u2, _ := r.CreateUser(*models.NewUser("third", "third", "third"))
	filter := models.UserFilter{IDs: []uint64{u.ID, u1.ID, u2.ID}}
	users, err := r.GetUsers(&filter)

	assert.NoError(t, err)
	assert.NotNil(t, users)
	truncTable("users")
}

func TestGetUsersByEmail(t *testing.T) {
	u, _ := r.CreateUser(*TestUser)
	filter := models.UserFilter{Email: &u.Email}
	users, err := r.GetUsers(&filter)

	assert.NoError(t, err)
	assert.NotNil(t, users)

	truncTable("users")
}

func TestGetUsersBySearch(t *testing.T) {
	u, _ := r.CreateUser(*TestUser)
	filter := models.UserFilter{Search: &u.FirstName}
	users, err := r.GetUsers(&filter)

	assert.NoError(t, err)
	assert.NotNil(t, users)

	truncTable("users")
}
