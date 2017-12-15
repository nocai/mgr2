package models

import (
	"log"
)

type Pager struct {
	Page      int64
	Rows      int64
	PageCount int64
	Total    int64       `json:"total"`
	PageList interface{} `json:"rows"`
}

// New
func NewPager(pageNo, pageSize int64, total int64, pageList interface{}) *Pager {
	if pageNo <= 0 || pageSize <= 0 {
		log.Panicln("The pageNo and pageSize can't be 0")
	}

	var pageCount int64
	if total%pageSize == 0 {
		pageCount = total / pageSize
	} else {
		pageCount = total/pageSize + 1
	}

	return &Pager{Page: pageNo, Rows: pageSize, PageCount: pageCount, Total: total, PageList: pageList}
}
