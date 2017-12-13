package userser

import (
	"mgr2/models"
)

type InvalidEnum int

const (
	Invalid InvalidEnum = iota
	Valid
)

type User struct {
	Id        int64
	Username  string
	Password  string
	GmtCreate []uint8
	Invalid   InvalidEnum
}

func GetUserById(id int64) *User {
	db := models.GetDB()

	var user User
	row := db.QueryRow("select tmu.id, tmu.username, tmu.password, tmu.create_time, tmu.invalid from t_mgr_user tmu where tmu.id = ?", id)

	err := row.Scan(
		&user.Id,
		&user.Username,
		&user.Password,
		&user.GmtCreate,
		&user.Invalid,
	)

	if err != nil {
		panic(err)
	}

	return &user
}
