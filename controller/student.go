package controller

import (
	"encoding/json"
	"errors"
	"fngc/mailer"
	"fngc/models"
	"net/http"
)


func ExaminationApply(w http.ResponseWriter, r *http.Request) {
	var examData *models.ExamPreparation

	err := decodeJSONBody(w, r, &examData)
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
	} //decode json request into user object

	user, err := examData.ExamApplication()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(models.ValidResponse(http.StatusOK, err.Error(), "error applying for user exam. check body for details "))
		return
	}

	if err = mailer.SendExamPrepMail(user); err != nil {
		models.LogError(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.ValidResponse(http.StatusOK, examData, "success"))
}

func StudentCurriculum(w http.ResponseWriter, r *http.Request) {

}

func NewTutor(w http.ResponseWriter, r *http.Request) {

}
