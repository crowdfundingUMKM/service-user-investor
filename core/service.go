package core

import (
	"errors"
	"os"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	IsPhoneAvailable(input CheckPhoneInput) (bool, error)

	DeactivateAccountUser(input DeactiveUserInput, adminId string) (bool, error)

	ActivateAccountUser(input DeactiveUserInput, adminId string) (bool, error)

	DeleteAccountUser(UnixID string) (User, error)

	GetAllUsers() ([]User, error)

	SaveAvatar(UnixID string, fileLocation string) (User, error)

	GetUserByUnixID(UnixID string) (User, error)
	UpdateUserByUnixID(UnixID string, input UpdateUserInput) (User, error)
	UpdatePasswordByUnixID(UnixID string, input UpdatePasswordInput) (User, error)

	SaveToken(UnixID string, Token string) (User, error)
	DeleteToken(UnixID string) (User, error)

	ReportAdmin(UnixID string, input ReportToAdminInput) (NotifInvestor, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.UnixID = uuid.New().String()[:12]
	user.Name = input.Name
	user.Email = input.Email
	user.Phone = input.Phone
	user.AvatarFileName = "/crwdstorage/avatar_investor/dafault-avatar.png"

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
	user.UpdateIDAdmin = adminId
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
	user.UpdateIDAdmin = adminId
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

// delete user by admin
func (s *service) DeleteAccountUser(UnixID string) (User, error) {
	// fin user by unix id
	user, err := s.repository.FindByUnixID(UnixID)
	_, err = s.repository.DeleteUser(user)

	if err != nil {
		return user, err
	}

	return user, nil
}

// get all users
func (s *service) GetAllUsers() ([]User, error) {
	users, err := s.repository.GetAllUser()
	if err != nil {
		return users, err
	}
	return users, nil
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

func (s *service) UpdateUserByUnixID(UnixID string, input UpdateUserInput) (User, error) {
	user, err := s.repository.FindByUnixID(UnixID)
	if err != nil {
		return user, err
	}

	if user.UnixID == "" {
		return user, errors.New("No user found on with that ID")
	}

	user.Name = input.Name
	user.Phone = input.Phone
	user.BioUser = input.BioUser
	user.Addreas = input.Addreas
	user.Country = input.Country
	user.FBLink = input.FBLink
	user.IGLink = input.IGLink

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *service) UpdatePasswordByUnixID(UnixID string, input UpdatePasswordInput) (User, error) {
	user, err := s.repository.FindByUnixID(UnixID)
	if err != nil {
		return user, err
	}

	if user.UnixID == "" {
		return user, errors.New("No user found on with that ID")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.OldPassword))

	if err != nil {
		return user, err
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword), bcrypt.MinCost)

	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash)

	updatedUser, err := s.repository.UpdatePassword(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

// logout
func (s *service) DeleteToken(UnixID string) (User, error) {
	user, err := s.repository.FindByUnixID(UnixID)
	if err != nil {
		return user, err
	}

	if user.UnixID == "" {
		return user, errors.New("No user found on with that ID")
	}

	user.Token = ""

	updatedUser, err := s.repository.UpdateToken(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *service) SaveAvatar(UnixID string, fileLocation string) (User, error) {
	user, err := s.repository.FindByUnixID(UnixID)
	if err != nil {
		return user, err
	}

	user.AvatarFileName = fileLocation

	updatedUser, err := s.repository.UploadAvatarImage(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

// ReportAdmin
func (s *service) ReportAdmin(UnixID string, input ReportToAdminInput) (NotifInvestor, error) {
	notif := NotifInvestor{}
	notif.UserInvestorId = UnixID
	notif.Title = input.Title
	notif.Description = input.Description
	notif.TypeError = input.TypeError
	notif.StatusNotif = 1

	newNotif, err := s.repository.SaveReport(notif)
	if err != nil {
		return newNotif, err
	}
	return newNotif, nil
}
