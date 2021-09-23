package controller

import "net/http"

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
