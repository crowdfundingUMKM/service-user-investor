package core

type UserInvestorFormatter struct {
	ID            int    `json:"id"`
	UnixID        string `json:"unix_id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	BioUser       string `json:"bio_user"`
	Avatar        string `json:"avatar_file_name"`
	StatusAccount string `json:"status_account"`
}

func FormatterUser(user User, token string) UserInvestorFormatter {
	formatter := UserInvestorFormatter{
		ID:            user.ID,
		UnixID:        user.UnixID,
		Name:          user.Name,
		Email:         user.Email,
		Phone:         user.Phone,
		BioUser:       user.BioUser,
		Avatar:        user.AvatarFileName,
		StatusAccount: user.StatusAccount,
	}
	return formatter
}

type UserDetailFormatter struct {
	ID            int    `json:"id"`
	UnixID        string `json:"unix_id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	BioUser       string `json:"bio_user"`
	Avatar        string `json:"avatar_file_name"`
	StatusAccount string `json:"status_account"`
}

func FormatterUserDetail(user User, updatedUser User) UserDetailFormatter {
	formatter := UserDetailFormatter{
		ID:            user.ID,
		UnixID:        user.UnixID,
		Name:          user.Name,
		Email:         user.Email,
		Phone:         user.Phone,
		BioUser:       user.BioUser,
		Avatar:        user.AvatarFileName,
		StatusAccount: user.StatusAccount,
	}
	// read data before update if null use old data
	if updatedUser.Name != "" {
		formatter.Name = updatedUser.Name
	}
	if updatedUser.Phone != "" {
		formatter.Phone = updatedUser.Phone
	}
	if updatedUser.BioUser != "" {
		formatter.BioUser = updatedUser.BioUser
	}
	if updatedUser.AvatarFileName != "" {
		formatter.Avatar = updatedUser.AvatarFileName
	}
	if updatedUser.StatusAccount != "" {
		formatter.StatusAccount = updatedUser.StatusAccount
	}
	return formatter
}
