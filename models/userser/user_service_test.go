package userser

import (
	"testing"
	"fmt"
)

func TestGetUserById(t *testing.T) {
	u := GetUserById(1)
	fmt.Println(u)
	//fmt.Println(string(u.GmtCreate))



}
