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

	var mailBody *regMailData
	mailBody.Email = regData.Email
	mailBody.FullName = regData.FullName
	mailBody.VerificationCode = verificationOTP

	var verify *models.VerifyUser
	_ = verify.SaveOTP(regData.Email, verificationOTP)

	// _ = godotenv.Load("conf.env")
	mailSubject := os.Getenv("mail_subject_prefix") + "Registration Successful"
	newRequestData := NewRequest(mailBody.Email, mailSubject)
	go newRequestData.AppSendMail(templatePath, mailBody)

	return nil
}

func SendExamPrepMail(user models.User) error {
	templatePath := os.Getenv("template_path") + "registration.html"

	var mailBody *regMailData
	mailBody.Email = user.Email
	mailBody.FullName = user.FullName

	// _ = godotenv.Load("conf.env")
	mailSubject := os.Getenv("mail_subject_prefix") + "Exam application successful"
	newRequestData := NewRequest(mailBody.Email, mailSubject)
	go newRequestData.AppSendMail(templatePath, mailBody)

	return nil
}

func SendStudyAbroad(user models.StudyAbroad) error {
	templatePath := os.Getenv("template_path") + "abroad.html"

	var mailBody *regMailData
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

	var mailBody *regMailData
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