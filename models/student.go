package models

import "errors"

func (studExam *ExamPreparation) ExamApplication() (User, error) {
	err := studExam.validateData()
	if err != nil {
		LogError(err)
		return User{}, err
	}

	var user User
	if err := db.Where("user_id = ?", studExam.StudentID).Find(&user).Error; err != nil {
		LogError(err)
		return User{}, errors.New("invalid user id")
	}

	if err := db.Create(&studExam).Error; err != nil {
		LogError(errors.New("error creating new examination profile"))
	}

	return user, nil
}

func (examData *ExamPreparation) validateData() error {
	if examData.ExamProfile == "" {
		return errors.New("invalid exam profile. exam profile cannot be null")
	}

	if examData.ClassType == "" {
		return errors.New("invalid class type. class type cannot be null")
	}

	if examData.TrainingType == "" {
		return errors.New("invalid training type. training type cannot be null")
	}

	if examData.TrainigDuration == "" {
		return errors.New("kindly specify training duration")
	}

	if examData.PersonalTutor == true {
		if examData.TutorType == "" {
			return errors.New("kindly specify type of tutoring. invalid tutor type")
		}
	}

	if examData.ExamCost == "" {
		return errors.New("kindly specify cost of examination")
	}

	return nil
}
