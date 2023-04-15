package investor

type UserInvestorFormatter struct {
	ID     int    `json:"id"`
	UnixID string `json:"unix_id"`
	Name   string `json:"name"`
	Phone  int    `json:"phone"`
	Email  string `json:"email"`
	Token  string `json:"token"`
}

func FormatterUser(user User, token string) UserInvestorFormatter {
	formatter := UserInvestorFormatter{
		ID:     user.ID,
		UnixID: user.UnixID,
		Name:   user.Name,
		Phone:  user.Phone,
		Email:  user.Email,
		Token:  token,
	}
	return formatter
}
