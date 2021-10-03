package controller

import (
	"encoding/json"
	"errors"
	"fngc/mailer"
	"fngc/models"
	"net/http"
)

//ExaminationApply godoc
//@Summary Handle apply for student examination
//@Description Accept JSON data for examination objects and returns valid response
//@Accept json
//@produce json
//@Tags Tutor
//@Success 200 {object} models.TutorDashboard	"ok"
//@Failure 400 {object} models.ResponseBody "Check Response Message"
//@Router /tutor/dashboard/ [get]
func TutorDashboard(w http.ResponseWriter, r *http.Request) {

}

//TotorEducationSetting godoc
//@Summary Handle setting up tutor educational status
//@Description Accept JSON data of User educational registration objects and returns valid response
//@Accept json
//@produce json
//@Tags Authorization
//@Param   TutorEducationalData      body models.TutorEducationalData true  "The Tutor Educational Data"
//@Success 200 {object} models.TutorEducationalData	"ok"
//@Failure 400 {object} models.ResponseBody "Check Response Message"
//@Router /tutor/education/ [post]
func TutorEducationSetting(w http.ResponseWriter, r *http.Request) {
	var edData models.TutorEducationalData
	err := decodeJSONBody(w, r, &edData)
	if err != nil {
		var mr *malformedRequest
		if errors.As(err, &mr) {
			models.LogError(err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(models.ValidResponse(http.StatusInternalServerError, "error passing json data. contact support", "error"))
			return
		} else {
			models.LogError(err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(models.ValidResponse(http.StatusInternalServerError, "internal server error", "error"))
			return
		}
	}


	if edData.TutorID == "" {
		errMessage := errors.New("invalid tutor id")
		models.LogError(errMessage)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(models.ValidResponse(http.StatusOK, errMessage, "error"))
		return
	}

	err, tutorUserData := edData.HandleTutorEducation()
	if err != nil {
		models.LogError(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(models.ValidResponse(http.StatusOK, err, "error"))
		return
	}

	err = mailer.SendNewTutorMail(tutorUserData)
	if err != nil {
		models.LogError(err)
	} //send mail notification

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.ValidResponse(http.StatusOK, edData, "success"))

}
