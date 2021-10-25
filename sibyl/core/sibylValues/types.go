package sibylValues

import "time"

type UserPermission int

type User struct {
	UserID  int64     `json:"user_id" gorm:"primaryKey"`
	Banned  bool      `json:"banned"`
	Reason  string    `json:"reason"`
	Message string    `json:"message"`
	Date    time.Time `json:"date"`
}

type Token struct {
	// the user id
	Id int64 `json:"id" gorm:"primaryKey"`

	// the user hash
	Hash string `json:"hash"`

	// the user's permissions
	Permission UserPermission `json:"permission"`

	// the user's last usage time
	LastUsage time.Time `json:"-"`

	// Creation time
	CreatedAt time.Time `json:"created_at"`
}

type EndpointError struct {
	ErrorCode int    `json:"code"`
	Message   string `json:"message"`
	Origin    string `json:"origin"`
}
