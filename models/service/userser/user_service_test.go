package userser

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"os"
	"log"
)

func TestGetUserById(t *testing.T) {
	u := GetUserById(0)
	assert.True(t, u != nil && u.Id == 1)
}

func TestA(t *testing.T) {
	assert.True(t, InvalidEnum(1).IsDefined(), "1 is defined")
	assert.False(t, InvalidEnum(5).IsDefined(), "5 is not defined")
}

func TestB(t *testing.T) {
	var std = log.New(os.Stderr, "", log.LstdFlags)
	std.SetPrefix("TEST")
	std.Println("a")
	std.Panicln("panic")
	std.Fatalln("exit")
	std.Println("b")

}