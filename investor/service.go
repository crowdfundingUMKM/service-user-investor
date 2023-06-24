package investor

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"service-user-investor/helper"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	IsPhoneAvailable(input CheckPhoneInput) (bool, error)

	DeactivateAccountUser(input DeactiveUserInput, adminId string) (bool, error)
	GetAdminId(input AdminIdInput) (string, error)

	ActivateAccountUser(input DeactiveUserInput, adminId string) (bool, error)

	GetUserByUnixID(UnixID string) (User, error)

	SaveToken(UnixID string, Token string) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAdminId(input AdminIdInput) (string, error) {
	adminID := helper.UserAdmin{}
	adminID.UnixAdmin = input.UnixID
	// fetch get /getAdminID from service api
	serviceAdmin := os.Getenv("SERVICE_ADMIN")
	// if service admin is empty return error
	if serviceAdmin == "" {
		return adminID.UnixAdmin, errors.New("Service admin is empty")
	}
	resp, err := http.Get(serviceAdmin + "/api/v1/admin/getAdminID/" + adminID.UnixAdmin)
	if err != nil {
		return adminID.UnixAdmin, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return adminID.UnixAdmin, errors.New("Failed to get admin status")
	}

	var response helper.AdminStatusResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	if response.Meta.Code != 200 {
		return "", errors.New(response.Meta.Message)
	} else if response.Data.StatusAccountAdmin == "deactive" {
		return "", errors.New("Admin account is deactive")
	} else if response.Data.StatusAccountAdmin == "active" {
		return adminID.UnixAdmin, nil
	} else {
		return "", errors.New("Invalid admin status")
	}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.UnixID = uuid.New().String()[:12]
	user.Name = input.Name
	user.Email = input.Email
	user.Phone = input.Phone

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash)
	// convert data os env to string
	user.StatusAccount = string(os.Getenv("STATUS_ACCOUNT"))

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("No user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))

	if err != nil {
		return user, err
	}

	return user, nil
}

// save token to database
func (s *service) SaveToken(UnixID string, Token string) (User, error) {
	user, err := s.repository.FindByUnixID(UnixID)
	user.Token = Token
	_, err = s.repository.UpdateToken(user)

	if err != nil {
		return user, err
	}

	return user, nil
}

//end save token to database

func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *service) IsPhoneAvailable(input CheckPhoneInput) (bool, error) {
	phone := input.Phone

	user, err := s.repository.FindByPhone(phone)
	if err != nil {
		return false, err
	}

	if user.UnixID == "" {
		return true, nil
	}

	return false, nil
}

func (s *service) DeactivateAccountUser(input DeactiveUserInput, adminId string) (bool, error) {
	// fin user by unix id
	user, err := s.repository.FindByUnixID(input.UnixID)
	if err != nil {
		return false, err
	}
	if adminId == "" {
		return false, errors.New("Admin ID is empty")
	}
	user.UpdateByAdmin = adminId
	user.StatusAccount = "deactive"
	_, err = s.repository.UpdateStatusAccount(user)

	if err != nil {
		return false, err
	}

	if user.UnixID == "" {
		return true, nil
	}
	return true, nil
}

func (s *service) ActivateAccountUser(input DeactiveUserInput, adminId string) (bool, error) {
	// fin user by unix id
	user, err := s.repository.FindByUnixID(input.UnixID)
	if err != nil {
		return false, err
	}
	if adminId == "" {
		return false, errors.New("Admin ID is empty")
	}
	user.UpdateByAdmin = adminId
	user.StatusAccount = "active"
	_, err = s.repository.UpdateStatusAccount(user)

	if err != nil {
		return false, err
	}

	if user.UnixID == "" {
		return true, nil
	}
	return true, nil
}

func (s *service) GetUserByUnixID(UnixID string) (User, error) {
	user, err := s.repository.FindByUnixID(UnixID)
	if err != nil {
		return user, err
	}

	if user.UnixID == "" {
		return user, errors.New("No user found on with that ID")
	}

	return user, nil
}
