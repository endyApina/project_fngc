package controller

import "net/http"

//ExaminationApply godoc
//@Summary Handle apply for student examination
//@Description Accept JSON data for examination objects and returns valid response
//@Accept json
//@produce json
//@Tags Student
//@Param   ExaminationPreparationData      body models.ExamPreparation true  "The Student Examination Data"
//@Success 200 {object} models.ExamPreparation	"ok"
//@Failure 400 {object} models.ResponseBody "Check Response Message"
//@Router /student/examination/ [post]
func ExaminationApply(w http.ResponseWriter, r *http.Request) {

}

//StudentCurriculum handle student curriculum
//@Summary Handles getting student curriculum
//@Description Gets the list of all curricul courses  being taken by the student
//@Accept json
//@produce json
//@Tags Student
//@Success 200 {object} models.StudentCurriculum	"ok"
//@Failure 400 {object} models.ResponseBody "Check Response Message"
//@Router /student/curriculum/ [get]
func StudentCurriculum(w http.ResponseWriter, r *http.Request) {

}

//NewTutor godoc
//@Summary Handle student request for new tutor
//@Description Accept JSON data for new tutor request
//@Accept json
//@produce json
//@Tags Student
//@Param   RequestData      body models.RequestTutor true  "The Student Examination Data"
//@Success 200 {object} models.RequestTutor	"ok"
//@Failure 400 {object} models.ResponseBody "Check Response Message"
//@Router /student/tutor/request [post]
func NewTutor(w http.ResponseWriter, r *http.Request) {

}
