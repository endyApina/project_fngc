package mailer

import (
	"bytes"
	"html/template"
	"log"
	"os"

	"github.com/joho/godotenv"
	gomail "gopkg.in/gomail.v2"
)

//request ldjf
type request struct {
	to      string
	subject string
	body    string
}

//NewRequest holds new request data
func NewRequest(to string, subject string) *request {
	return &request{
		to:      to,
		subject: subject,
	}
}

func (r *request) AppSendMail(tempName string, item interface{}) {
	err := r.parseTemplate(tempName, item)
	if err != nil {
		log.Println(err.Error())
		f, _ := os.OpenFile("logs/errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer f.Close()

		logger := log.New(f, "Error: ", log.LstdFlags)
		logger.Println(err.Error())
	}
	// models.CheckErr(err, true, "Errors parsing temnplate with data")
	if ok := r.sendEmail(); ok {

	} else {
		defer recover()
		log.Println(err.Error())
		f, _ := os.OpenFile("logs/errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer f.Close()

		logger := log.New(f, "Error: ", log.LstdFlags)
		logger.Println(err.Error())
	}
}

func (r *request) parseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}

func (r *request) sendEmail() bool {
	_ = godotenv.Load("conf.env")
	m := gomail.NewMessage()
	m.SetAddressHeader("From", os.Getenv("mailer_email"), os.Getenv("mailer_header"))
	m.SetHeader("To", r.to)
	m.SetHeader("Subject", r.subject)
	m.SetBody("text/html", r.body)

	d := gomail.NewDialer(os.Getenv("mailer_smtp"), 587, os.Getenv("mailer_email"), os.Getenv("mailer_password"))
	if err := d.DialAndSend(m); err != nil {
		log.Println(err.Error())
		f, _ := os.OpenFile("logs/errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer f.Close()

		logger := log.New(f, "Error: ", log.LstdFlags)
		logger.Println(err.Error())
		// models.CheckErr(err, false, err.Error())
		return false
	}
	log.Println("Email Sent to: " + r.to)
	return true
}
