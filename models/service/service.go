package service

import (
	"errors"
	"mgr2/conf"
)

var (
	ErrQuery  = errors.New(conf.MsgQuery)
	ErrInsert = errors.New(conf.MsgInsert)
	ErrUpdate = errors.New(conf.MsgUpdate)
	ErrDelete = errors.New(conf.MsgDelete)

	ErrArgument        = errors.New(conf.MsgArgument)
	ErrDataDuplication = errors.New(conf.MsgDataDuplication)
)
