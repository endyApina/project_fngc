package controller

import (
	"fmt"
	"net/http"
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

}
