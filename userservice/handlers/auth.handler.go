package handlers

import (
	_ "crypto/rand"
	"encoding/json"
	"github.com/gookit/validate"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
	"userservice/dto"
	"userservice/helper"
	"userservice/services"
)

type AuthHandler struct {
	UserService         services.UserService
	JwtService          services.JwtService
	NotificationService services.NotificationService
}

func (au AuthHandler) ManualLogin(w http.ResponseWriter, r *http.Request) {
	loginRequest := dto.LoginDTO{}

	json.NewDecoder(r.Body).Decode(&loginRequest)

	v := validate.Struct(loginRequest)
	if !v.Validate() {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(v.Errors)
		return
	}

	findUserEmail := au.UserService.FindUserByEmail(loginRequest.Email)

	if findUserEmail == nil {
		logrus.Println("i made it here")
		message := map[string]string{"error": "Invalid Credentials"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&message)
		return
	}
	user := findUserEmail

	checkPassword := user.CheckPasswordHash(loginRequest.Password)
	if checkPassword != nil {
		message := map[string]string{"error": "Invalid Credentials"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message)
		return
	}
	logrus.Println(string(user.Role))
	generateToken, err := au.JwtService.GenerateTokenPair(int(user.ID), string(user.Role))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logrus.Println(err.Error())
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	jsonb, err := json.Marshal(generateToken)
	if err != nil {
		message := map[string]string{"error": "An Error Occurred"}
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Println(err)
		json.NewEncoder(w).Encode(message)
		return
	}
	autoresponder := dto.AuthResponse{}
	if err := json.Unmarshal(jsonb, &autoresponder); err != nil {
		message := map[string]string{"error": "An Error Occurred"}
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Println(err)
		json.NewEncoder(w).Encode(message)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(autoresponder)

}

func GoogleLogin(w http.ResponseWriter, r *http.Request) {

}

func FacebookLogin(w http.ResponseWriter, r *http.Request) {

}

func (au AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {

	tokenRequest := dto.RefreshTokenDTO{}

	json.NewDecoder(r.Body).Decode(&tokenRequest)

	v := validate.Struct(tokenRequest)
	if !v.Validate() {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(v.Errors)
		return
	}

	sub, error := au.JwtService.ValidateToken(tokenRequest.RefreshToken)

	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
	}
	payload, ok := sub.(map[string]interface{})

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		message := map[string]string{"error": "could not parse value"}
		json.NewEncoder(w).Encode(message)
		return

	}
	check := payload["sub"]
	result, ok := check.(dto.ValidateRefreshToken)

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		message := map[string]string{"error": "could not parse value"}
		json.NewEncoder(w).Encode(message)
		return

	}

	if result.IsRefresh != true {
		w.WriteHeader(http.StatusBadRequest)
		message := map[string]string{"error": "require refresh token but received access token"}
		json.NewEncoder(w).Encode(message)
		return

	}
	user := au.UserService.FindUserById(result.UserId)

	if user == nil {
		message := map[string]string{"error": "invalid token"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&message)
		return
	}

	generateToken, err := au.JwtService.GenerateTokenPair(int(user.ID), string(user.Role))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Println(err.Error())
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	jsonb, err := json.Marshal(generateToken)
	if err != nil {
		message := map[string]string{"error": "An Error Occurred"}
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Println(err)
		json.NewEncoder(w).Encode(message)
		return
	}
	autoresponder := dto.AuthResponse{}
	if er := json.Unmarshal(jsonb, &autoresponder); err != nil {
		message := map[string]string{"error": "An Error Occurred"}
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Println(er)
		json.NewEncoder(w).Encode(message)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(autoresponder)

}

func (au AuthHandler) VerifyOTPAndAuthenticate(w http.ResponseWriter, r *http.Request) {
	newRequest := dto.VerifyOTPDTO{}

	json.NewDecoder(r.Body).Decode(&newRequest)

	v := validate.Struct(newRequest)
	if !v.Validate() {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(v.Errors)
		return
	}

	findOTPCode := au.UserService.FindUserOTPCode(newRequest.OTPCode)

	if findOTPCode != nil {
		message := map[string]string{"error": "Invalid OTP code"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&message)
		return
	}

	if findOTPCode.OTPExpireTime.Unix() < time.Now().Local().Unix() {
		message := map[string]string{"error": "OTP Code Expired"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&message)
		return
	}

	generateToken, err := au.JwtService.GenerateTokenPair(int(findOTPCode.ID), string(findOTPCode.Role))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logrus.Println(err.Error())
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	jsonb, err := json.Marshal(generateToken)
	if err != nil {
		// do error check
		logrus.Println(err)
		return
	}
	autoresponder := dto.AuthResponse{}
	if err := json.Unmarshal(jsonb, &autoresponder); err != nil {
		// do error check
		logrus.Println(err)
		return
	}
	json.NewEncoder(w).Encode(autoresponder)

}

