package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gookit/validate"
	"github.com/sirupsen/logrus"
	"net/http"
	"userservice/dto"
	"userservice/services"
)

func ManualLogin(w http.ResponseWriter, r *http.Request) {
	loginRequest := dto.LoginDTO{}

	json.NewDecoder(r.Body).Decode(&loginRequest)

	v := validate.Struct(loginRequest)
	if v.Validate(); v.Errors != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(v.Errors)
		return
	}

	fmt.Print("i made it here")
	findUserEmail := services.UserService{}.FindUserByEmail(loginRequest.Email)

	if findUserEmail == nil {
		logrus.Println("i made it here")
		message := map[string]string{"error": "Invalid Email"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message)
		return
	}
	user := *findUserEmail

	checkPassword := user.CheckPasswordHash(loginRequest.Password)
	if checkPassword != nil {
		message := map[string]string{"error": "Invalid Credentials"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message)
		return
	}
	generateToken, err := services.GenerateJWT(int(user.ID))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logrus.Println(err.Error())
		json.NewEncoder(w).Encode(err)
		return
	}
	authresponse := dto.AuthResponse{
		AccessToken: generateToken,
	}

	json.NewEncoder(w).Encode(authresponse)

}

func GoogleLogin(w http.ResponseWriter, r *http.Request) {

}

func FacebookLogin(w http.ResponseWriter, r *http.Request) {

}

func RefreshToken(w http.ResponseWriter, r *http.Request) {

}

func Me(w http.ResponseWriter, r *http.Request) {
	// pass token to get user

}

func VerifyOTP(w http.ResponseWriter, r *http.Request) {

}

func ResendOTP(w http.ResponseWriter, r *http.Request) {

}

func ForgotPassword(w http.ResponseWriter, r *http.Request) {

}

func ResetPassword(w http.ResponseWriter, r *http.Request) {

}

func ManualRegistration(w http.ResponseWriter, r *http.Request) {
	newUserRequest := dto.CreateNewUserDTO{}
	json.NewDecoder(r.Body).Decode(&newUserRequest)

	v := validate.Struct(newUserRequest)
	if v.Validate(); v.Errors != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(v.Errors)
		return
	}

	newUser := services.UserService{}.CreateUser(newUserRequest)

	if newUser != nil {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(&newUser)

	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

}
