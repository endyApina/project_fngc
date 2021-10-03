package mailer

import (
	"fngc/models"
	"os"

	"github.com/joho/godotenv"
)

type regMailData struct {
	FullName         string `json:"full_name"`
	Email            string `json:"email"`
	VerificationCode string `json:"verification_code"`
	FirstName        string `json:"first_name"`
	Message          string `json:"message"`
}

func SendRegistrationMail(regData models.User) error {
	templatePath := os.Getenv("template_path") + "registration.html"
	verificationOTP := models.GenerateRandomString(6)

	var mailBody regMailData
	mailBody.Email = regData.Email
	mailBody.FullName = regData.FullName
	mailBody.VerificationCode = verificationOTP

	var verify models.VerifyUser
	err := verify.SaveOTP(regData.Email, verificationOTP)
	if err != nil {
		models.LogError(err)
		return err
	}

	// _ = godotenv.Load("conf.env")
	mailSubject := os.Getenv("mail_subject_prefix") + "Registration Successful"
	newRequestData := NewRequest(mailBody.Email, mailSubject)
	go newRequestData.AppSendMail(templatePath, mailBody)

	return nil
}

func SendVerifiedMail(regData models.User) error {
	templatePath := os.Getenv("template_path") + "verified.html"

	var mailBody regMailData
	mailBody.Email = regData.Email
	mailBody.FullName = regData.FullName

	// _ = godotenv.Load("conf.env")
	mailSubject := os.Getenv("mail_subject_prefix") + "Email Verified"
	newRequestData := NewRequest(mailBody.Email, mailSubject)
	go newRequestData.AppSendMail(templatePath, mailBody)

	return nil
}

func SendNewTutorMail(regData models.User) error {
	templatePath := os.Getenv("template_path") + "tutor.html"

	var mailBody regMailData
	mailBody.Email = regData.Email
	mailBody.FullName = regData.FullName

	// _ = godotenv.Load("conf.env")
	mailSubject := os.Getenv("mail_subject_prefix") + "New Tutor Education details"
	newRequestData := NewRequest(mailBody.Email, mailSubject)
	go newRequestData.AppSendMail(templatePath, mailBody)

	return nil
}


func SendExamPrepMail(user models.User) error {
	templatePath := os.Getenv("template_path") + "examination.html"

	var mailBody regMailData
	mailBody.Email = user.Email
	mailBody.FullName = user.FullName

	// _ = godotenv.Load("conf.env")
	mailSubject := os.Getenv("mail_subject_prefix") + "Exam application successful"
	newRequestData := NewRequest(mailBody.Email, mailSubject)
	go newRequestData.AppSendMail(templatePath, mailBody)

	adminMail := os.Getenv("mail_subject_prefix") + "New Exam Preparation"
	adminRequestData := NewRequest(os.Getenv("admin_email"), adminMail)
	adminPath := os.Getenv("template_path") + "admin-exam.html"

	go adminRequestData.AppSendMail(adminPath, mailBody)

	return nil
}

func SendStudyAbroad(user models.StudyAbroad) error {
	templatePath := os.Getenv("template_path") + "abroad.html"

	var mailBody regMailData
	mailBody.Email = user.Email
	mailBody.FullName = user.FirstName

	// _ = godotenv.Load("conf.env")
	mailSubject := os.Getenv("mail_subject_prefix") + "Application to study abroad"
	newRequestData := NewRequest(mailBody.Email, mailSubject)
	go newRequestData.AppSendMail(templatePath, mailBody)

	return nil
}

func SendContactUs(user models.ContactUs) error {
	templatePath := os.Getenv("template_path") + "contact.html"
	_ = godotenv.Load("conf.env")

	var mailBody regMailData
	mailBody.Email = user.Email
	mailBody.FullName = user.FirstName + " " + user.LastName
	mailBody.Message = user.Message

	adminEmail := os.Getenv("admin_email")

	// _ = godotenv.Load("conf.env")
	mailSubject := os.Getenv("mail_subject_prefix") + user.Subject
	newRequestData := NewRequest(adminEmail, mailSubject)
	go newRequestData.AppSendMail(templatePath, mailBody)

	return nil
}
