package service

import (
	"github.com/anmollp/golang-microservices/mvc/model"
	"github.com/anmollp/golang-microservices/mvc/utils"
)

func GetUser(userId int64) (*model.User, *utils.ApplicationError) {
	return model.GetUser(userId)
}
