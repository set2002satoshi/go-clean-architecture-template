package types

import (
	"fmt"
)

const (
	INITIAL_ID       = 0
	INITIAL_REVISION = 1
	DEFAULT_NAME     = "anonymous"
)

type IDENTIFICATION int

func NewIDENTIFICATION(id int) (IDENTIFICATION, error) {
	if id < 0 {
		return IDENTIFICATION(0), fmt.Errorf("couldn't setting negative number")
	}
	return IDENTIFICATION(id), nil
}

func NewOneOrMoreIDENTIFICATION(id int) (IDENTIFICATION, error) {
	if id == 0 {
		return 0, fmt.Errorf("couldn't setting negative number or zero")
	}
	return NewIDENTIFICATION(id)
}

type REVISION int

// type Roll int

// const (
// 	AdminRoll Roll = iota + 1
// 	MembersRoll
// )

// func NewRoll(i int) (Roll, error) {
// 	t := Roll(i)
// 	switch t {
// 	case AdminRoll:
// 		return AdminRoll, nil
// 	case MembersRoll:
// 		return MembersRoll, nil
// 	}

// 	return t, fmt.Errorf("invalid user roll")
// }

type AccessROLL string

const (
	ADMIN_ROLL  = AccessROLL("admin")
	MEMBER_ROLL = AccessROLL("member")
	EDITOR_ROLL = AccessROLL("editor")
	GUESTS      = AccessROLL("guests")
)

func NewAccessROLL(user string) (AccessROLL, error) {

	ur := AccessROLL(user)

	switch ur {
	case ADMIN_ROLL:
		return ur, nil
	case MEMBER_ROLL:
		return ur, nil
	case EDITOR_ROLL:
		return ur, nil
	case GUESTS:
		return ur, nil
	default:

	}

	return ur, fmt.Errorf("%s is disabled roll", ur)
}
