package models

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func (regData User) HandleRegistration() error {
	_ = godotenv.Load("conf.env")
	err := regData.validateData() //validate registration data
	if err != nil {
		return err
	}

	if err = db.Where("email = ?", regData.Email).Find(&User{}).Error; err == nil {
		return errors.New("user email address already exist")
	} // check if user email exists

	if err = regData.verifyStatus(); err != nil {
		return err
	} //verify user status

	regData.Status = os.Getenv("unverified_status")
	regData.UserID = uuid.NewString()
	hashPassword, _ := HashPassword(regData.Password)
	regData.Password = hashPassword

	if err = db.Create(&regData).Error; err != nil {
		return err
	}

	regData.Password = ""

	return nil
}

func (regData User) validateData() error {
	if regData.Email == "" {
		return errors.New("email address cannot be empty")
	}

	if regData.FullName == "" {
		return errors.New("full name cannot be empty")
	}

	if regData.UserType == "" {
		return errors.New("invalid user type")
	}

	if regData.Password == "" {
		return errors.New("user must have a secured password")
	}

	return nil
}

func (regData *User) verifyStatus() error {
	_ = godotenv.Load("conf.env")
	if regData.UserType == os.Getenv("student_type") {
		return nil
	}

	if regData.UserType == os.Getenv("tutor_type") {
		return nil
	}

	return errors.New("invalid user type")
}

func (user *User) Login(loginData LoginData) error {
	_ = godotenv.Load("conf.env")
	if loginData.Email == "" {
		return errors.New("empty email string")
	}

	if loginData.Password == "" {
		return errors.New("invalid login credentials")
	}

	loginAttempt := addLoginAttempt(loginData.Email)
	if err := db.Where("email = ?", loginData.Email).Find(&user).Error; err != nil {
		updateLoginAttempt(loginAttempt, "failed", "user does not exist")
		errorMessage := errors.New("user does not exist")
		LogError(err)
		return errorMessage
	} //check if user email exist

	passwordMatch := checkPasswordHash(loginData.Password, user.Password)
	if !passwordMatch {
		updateLoginAttempt(loginAttempt, "failed", "invalid login credentials")
		return errors.New("invalid login credentials")
	}
	if user.Status == os.Getenv("unverified_status") {
		return errors.New("unverified user")
	}

	return nil
}

func (authData *AuthToken) GenerateTokenString(email string) error {
	_ = godotenv.Load("conf.env")
	tokenClaim := os.Getenv("jwt_secret")

	expiresAt := time.Now().Add(time.Minute * 10000).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   expiresAt,
	})

	tokenString, err := token.SignedString([]byte(tokenClaim))
	if err != nil {
		return err
	}

	authData.Token = tokenString
	authData.TokenType = "Bearer"
	authData.ExpiresIn = expiresAt

	return nil
}

func (user *User) VerifyOTP(otpBody VerifyUser) error {
	if err := db.Where("email = ? AND verification_otp = ?", otpBody.Email, otpBody.VerificationOTP).Find(&VerifyUser{}).Error; err != nil {
		LogError(err)
		return errors.New("invalid otp verfication code")
	}
	db.Delete(&otpBody)
	if err := db.Where("email = ?", otpBody.Email).Find(&user).Error; err != nil {
		errM := errors.New("error finding user object")
		LogError(errM)
		return errM
	}
	if user.Email == "" {
		LogError(errors.New("error getting user data when updating"))
		return errors.New("error getting user data when updating")
	}

	verified := os.Getenv("verified_status")
	if err := db.Model(&User{}).Where("email = ?", otpBody.Email).Update("status", verified).Error; err != nil {
		LogError(err)
		return errors.New("error verifying user data")
	}

	return nil
}

func (tutor *TutorRegistration) RegisterTutor() error {
	_ = godotenv.Load("conf.env")
	err := tutor.validateData()
	if err != nil {
		return err
	}
	tutor.Password = ""

	if err := db.Create(&tutor).Error; err != nil {
		return err
	}

	return nil
}

func (tutorData *TutorRegistration) validateData() error {
	if tutorData.FirstName == "" || tutorData.LastName == "" {
		return errors.New("empty name field")
	}

	if tutorData.Email == "" || tutorData.TutorType == "" {
		return errors.New("empty tutor or email field")
	}

	if tutorData.NigerianResident {
		if tutorData.StateOfResidence == "" {
			return errors.New("unspciefied state of residence")
		}
	} else {
		if tutorData.Residence == "" {
			return errors.New("unspciefied state of residence")
		}
	}

	return nil
}
