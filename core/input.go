package core

type RegisterUserInput struct {
	Name     string `json:"name" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}

type CheckPhoneInput struct {
	Phone string `json:"phone" binding:"required"`
}

type DeactiveUserInput struct {
	UnixID string `json:"unix_id" binding:"required"`
}

type ActiveUserInput struct {
	UnixID string `json:"unix_id" binding:"required"`
}

type GetUserIdInput struct {
	UnixID string `uri:"unix_id" binding:"required"`
}

type UpdateUserInput struct {
	Name    string `json:"name" binding:"required"`
	Phone   string `json:"phone" binding:"required"`
	BioUser string `json:"bio_user" binding:"required"`
}

type UpdatePasswordInput struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}
