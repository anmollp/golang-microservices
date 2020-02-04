package model

import (
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUserNotFound(t *testing.T) {
    //Initialisation
    
    //Execution
	user, err := GetUser(0)

	//Validation
	assert.Nil(t, user, "We were not expecting a user with id 0.")
	assert.NotNil(t, err, "We were expecting an error when user id is 0.")
	assert.EqualValues(t, err.StatusCode, http.StatusNotFound)
	assert.EqualValues(t, err.Code, "Not Found.")
	assert.EqualValues(t, err.Message, "User 0 not found.")
}

func TestGetUserNoError(t *testing.T) {

	user, err := GetUser(123)

	assert.NotNil(t, user)
	assert.Nil(t, err)
	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, user.FirstName, "Anmol")
	assert.EqualValues(t, user.LastName, "Patil")
	assert.EqualValues(t, user.Email, "myemail@gmail.com")

}
