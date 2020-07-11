package model

import (
	"fmt"
	"strings"
)

// User represents a structure of an user.
type User struct {
	Name string
	Age  int
}

func (u *User) String() string {
	return strings.Join([]string{
		fmt.Sprintf("Name: %s", u.Name),
		fmt.Sprintf("Age: %d", u.Age),
	}, ", ")
}
