package database

import sv "gitlab.com/Dank-del/SibylAPI-Go/sibyl/core/sibylValues"

func GetUserFromId(id int64) (*sv.User, error) {
	u := sv.User{}
	SESSION.Where("user_id = ?", id).Take(&u)
	if u.UserID != id {
		// not found
		return nil, nil
	}

	return &u, nil
}

func NewUser(u *sv.User) {
	tx := SESSION.Begin()
	tx.Save(u)
	tx.Commit()
}
