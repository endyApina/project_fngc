package models

import (
	"log"
	"math/rand"
	"os"

	"golang.org/x/crypto/bcrypt"
)

//LogError logs all error to file
func LogError(err error) {
	f, _ := os.OpenFile("logs/errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	logger := log.New(f, "Error: ", log.LstdFlags)
	if err != nil {
		logger.Println(err.Error())
	}
}

func ValidResponse(code int, body interface{}, message string) *ResponseBody {
	var response *ResponseBody
	response.Code = code
	response.Message = message
	response.Body = body

	return response
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func GenerateRandomString(lenght int) string {
	var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")
	b := make([]rune, lenght)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}

func addLoginAttempt(email string) LoginHistory {
	var history LoginHistory
	history.Email = email
	db.Create(&history)
	return history
}

func updateLoginAttempt(history LoginHistory, status string, message string) {
	history.Status = status
	history.Message = message
	db.Model(&history).Updates(LoginHistory{Status: history.Status, Message: history.Message})
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (otpBody *VerifyUser) SaveOTP(email, otp string) error {
	if err := db.Where("email = ?", email).Find(&otpBody).Error; err == nil {
		return nil
	}

	otpBody.Email = email
	otpBody.VerificationOTP = otp
	return nil
}
