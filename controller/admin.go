package controller

import (
	"encoding/json"
	"fngc/models"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

//GetExaminationProfile godoc
//@Summary Retrieves a list of all examination profile
//@Description Retrieves all examination profiles
//@Accept json
//@produce json
//@Tags Admin
//@Success 200 {object} models.ExamPreparation	"ok"
//@Failure 400 {object} models.ResponseBody "Check Response Message"
//@Router /admin/examination/all [get]
func GetAllExaminationProfile(w http.ResponseWriter, r *http.Request) {
	var examProfile []models.ExamPreparation
	examProfile, err := models.GetAllExam()
	if err != nil {
		models.LogError(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(models.ValidResponse(http.StatusOK, err.Error(), "error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.ValidResponse(http.StatusOK, examProfile, "success"))
}

//GetExaminationProfile godoc
//@Summary Retrieves a list of all examination profile
//@Description Retrieves all examination profiles
//@Accept json
//@produce json
//@Tags Admin
//@Success 200 {object} models.ExamPreparation	"ok"
//@Failure 400 {object} models.ResponseBody "Check Response Message"
//@Router /admin/examination/{profile} [get]
func GetExaminationProfile(w http.ResponseWriter, r *http.Request) {
	profile := chi.URLParam(r, "profile")
	var examProfile []models.ExamPreparation
	examProfile, err := models.GetProfileExam(profile)
	if err != nil {
		models.LogError(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(models.ValidResponse(http.StatusOK, err.Error(), "error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.ValidResponse(http.StatusOK, examProfile, "success"))
}

//GetAllTutor godoc
//@Summary Retrieves a list of all tutors
//@Description Retrieves all tutor data
//@Accept json
//@produce json
//@Tags Admin
//@Success 200 {object} models.User	"ok"
//@Failure 400 {object} models.ResponseBody "Check Response Message"
//@Router /admin/tutor/all [get]
func GetAllTutor(w http.ResponseWriter, r *http.Request) {
	_ = godotenv.Load("conf.env")
	var allTutor []models.User
	userType := os.Getenv("tutor_type")
	allTutor, err := models.GetAllUser(userType)
	if err != nil {
		models.LogError(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(models.ValidResponse(http.StatusOK, err.Error(), "error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.ValidResponse(http.StatusOK, allTutor, "success"))
}

//GetAllStudents godoc
//@Summary Retrieves a list of all students
//@Description Retrieves all students data
//@Accept json
//@produce json
//@Tags Admin
//@Success 200 {object} models.User	"ok"
//@Failure 400 {object} models.ResponseBody "Check Response Message"
//@Router /admin/student/all [get]
func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	_ = godotenv.Load("conf.env")
	var allTutor []models.User
	userType := os.Getenv("student_type")
	allTutor, err := models.GetAllUser(userType)
	if err != nil {
		models.LogError(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(models.ValidResponse(http.StatusOK, err.Error(), "error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.ValidResponse(http.StatusOK, allTutor, "success"))
}

//GetAllAbroadStudies godoc
//@Summary Retrieves a list of all abroad application students
//@Description Retrieves all abroad student data
//@Accept json
//@produce json
//@Tags Admin
//@Success 200 {object} models.StudyAbroad	"ok"
//@Failure 400 {object} models.ResponseBody "Check Response Message"
//@Router /admin/student/abroad/all [get]
func GetAllAbroadStudies(w http.ResponseWriter, r *http.Request) {
	_ = godotenv.Load("conf.env")
	var allStudies []models.StudyAbroad
	allStudies, err := models.GetAllAbroadStudies()
	if err != nil {
		models.LogError(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(models.ValidResponse(http.StatusOK, err.Error(), "error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.ValidResponse(http.StatusOK, allStudies, "success"))
}
