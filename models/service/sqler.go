package service

import "bytes"

type Sqler interface {
	GetSql() string
	GetArgs() []interface{}
	GetSqlAndArgs() (string, []interface{})
}

type BaseSql struct {
	sql   bytes.Buffer
	args  []interface{}
}

func (this BaseSql) GetSql() string {
	return this.sql.String()
}

func (this BaseSql) GetArgs() []interface{} {
	return this.args
}

func (this BaseSql) GetSqlAndArgs() (string, []interface{}) {
	return this.sql.String(), this.args
}

func (this *BaseSql) AppendSql(sql string) *BaseSql {
	this.sql.WriteString(sql)
	return this
}

func (this *BaseSql) AppendArg(arg interface{}) *BaseSql {
	this.args = append(this.args, arg)
	return this
}