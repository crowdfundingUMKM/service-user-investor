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
	Token         string `json:"token"`
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
		Token:         token,
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
	Addreas       string `json:"address"`
	Country       string `json:"country"`
	FBLink        string `json:"fb_link"`
	IGLink        string `json:"ig_link"`
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
		Addreas:       user.Addreas,
		Country:       user.Country,
		FBLink:        user.FBLink,
		IGLink:        user.IGLink,
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
	if updatedUser.Addreas != "" {
		formatter.Addreas = updatedUser.Addreas
	}
	if updatedUser.Country != "" {
		formatter.Country = updatedUser.Country
	}
	if updatedUser.FBLink != "" {
		formatter.FBLink = updatedUser.FBLink
	}
	if updatedUser.IGLink != "" {
		formatter.IGLink = updatedUser.IGLink
	}
	return formatter
}

// notify formater
type NotifyFormatter struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	TypeError   string `json:"type_error"`
	StatusNotif int    `json:"status_notif"`
}

func FormatterNotify(notify NotifInvestor) NotifyFormatter {
	formatter := NotifyFormatter{
		ID:          notify.ID,
		Title:       notify.Title,
		Description: notify.Description,
		TypeError:   notify.TypeError,
		StatusNotif: notify.StatusNotif,
	}
	return formatter
}

// for api to other service
type UserInvestor struct {
	UnixInvestor          string `json:"unix_investor"`
	StatusAccountInvestor string `json:"status_account_investor"`
}

// get user Investor status
func FormatterUserInvestorID(user User) UserInvestor {
	formatter := UserInvestor{
		UnixInvestor:          user.UnixID,
		StatusAccountInvestor: user.StatusAccount,
	}
	return formatter
}
