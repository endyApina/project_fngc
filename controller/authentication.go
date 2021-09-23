package controller

import (
	"encoding/json"
	"errors"
	"fngc/mailer"
	"fngc/models"
	"net/http"
)

//Registration godoc
//@Summary Handle unique User Registration
//@Description Accept JSON data of Student User objects and returns valid response
//@Accept json
//@produce json
//@Tags Authorization
//@Param   StudentData      body models.UserRegistrationData true  "The Student Registration Data"
//@Success 200 {object} models.UserRegistrationData	"ok"
//@Failure 400 {object} models.ResponseBody "Check Response Message"
//@Router /auth/signup [post]
func Registration(w http.ResponseWriter, r *http.Request) {
	var registrationData models.User

	err := decodeJSONBody(w, r, &registrationData)
	if err != nil {
		var mr *malformedRequest
		if errors.As(err, &mr) {
			models.LogError(err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(models.ValidResponse(http.StatusInternalServerError, "error passing json data. contact support", "error"))
		} else {
			models.LogError(err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(models.ValidResponse(http.StatusInternalServerError, "internal server error", "error"))
		}
	} //decode json request into user object

	if err = registrationData.HandleRegistration(); err != nil {
		models.LogError(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(models.ValidResponse(http.StatusOK, err.Error(), "error registering new user"))
	} //business logic to register a new user

	if err = mailer.SendRegistrationMail(registrationData); err != nil {
		models.LogError(err)
	} //send mail notification

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.ValidResponse(http.StatusOK, registrationData, "success"))
}

//UserLogin godoc
//@Summary Handle unique Unique User Login
//@Description Accept JSON data of User Login objects and returns valid response
//@Accept json
//@produce json
//@Tags Authorization
//@Param   StudentData      body models.LoginData true  "The Tutor Login Data"
//@Success 200 {object} models.LoginData	"ok"
//@Failure 400 {object} models.ResponseBody "Check Response Message"
//@Router /auth/login [post]
func UserLogin(w http.ResponseWriter, r *http.Request) {
	var loginData *models.LoginData
	var loggedIn *models.LoggedInData

	err := decodeJSONBody(w, r, &loginData)
	if err != nil {
		var mr *malformedRequest
		if errors.As(err, &mr) {
			models.LogError(err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(models.ValidResponse(http.StatusInternalServerError, "error passing json data. contact support", "error"))
		} else {
			models.LogError(err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(models.ValidResponse(http.StatusInternalServerError, "internal server error", "error"))
		}
	} //decode json request into user object

	if err = loggedIn.User.Login(loginData); err != nil {
		models.LogError(err)
		if err.Error() == "unverified user" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(models.ValidResponse(http.StatusOK, err.Error(), "error"))
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(models.ValidResponse(http.StatusBadRequest, err.Error(), "error"))
		}
	}

	loggedIn.User.Password = ""

	if err = loggedIn.Token.GenerateTokenString(loginData.Email); err != nil {
		models.LogError(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(models.ValidResponse(http.StatusOK, err.Error(), "error"))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.ValidResponse(http.StatusOK, loggedIn, "success"))
}

//VerifyUser godoc
//@Summary Handle verifying user otp
//@Description Accept JSON data of User Reset password objects and returns valid response
//@Accept json
//@produce json
//@Tags Authorization
//@Param   StudentData      body models.VerifyUser true  "The User Verification Data"
//@Success 200 {object} models.VerifyUser	"ok"
//@Failure 400 {object} models.ResponseBody "Check Response Message"
//@Router /auth/verifyotp [post]
func VerifyUser(w http.ResponseWriter, r *http.Request) {
	var loggedInData *models.LoggedInData
	var otpBody *models.VerifyUser
	err := decodeJSONBody(w, r, &otpBody)
	if err != nil {
		var mr *malformedRequest
		if errors.As(err, &mr) {
			models.LogError(err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(models.ValidResponse(http.StatusInternalServerError, "error passing json data. contact support", "error"))
		} else {
			models.LogError(err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(models.ValidResponse(http.StatusInternalServerError, "internal server error", "error"))
		}
	}

	if err = loggedInData.User.VerifyOTP(otpBody); err != nil {
		models.LogError(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(models.ValidResponse(http.StatusOK, err.Error(), "error"))
	}

	if err = loggedInData.Token.GenerateTokenString(otpBody.Email); err != nil {
		models.LogError(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(models.ValidResponse(http.StatusOK, err.Error(), "error"))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.ValidResponse(http.StatusOK, loggedInData, "success"))

}

//ResetPassword godoc
//@Summary Handle resetting a user password
//@Description Accept JSON data of User Reset password objects and returns valid response
//@Accept json
//@produce json
//@Tags Authorization
//@Param   StudentData      body models.ResetPassword true  "The User Data"
//@Success 200 {object} models.ResetPassword	"ok"
//@Failure 400 {object} models.ResponseBody "Check Response Message"
//@Router /auth/resetpassword [post]
func ResetPassword(w http.ResponseWriter, r *http.Request) {

}
