package investor

import (
	"time"
)

type User struct {
	ID             int
	Name           string
	Email          string
	Phone          int
	PasswordHash   string
	AvatarFileName string
	Role           string
	Token          string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
