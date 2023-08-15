package services

import (
	"context"
	"encoding/json"
	"fmt"
	"moflo-be/constants"
	"moflo-be/models"
	"moflo-be/utils"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func UserRegister(response http.ResponseWriter, request *http.Request) {
	var user models.User
	if json.NewDecoder(request.Body).Decode(&user) != nil {
		utils.CreateErrorResponse(response, http.StatusBadRequest, constants.ErrorFieldInvalid)
		return
	}
	if !utils.IsNotNil(user.FullName, user.Password, user.Username, user.Email) {
		utils.CreateErrorResponse(response, http.StatusBadRequest, constants.ErrorFieldEmpty)
		return
	}
	if !utils.IsUsername(user.Username) || !utils.IsPassword(user.Password) {
		utils.CreateErrorResponse(response, http.StatusBadRequest, constants.ErrorFieldFormatInvalid)
		return
	}
	hashedPasswordByte, errBcrypt := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(hashedPasswordByte[:])
	ctx, timeout := context.WithTimeout(context.Background(), constants.TimeOut)
	defer timeout()
	_, errInsert := models.CollectionUser().InsertOne(ctx, user)
	if errInsert != nil || errBcrypt != nil {
		if mongo.IsDuplicateKeyError(errInsert) {
			duplicateKey := utils.ExtractDuplicateKeyError(errInsert.Error())
			utils.CreateErrorResponse(response, http.StatusConflict, fmt.Sprintf("Terdapat %s yang sama", duplicateKey))
			return
		}
		utils.CreateErrorResponse(response, http.StatusInternalServerError, constants.ErrorInternalServerError)
		return
	}
	if ctx.Err() != nil {
		utils.CreateErrorResponse(response, http.StatusRequestTimeout, constants.ErrorTimeout)
		return
	}
	utils.CreateSuccessResponse(response, constants.SuccessRegister, nil)
}

func UserLogin(response http.ResponseWriter, request *http.Request) {
	var user models.User
	var availableUser models.User
	var filter interface{}
	if json.NewDecoder(request.Body).Decode(&user) != nil {
		utils.CreateErrorResponse(response, http.StatusBadRequest, constants.ErrorFieldFormatInvalid)
		return
	}
	if !utils.IsNotNil(user.Username) && utils.IsNotNil(user.Password) {
		filter = bson.M{"username": user.Username}
	} else if !utils.IsNotNil(user.Email) && utils.IsNotNil(user.Password) {
		filter = bson.M{"email": user.Email}
	} else {
		utils.CreateErrorResponse(response, http.StatusBadRequest, constants.ErrorFieldEmpty)
		return
	}
	ctx, timeout := context.WithTimeout(context.Background(), constants.TimeOut)
	defer timeout()
	if models.CollectionUser().FindOne(ctx, filter).Decode(&availableUser) != nil {
		utils.CreateErrorResponse(response, http.StatusInternalServerError, constants.ErrorInternalServerError)
		return
	}
	if !utils.IsNotNil(availableUser.Username) {
		utils.CreateErrorResponse(response, http.StatusNotFound, constants.ErrorUserNotFound)
		return
	}
	if bcrypt.CompareHashAndPassword([]byte(availableUser.Password), []byte(user.Password)) != nil {
		utils.CreateErrorResponse(response, http.StatusForbidden, constants.ErrorWrongPassword)
		return
	}
	var data interface{}
	if json.Unmarshal([]byte(fmt.Sprintf(
		`{"username":"%s","email":"%s","fullName":"%s"}`,
		availableUser.Username,
		availableUser.Email,
		availableUser.FullName)), &data) != nil {
		utils.CreateErrorResponse(response, http.StatusInternalServerError, constants.ErrorInternalServerError)
		return
	}
	utils.CreateSuccessResponse(response, constants.SuccessLogin, data)
}
