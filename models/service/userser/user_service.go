package userser

import (
	"log"
	"mgr2/models"
	"mgr2/models/service"
	"time"
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

func (this User) GetSqler() service.Sqler {
	var sqler service.BaseSql
	sqler.AppendSql("select t.id, t.username, t.password, t.create_time, t.update_time, t.invalid from t_mgr_user t where 1 = 1")
	if this.Id != 0 {
		sqler.AppendSql(" and t.id = ?")
		sqler.AppendArg(this.Id)
	}
	if this.Username != "" {
		sqler.AppendSql(" and t.username = ?")
		sqler.AppendArg(this.Username)
	}
	if this.Password != "" {
		sqler.AppendSql(" and t.password = ?")
		sqler.AppendArg(this.Password)
	}
	if !this.GmtCreate.IsZero() {
		sqler.AppendSql(" and t.create_time = ?")
		sqler.AppendArg(this.GmtCreate)
	}
	if !this.GmtUpdate.IsZero() {
		sqler.AppendSql(" and t.update_time = ?")
		sqler.AppendArg(this.GmtUpdate)
	}

	return &sqler
}

func QueryUser(userKey User) []User {
	sqler := userKey.GetSqler()

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

func GetUserById(id int64) *User {
	if id <= 0 {
		//log.Println("id =", id)
		log.Panicln("id = ", id)
		return nil
	}
	userKey := User{Id: id}

	users := QueryUser(userKey)
	if users != nil && len(users) > 0 {
		return &users[0]
	}
	return nil
}
