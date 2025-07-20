package controllers

import (
	utils "github.com/nchikkam/context-findr-be/utils/infrastructure"
	"go.mongodb.org/mongo-driver/mongo"
)

func getUserCollection() *mongo.Collection {
	return utils.DataBase.Collection("users")
}

func getFileCollection() *mongo.Collection {
	return utils.DataBase.Collection("files")
}
