package controller

import "net/http"

//Registration godoc
//@Summary Handle unique Student User Registration
//@Description Accept JSON data of Student User objects and returns valid response
//@Accept json
//@produce json
//@Tags Authorization
//@Param   StudentData      body models.StudentRegistrationData true  "The Student Registration Data"
//@Success 200 {object} models.StudentRegistrationData	"ok"
//@Failure 400 {object} models.ResponseBody "Check Response Message"
//@Router /auth/signup/student [post]
func StudentRegistration(w http.ResponseWriter, r *http.Request) {

}

//TutorRegistration godoc
//@Summary Handle unique Tutor User Registration
//@Description Accept JSON data of Tutor User objects and returns valid response
//@Accept json
//@produce json
//@Tags Authorization
//@Param   StudentData      body models.TutorRegistrationData true  "The Tutor Registration Data"
//@Success 200 {object} models.TutorRegistrationData	"ok"
//@Failure 400 {object} models.ResponseBody "Check Response Message"
//@Router /auth/signup/tutor [post]
func TutorRegistration(w http.ResponseWriter, r *http.Request) {

}

//UserLogin godoc
//@Summary Handle unique Unique User Registration
//@Description Accept JSON data of User Login objects and returns valid response
//@Accept json
//@produce json
//@Tags Authorization
//@Param   StudentData      body models.LoginData true  "The Tutor Login Data"
//@Success 200 {object} models.LoginData	"ok"
//@Failure 400 {object} models.ResponseBody "Check Response Message"
//@Router /auth/login [post]
func UserLogin(w http.ResponseWriter, r *http.Request) {

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
