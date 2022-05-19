package controller

import (
	"encoding/json"
	"fngc/models"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)


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
