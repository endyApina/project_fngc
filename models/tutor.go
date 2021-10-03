package models

import "errors"

func (tutor *TutorEducationalData) HandleTutorEducation() (error, User) {
	var tutorData User
	if err := db.Where("user_id = ?", tutor.TutorID).Find(&tutorData).Error; err != nil {
		return errors.New("tutor id does not exist"), tutorData
	}
	err := tutor.handleDegree()
	if err != nil {
		return err, User{}
	}

	err = tutor.handleWorkExperience()
	if err != nil {
		return err, User{}
	}

	err = tutor.handleSchedule()
	if err != nil {
		return err, User{}
	}

	var education TutorEducation
	education.CertifiedTeacher = tutor.CertifiedTeacher
	education.Country = tutor.Country
	education.SubjectSpeciality = tutor.SubjectSpeciality
	education.TutorID = tutor.TutorID

	if err = db.Create(&education).Error; err != nil {
		return errors.New("error creating education data"), User{}
	}

	return nil, User{}
}

func (tutor *TutorEducationalData) handleSchedule() error {
	err := tutor.Schedule.validateSchedule()
	if err != nil {
		return err
	}

	if err = db.Create(&tutor).Error; err != nil {
		return errors.New("error creating new tutor schedule")
	}
	return nil
}

func (schedule *TutorSchedule) validateSchedule() error {
	if schedule.Hours == "" {
		return errors.New("empty number of work hours")
	}

	if len(schedule.Weekdays) == 0 {
		return errors.New("empty working weekdays")
	}

	if schedule.Weekends {
		if len(schedule.WeekendHours) == 0 {
			return errors.New(("empty weekend hours"))
		}
	}

	return nil
}

func (tutor *TutorEducationalData) handleDegree() error {
	if len(tutor.EducationalDegree) == 0 {
		return errors.New("empty degree array")
	}

	for _, degree := range tutor.EducationalDegree {
		degree.TutorID = tutor.TutorID
		err := degree.validateObejct()
		if err != nil {
			return err
		}
		if err := db.Create(&degree).Error; err != nil {
			return errors.New("error creating new degree")
		}
	}

	return nil
}

func (degree *TutorDegree) validateObejct() error {
	if degree.Degree == "" {
		return errors.New("empty degree string at")
	}

	if degree.College == "" {
		return errors.New("empty degree college")
	}

	if degree.Major == "" {
		return errors.New("empty degree major")
	}

	if degree.Year == "" {
		return errors.New("empty year of degree")
	}

	if degree.TutorID == "" {
		return errors.New("empty tutor id")
	}

	return nil
}

func (tutor *TutorEducationalData) handleWorkExperience() error {
	if len(tutor.WorkExperience) == 0 {
		return errors.New("empty degree array")
	}

	for _, employment := range tutor.WorkExperience {
		employment.TutorID = tutor.TutorID
		err := employment.validateEmployment()
		if err != nil {
			return err
		}
		if err := db.Create(&employment).Error; err != nil {
			return errors.New("error creating new degree")
		}
	}

	return nil
}

func (employment *TutorEmployment) validateEmployment() error {
	if employment.Employer == "" {
		return errors.New("empty employer details")
	}

	if employment.TimeAgreement == "" {
		return errors.New("empty empty employment duration")
	}

	if employment.JobTitle == "" {
		return errors.New("empty employment job title")
	}

	if employment.DateOfEmployment == "" {
		return errors.New("empty employment date of employment")
	}

	if employment.JobIndustry == "" {
		return errors.New("empty job industry")
	}

	if employment.TutorID == "" {
		return errors.New("empty tutor id")
	}

	return nil
}
