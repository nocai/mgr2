package models

import (
	"bytes"
	"strconv"
	"mgr2/conf"
)

type Sqler interface {
	GetSql() string
	GetCountSql() string
	GetPageSql(pageNo, pageSize int64) string

	GetArgs() []interface{}

	AppendSqlAndArgs(sql string, arg interface{}) Sqler
}

type BaseSql struct {
	sql  bytes.Buffer
	args []interface{}
}

func (this BaseSql) GetSql() string {
	return this.sql.String()
}

func (this BaseSql) GetArgs() []interface{} {
	return this.args
}

func (this *BaseSql) AppendSqlAndArgs(sql string, arg interface{}) Sqler {
	this.sql.WriteString(sql)
	this.args = append(this.args, arg)
	return this
}

func (this *BaseSql) GetCountSql() string {
	if this.sql.Len() > 0 {
		return "select count(1) from (" + this.sql.String() + ") as tttttttttttttttttttttt"
	}
	panic(conf.MsgArgument)
}

func (this *BaseSql) GetPageSql(pageNo, pageSize int64) string {
	if pageNo > 0 && pageSize > 0 {
		startIndex := (pageNo - 1) * pageSize
		return this.GetSql() + " limit " + strconv.FormatInt(startIndex, 10) + ", " + strconv.FormatInt(pageSize, 10)
	}
	panic(conf.MsgArgument)
}

func NewDefaultSqler(sql string) Sqler {
	var sqler BaseSql
	sqler.sql.WriteString(sql)
	return &sqler
}
