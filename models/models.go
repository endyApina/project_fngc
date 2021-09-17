package models

import "time"

type DefaultModel struct {
	ID        int        `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

type StudentRegistrationData struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TutorRegistrationData struct {
	FullName  string `json:"full_name"`
	Email     string `json:"email"`
	TutorType string `json:"tutor_type"`
	Address   string `json:"address"`
	Password  string `json:"password"`
}

type UserData struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	UserType string `json:"user_type"`
	UserID   string `json:"user_id"`
}

type VerifyUser struct {
	Email            string `json:"email"`
	VerificationLink string `json:"verification_link"`
}

type ResetPassword struct {
	OldPassword       string `json:"old_password"`
	NewPassword       string `json:"new_password"`
	ResetPasswordLink string `json:"reset_password_link"`
}

type ResponseBody struct {
	Code    int         `json:"code"`
	Body    interface{} `json:"body"`
	Message string      `json:"message"`
}

type ExamPreparation struct {
	Profile         string `json:"profile"`
	ClassType       string `json:"class_type"`
	TrainingType    string `json:"training_type"`
	TrainigDuration string `json:"training_duration"`
	StudyPack       bool   `json:"study_pack"`
	PersonalTutor   bool   `json:"personal_tutor"`
	TutorType       string `json:"tutor_type"`
	PreparationID   string `json:"preparation_id"`
}

type Exam struct {
	DefaultModel
	Exam   string `json:"exam"`
	ExamID string `json:"exam_id"`
}

type StudentCurriculum struct {
	Subject   string `json:"subjject"`
	Class     string `json:"class"`
	ClassType string `json:"class_type"`
}

type RequestTutor struct {
	ExamID    string `json:"exam_id"`
	ExamType  string `json:"exam_type"`
	Gender    string `json:"gender"`
	StudentID string `json:"student_id"`
}

type Task struct {
	DefaultModel
	Task   string `json:"task"`
	UserID string `json:"user_id"`
}

type TutorDashboard struct {
	Task    []Task     `json:"task"`
	Student []UserData `json:"students"`
}

type StudyAbroad struct {
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Phone            string `json:"phone"`
	Email            string `json:"email"`
	DOB              string `json:"dob"`
	MaritalStatus    string `json:"marital_status"`
	Gender           string `json:"gender"`
	Nationality      string `json:"nationality"`
	Address          string `json:"address"`
	HighSchool       string `json:"high_school"`
	EducationalLevel string `json:"educational_level"`
	UniversityName   string `json:"university_name"`
	Course           string `json:"course"`
	Degree           string `json:"degree"`
	YearOfAdmission  string `json:"year_of_admission"`
	GraduationYear   string `json:"graduation_year"`
	DegreeLevel      string `json:"degree_level"`
	DegreeProgramme  string `json:"degree_prgramme"`
	FinancialStatus  string `json:"financial_status"`
	EnrollmentTerm   string `json:"enrollment_term"`
	EnrollmentYear   string `json:"enrollement_year"`
	StudyModel       string `json:"study_model"`
	CountryOfStudy   string `json:"country_of_study"`
}
