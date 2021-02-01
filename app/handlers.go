package app

import (
	"HumoStore/models"
	"HumoStore/tokens"
	"encoding/json"
	_ "fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (server *MainServer) CheckUserHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	responseBody := models.ResponseStatus{}
	requestBody := models.RequestLogin{}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Error Invalid json")
		return
	}

	check, err := server.UserSvc.CheckUser(requestBody.Login)
	if err != nil || check == false {
		responseBody.Ok = false
		responseBody.Message = "User not fount. Please check user!"
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(responseBody)
		return
	}
	if check == false {
		responseBody.Ok = false
		responseBody.Message = "User not exists. Please sign up"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(responseBody)
		return
	}
	if check == true {
		responseBody.Ok = true
		responseBody.Message = "User has founded."
		json.NewEncoder(w).Encode(responseBody)
		return
	}

}

func (server *MainServer) SignUp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var requestBody models.User
	var responseBody models.ResponseToken
	var StatusResponse models.StatusResponse
	var Status models.CredentialStatus

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Json is invalid")
		return
	}

	if len(requestBody.Password) < 9 {
		StatusResponse.Ok = false
		StatusResponse.Message = "Passwords must be at least 9 characters."
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(StatusResponse)
		return
	}

	err = server.UserSvc.SignUp(requestBody)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Status)
		return
	}

	userRole := 3
	err = server.UserSvc.UserRole(userRole, requestBody.Login)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userID, responseUser, err := server.UserSvc.GetUserByLogin(requestBody.Login)

	token := tokens.SetToken(userID,responseUser.Login)

	responseBody.Ok = true
	responseBody.Token = token
	responseBody.User = responseUser

	err = json.NewEncoder(w).Encode(responseBody)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}
