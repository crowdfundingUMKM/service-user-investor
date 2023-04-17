package investor

type UserInvestorFormatter struct {
	ID     int    `json:"id"`
	UnixID string `json:"unix_id"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Email  string `json:"email"`
	Token  string `json:"token"`
	Role   string `json:"role"`
}

func FormatterUser(user User, token string) UserInvestorFormatter {
	formatter := UserInvestorFormatter{
		ID:     user.ID,
		UnixID: user.UnixID,
		Name:   user.Name,
		Phone:  user.Phone,
		Email:  user.Email,
		Token:  token,
		Role:   user.Role,
	}
	return formatter
}
