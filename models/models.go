package models

import "time"

type DefaultModel struct {
	ID        int        `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	DefaultModel
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	UserType string `json:"user_type"`
	Address  string `json:"address"`
	Password string `json:"password"`
	Status   string `json:"status"`
	UserID   string `json:"user_id"`
}

type UserRegistrationData struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	UserType string `json:"user_type"`
	Address  string `json:"address"`
	Password string `json:"password"`
}

type UserData struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	UserType string `json:"user_type"`
	UserID   string `json:"user_id"`
}

type VerifyUser struct {
	Email           string `json:"email"`
	VerificationOTP string `json:"verification_otp"`
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
	DefaultModel
	ExamProfile     string `json:"exam_profile"`
	ClassType       string `json:"class_type"`
	TrainingType    string `json:"training_type"`
	TrainigDuration string `json:"training_duration"`
	StudyPack       bool   `json:"study_pack"`
	PersonalTutor   bool   `json:"personal_tutor"`
	TutorType       string `json:"tutor_type"`
	ExamID          string `json:"exam_id"`
	StudentID       string `json:"student_id"` //logged in user unique id
	ExamCost        string `json:"exam_cost"`
	ExamStatus      string `json:"ezam_status"`
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

type LoginHistory struct {
	DefaultModel
	Email   string `json:"email"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type AuthToken struct {
	TokenType string `json:"token_type"`
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
}

type LoggedInData struct {
	User  User      `json:"user"`
	Token AuthToken `json:"token_data"`
}

type ContactUs struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Subject     string `json:"subject"`
	Message     string `json:"message"`
}
