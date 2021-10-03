package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"fngc/mailer"
	"fngc/models"
	"io"
	"net/http"
	"strings"

	"github.com/golang/gddo/httputil/header"
)

func FNGCInit(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to FNGC backend api router")
}

//AddTask godoc
//@Summary Handle adding task for users
//@Description Accept JSON data for examination objects and returns valid response
//@Accept json
//@produce json
//@Tags Hybrid APIs
//@Param   TaskData      body models.Task true  "The Task Data"
//@Success 200 {object} models.Task	"ok"
//@Failure 400 {object} models.ResponseBody "Check Response Message"
//@Router /user/task/ [post]
func AddTask(w http.ResponseWriter, r *http.Request) {

}

//StudyAbroad godoc
//@Summary Handle data to process abroad study
//@Description Accept JSON data for study details objects and returns valid response
//@Accept json
//@produce json
//@Tags Hybrid APIs
//@Param   StudyAbroadData      body models.StudyAbroad true  "The Task Data"
//@Success 200 {object} models.StudyAbroad	"ok"
//@Failure 400 {object} models.ResponseBody "Check Response Message"
//@Router /user/studyabroad/ [post]
func StudyAbroad(w http.ResponseWriter, r *http.Request) {
	var studData models.StudyAbroad

	err := decodeJSONBody(w, r, &studData)
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

	if err = studData.CreateStudyAbroad(); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(models.ValidResponse(http.StatusInternalServerError, err.Error(), "error"))
	}

	if err = mailer.SendStudyAbroad(studData); err != nil {
		models.LogError(err)
	} //send mail notification

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.ValidResponse(http.StatusOK, studData, "success"))
}

//ContactUs godoc
//@Summary Handle contact us request
//@Description Accept JSON data for contact details and message objects then returns valid response
//@Accept json
//@produce json
//@Tags Hybrid APIs
//@Param   ContactUs      body models.ContactUs true  "The Task Data"
//@Success 200 {object} models.ContactUs	"ok"
//@Failure 400 {object} models.ResponseBody "Check Response Message"
//@Router /user/contactus/ [post]
func ContactUs(w http.ResponseWriter, r *http.Request) {
	var contact models.ContactUs

	err := decodeJSONBody(w, r, &contact)
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

	if err = contact.HandleContactUs(); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(models.ValidResponse(http.StatusInternalServerError, err.Error(), "error"))
		return
	}

	if err = mailer.SendContactUs(contact); err != nil {
		models.LogError(err)
	} //send mail notification

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.ValidResponse(http.StatusOK, contact, "success"))
}

//ContactUs godoc
//@Summary Handle sending customer review
//@Description Accept JSON data for customer review and message objects then returns valid response
//@Accept json
//@produce json
//@Tags Hybrid APIs
//@Param   Review      body models.Review true  "The Review Data"
//@Success 200 {object} models.Review	"ok"
//@Failure 400 {object} models.ResponseBody "Check Response Message"
//@Router /user/review/ [post]
func SendReview(w http.ResponseWriter, r *http.Request) {
	var reviewData models.Review

	err := decodeJSONBody(w, r, &reviewData)
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

	if err = reviewData.HandleReview(); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(models.ValidResponse(http.StatusInternalServerError, err.Error(), "error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.ValidResponse(http.StatusOK, reviewData, "success"))
}

type malformedRequest struct {
	status int
	msg    string
}

func (mr *malformedRequest) Error() string {
	return mr.msg
}

func decodeJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not applicable/json"
			return &malformedRequest{status: http.StatusUnsupportedMediaType, msg: msg}
		}
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Request body contains badly-formed JSON (at poistion %d", syntaxError.Offset)
			return &malformedRequest{status: http.StatusBadRequest, msg: msg}

		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			return &malformedRequest{status: http.StatusBadRequest, msg: msg}

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			return &malformedRequest{status: http.StatusBadRequest, msg: msg}

		case errors.Is(err, io.EOF):
			msg := "Request body must not be empty"
			return &malformedRequest{status: http.StatusBadRequest, msg: msg}

		case err.Error() == "http: request body too large":
			msg := "Request body must not be larger than 1MB"
			return &malformedRequest{status: http.StatusRequestEntityTooLarge, msg: msg}

		default:
			return err
		}
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		msg := "Request body must only contain a single JSON object"
		return &malformedRequest{status: http.StatusBadRequest, msg: msg}
	}

	return nil
}
