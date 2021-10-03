package models

import "errors"

func (study *StudyAbroad) CreateStudyAbroad() error {
	if err := study.validateData(); err != nil {
		LogError(err)
		return err
	}

	if err := db.Create(&study).Error; err != nil {
		LogError(err)
		return errors.New("error creating new abroad study data")
	}

	return nil
}

func (study *StudyAbroad) validateData() error {
	if study.FirstName == "" || study.LastName == "" {
		return errors.New("please input user first name and last name")
	}

	if study.Email == "" || study.Phone == "" {
		return errors.New("please input user email and phone records")
	}

	if study.DOB == "" || study.MaritalStatus == "" || study.Gender == "" || study.Nationality == "" || study.Address == "" {
		return errors.New("kindly verify that all user personal information has been filled. incomplete user information ")
	}

	if study.HighSchool == "yes" {
		if study.UniversityName == "" {
			return errors.New("empty university name. please specify university name")
		}

		if study.Course == "" {
			return errors.New("empty couse name. please specify name of course")
		}

		if study.Degree == "" {
			return errors.New("empty degree. please specify type of degree")
		}

		if study.YearOfAdmission == "" || study.GraduationYear == "" {
			return errors.New("kindly specify year of admission and graduation year")
		}
	}

	if study.HighSchool == "no" {
		if study.EducationalLevel == "" {
			return errors.New("kindly specify educational level")
		}
	}

	if study.HighSchool == "others" {
		if study.Course == "" {
			return errors.New("please specify course of study")
		}
	}

	if study.DegreeLevel == "" {
		return errors.New("kindly specify the degree level of interest")
	}

	if study.DegreeProgramme == "" {
		return errors.New("kinldy specify the program of interest")
	}

	if study.FinancialStatus == "" {
		return errors.New("kindly show financial status")
	}

	if study.EnrollmentTerm == "" || study.EnrollmentYear == "" {
		return errors.New("kindly specify the term and year of intrest")
	}

	if study.StudyModel == "" {
		return errors.New("kindly specify the model of study")
	}

	if study.CountryOfStudy == "" {
		return errors.New("kindly specify the country of interest")
	}

	return nil
}

func (contactData ContactUs) HandleContactUs() error {
	if err := contactData.validateData(); err != nil {
		LogError(err)
		return err
	}

	if err := db.Create(&contactData).Error; err != nil {
		LogError(err)
		return errors.New("error creating new contact data")
	}

	return nil
}

func (contactData ContactUs) validateData() error {
	if contactData.FirstName == "" || contactData.LastName == "" {
		return errors.New("user first name and last name cannot be empty")
	}

	if contactData.PhoneNumber == "" || contactData.Email == "" {
		return errors.New("kindly fill all records including phone number and email")
	}

	if contactData.Subject == "" || contactData.Message == "" {
		return errors.New("user subject and message cannot be empty")
	}

	return nil
}

func (reviewData Review) HandleReview() error {
	if err := reviewData.validateData(); err != nil {
		LogError(err)
		return err
	}

	if err := db.Create(&reviewData).Error; err != nil {
		LogError(err)
		return errors.New("error creating new review data")
	}

	return nil
}

func (reviewData Review) validateData() error {
	if reviewData.FullName == ""{
		return errors.New("user full name cannot be empty")
	}

	if reviewData.Avatar == "" {
		return errors.New("user avatar cannot be empty")
	}

	if reviewData.Review == "" {
		return errors.New("user review is empty")
	}

	return nil
}