func (au AuthHandler) ResendOTP(w http.ResponseWriter, r *http.Request) {
	newRequest := dto.ResendOTP{}

	json.NewDecoder(r.Body).Decode(&newRequest)

	user := au.UserService.FindUserByEmail(newRequest.Email)

	if user == nil {
		message := map[string]string{"error": "Email not Recognized"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&message)
		return
	}

	str := helper.GenerateOTP()
	emailBody := helper.GenerateOTPEmailTemplate(str)
	subject := "OTP Verification"
	defer au.NotificationService.SendEmail(user.Email, emailBody, subject)
	expirationTime := time.Now().Add(8 * time.Minute).Unix()

	data := map[string]interface{}{"otp_code": str, "otp_expire_time": expirationTime}

	result := au.UserService.UpdateBasicDetails(int(user.ID), data)

	if result != nil {
		message := map[string]string{"message": "OTP Re-sent"}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&message)
	} else {
		message := map[string]string{"error": "Could not Resend OTP"}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&message)
		return
	}

}

func (au AuthHandler) VerifyOTP(w http.ResponseWriter, r *http.Request) {
	newRequest := dto.VerifyOTPDTO{}

	json.NewDecoder(r.Body).Decode(&newRequest)

	v := validate.Struct(newRequest)
	if !v.Validate() {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(v.Errors)
		return
	}

	findOTPCode := au.UserService.FindUserOTPCode(newRequest.OTPCode)

	if findOTPCode != nil {
		message := map[string]string{"error": "Invalid OTP code"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&message)
		return
	}

	if findOTPCode.OTPExpireTime.Unix() < time.Now().Local().Unix() {
		message := map[string]string{"error": "OTP Code Expired"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&message)
		return
	}

	message := map[string]string{"message": "OTP Code Verified"}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&message)

}

func (au AuthHandler) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	newRequest := dto.ForgotPasswordDTO{}
	json.NewDecoder(r.Body).Decode(&newRequest)

	v := validate.Struct(newRequest)
	if !v.Validate() {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(v.Errors)
		return
	}

	findUserEmail := au.UserService.FindUserByEmail(newRequest.Email)

	if findUserEmail == nil {
		logrus.Println("i made it here")
		message := map[string]string{"error": "user with this email not found"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&message)
		return
	}

	str := helper.GenerateOTP()
	emailBody := helper.GenerateForgetPasswordEmailTemplate(str)
	subject := "Forgot Password"
	defer au.NotificationService.SendEmail(findUserEmail.Email, emailBody, subject)
	expirationTime := time.Now().Add(8 * time.Minute).Unix()

	data := map[string]interface{}{"otp_code": str, "otp_expire_time": expirationTime}

	result := au.UserService.UpdateBasicDetails(int(findUserEmail.ID), data)

	if result != nil {
		message := map[string]string{"message": "Reset Password initiated"}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&message)
	} else {
		message := map[string]string{"error": "Could not initiate Reset Password"}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&message)
		return
	}

}

func (au AuthHandler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	newRequest := dto.ResetPasswordDTO{}

	json.NewDecoder(r.Body).Decode(&newRequest)

	v := validate.Struct(newRequest)
	if !v.Validate() {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(v.Errors)
		return
	}

	findOTPCode := au.UserService.FindUserOTPCode(newRequest.OTPCode)

	if findOTPCode != nil {
		message := map[string]string{"error": "Invalid OTP code"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&message)
		return
	}

	if findOTPCode.OTPExpireTime.Unix() < time.Now().Local().Unix() {
		message := map[string]string{"error": "OTP Code Expired"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&message)
		return
	}
	data := map[string]interface{}{"otp_code": "", "otp_expire_time": 0, "password": newRequest.NewPasssword}

	result := au.UserService.UpdateBasicDetails(int(findOTPCode.ID), data)

	if result != nil {
		message := map[string]string{"message": "Reset Password initiated"}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&message)
	} else {
		message := map[string]string{"error": "Could not Reset Password"}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&message)
		return
	}

}

func (au AuthHandler) ManualRegistration(w http.ResponseWriter, r *http.Request) {
	newUserRequest := dto.CreateNewUserDTO{}
	json.NewDecoder(r.Body).Decode(&newUserRequest)

	v := validate.Struct(newUserRequest)
	if !v.Validate() {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(v.Errors)
		return
	}

	findUserEmail := au.UserService.FindUserByEmail(newUserRequest.Email)

	if findUserEmail != nil {
		logrus.Println("i made it here")
		message := map[string]string{"error": "user with this email exist already"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&message)
		return
	}

	str := helper.GenerateOTP()
	expirationTime := time.Now().Add(10 * time.Minute)

	newUserRequest.OTPCode = str
	newUserRequest.OTPExpireTime = expirationTime
	newUser := au.UserService.CreateUser(newUserRequest)

	if newUser != nil {
		emailBody := helper.GenerateOTPEmailTemplate(str)
		subject := "Very Email"
		defer au.NotificationService.SendEmail(newUser.Email, emailBody, subject)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(&newUser)

	} else {
		message := map[string]string{"error": "Failed to register user"}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&message)
	}

}
