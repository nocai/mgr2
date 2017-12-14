package userser

import (
	"log"
	"mgr2/models"
	"time"
	"mgr2/models/service"
)

type InvalidEnum int

func (this InvalidEnum) IsDefined() bool {
	return this == Invalid || this == Valid
}

func (this InvalidEnum) String() string {
	if this.IsDefined() {
		if this == Invalid {
			return "InValid"
		} else {
			return "Valid"
		}
	}
	panic("this InvalidEnum is not defined")
}

const (
	Invalid InvalidEnum = iota
	Valid
)

type User struct {
	Id        int64
	Username  string
	Password  string
	GmtCreate time.Time
	GmtUpdate time.Time
	Invalid   InvalidEnum
}

type UserKey struct {
	User

	GmtCreateStart time.Time
	GmtCreateEnd   time.Time
	GmtUpdateStart time.Time
	GmtUpdateEnd   time.Time
}

func (this UserKey) GetSqler() models.Sqler {
	sqler := models.NewDefaultSqler("select t.id, t.username, t.password, t.create_time, t.update_time, t.invalid from t_mgr_user t where 1 = 1")

	if this.Id != 0 {
		sqler.AppendSqlAndArgs(" and t.id = ?", this.Id)
	}
	if this.Username != "" {
		sqler.AppendSqlAndArgs(" and t.username = ?", this.Username)
	}
	if this.Password != "" {
		sqler.AppendSqlAndArgs(" and t.password = ?", this.Password)
	}
	if !this.GmtCreate.IsZero() {
		sqler.AppendSqlAndArgs(" and t.create_time = ?", this.GmtCreate)
	}
	if !this.GmtUpdate.IsZero() {
		sqler.AppendSqlAndArgs(" and t.update_time = ?", this.GmtUpdate)
	}

	if !this.GmtCreateStart.IsZero() {
		sqler.AppendSqlAndArgs(" and t.create_time > ?", this.GmtCreateStart)
	}
	if !this.GmtCreateEnd.IsZero() {
		sqler.AppendSqlAndArgs(" and t.create_time < ?", this.GmtCreateEnd)
	}

	if !this.GmtUpdateStart.IsZero() {
		sqler.AppendSqlAndArgs(" and t.update_time > ?", this.GmtUpdateStart)
	}
	if !this.GmtUpdateEnd.IsZero() {
		sqler.AppendSqlAndArgs(" and t.update_time < ?", this.GmtUpdateEnd)
	}
	return sqler
}

func QueryUser(key UserKey) []User {
	sqler := key.GetSqler()

	db := models.GetDB()
	rows, err := db.Query(sqler.GetSql(), sqler.GetArgs()...)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var result []User
	for rows.Next() {
		var temp User
		if err := rows.Scan(&temp.Id, &temp.Username, &temp.Password, &temp.GmtCreate, &temp.GmtUpdate, &temp.Invalid); err != nil {
			panic(err)
		}
		result = append(result, temp)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}

	return result
}

func PageUser(pageNo, pageSize int64, key User) *models.Pager {
return nil
}

func GetUserById(id int64) *User {
	if id <= 0 {
		log.Println("id =", id)
		panic(service.ErrArgument)
	}

	users := QueryUser(UserKey{
		User:User{
			Id:id,
		},
	})
	if users != nil && len(users) > 0 {
		return &users[0]
	}
	return nil
}
