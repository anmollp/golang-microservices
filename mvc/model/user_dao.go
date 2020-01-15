package model

import (
	"github.com/anmollp/golang-microservices/mvc/utils"
	"net/http"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Anmol", LastName: "Patil", Email: "myemail@gmail.com"},
	}
)
func GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{Message: "User not found.", StatusCode:http.StatusNotFound, Code: "Not Found."}
}
