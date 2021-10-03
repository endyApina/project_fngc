package models

import (
	"errors"

	"github.com/joho/godotenv"
)

func GetAllExam() ([]ExamPreparation, error) {
	var allExams []ExamPreparation
	if err := db.Find(&allExams).Error; err != nil {
		return []ExamPreparation{}, errors.New("error finding all examination profiles")
	}

	return allExams, nil
}

func GetProfileExam(profile string) ([]ExamPreparation, error) {
	var allExams []ExamPreparation
	if err := db.Where("exam_profile = ?", profile).Find(&allExams).Error; err != nil {
		return []ExamPreparation{}, errors.New("error finding all examination profiles")
	}

	return allExams, nil
}

func GetAllUser(userType string) ([]User, error) {
	_ = godotenv.Load("conf.env")
	var tutor []User
	if err := db.Where("user_type = ?", userType).Find(&tutor).Error; err != nil {
		return []User{}, errors.New("error finding all tutor profiles")
	}

	return tutor, nil
}

func GetAllAbroadStudies() ([]StudyAbroad, error) {
	var allStudies []StudyAbroad
	if err := db.Find(&allStudies).Error; err != nil {
		return []StudyAbroad{}, errors.New("error finding all abroad studies")
	}

	return allStudies, nil
}
