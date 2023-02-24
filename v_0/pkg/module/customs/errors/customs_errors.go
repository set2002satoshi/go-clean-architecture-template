package errors

import "fmt"

const (
	ERR0000 = "ERROR0000"

	UNDEFINED = "UNDEFINED"
)

var ErrMap = map[string]string{
	ERR0000: "error code undetermined",

	UNDEFINED: "undefined",
}

type (
	customError struct {
		ErrorMap map[string]string
	}
	iCustomError interface {
		error
		add(code string)
		addSet(code, message string)
		combine(err error)
		isContain(code string) bool
		isEmpty() bool
		getMap() map[string]string
		wrap(code, message string)
	}
)

var _ iCustomError = &customError{}

func NewCustomError(codes ...string) error {
	e := &customError{ErrorMap: map[string]string{}}
	for _, s := range codes {
		e.add(s)
	}
	return e
}

func New(s string) error {
	c := &customError{ErrorMap: map[string]string{}}
	c.addSet(UNDEFINED, s)
	return c
}

func Combine(origin, new error) error {
	if origin == nil && new == nil {
		return nil
	}
	if oErr, ok := origin.(iCustomError); ok {
		if new != nil {
			oErr.combine(new)
		}
		return oErr
	}
	if nErr, ok := new.(iCustomError); ok {
		if origin != nil {
			nErr.addSet(UNDEFINED, origin.Error())
		}
		return nErr
	}
	cErr := &customError{ErrorMap: map[string]string{}}
	cErr.combine(origin)
	cErr.combine(new)
	return cErr
}

func Wrap(err error, code, message string) error {
	if cErr, ok := err.(iCustomError); ok {
		cErr.wrap(code, message)
		return cErr
	}
	cErr := &customError{ErrorMap: map[string]string{}}
	cErr.wrap(code, message)
	return cErr
}

func Add(err error, code string) error {
	if cErr, ok := err.(iCustomError); ok {
		cErr.add(code)
		return cErr
	}
	cErr := &customError{ErrorMap: map[string]string{}}
	if err != nil {
		cErr.combine(err)
	}
	cErr.add(code)
	return cErr
}

func IsContain(err error, code string) bool {
	if cErr, ok := err.(iCustomError); ok {
		return cErr.isContain(code)
	}
	return false
}

func IsEmpty(err error) bool {
	if cErr, ok := err.(iCustomError); ok {
		return cErr.isEmpty()
	}
	return err == nil
}

func ToMap(err error) map[string]string {
	if cErr, ok := err.(iCustomError); ok {
		return cErr.getMap()
	}
	c := &customError{ErrorMap: map[string]string{}}
	if err != nil {
		c.addSet(UNDEFINED, err.Error())
	}
	return c.getMap()
}

func (c *customError) add(code string) {
	if _, ok := ErrMap[code]; ok {
		c.ErrorMap[code] = getMessage(code)
	}
}

func (c *customError) addSet(code, message string) {
	if _, ok := c.ErrorMap[code]; ok {
		c.ErrorMap[code] = fmt.Sprintf("%s, %s", c.ErrorMap[code], message)
		return
	}
	c.ErrorMap[code] = message
}
func (c *customError) combine(err error) {
	if cErr, ok := err.(iCustomError); ok {
		for k, v := range cErr.getMap() {
			c.addSet(k, v)
		}
		return
	}
	if err != nil {
		c.addSet(UNDEFINED, err.Error())
	}
}
func (c *customError) isContain(code string) bool {
	_, ok := c.ErrorMap[code]
	return ok
}

func (c *customError) isEmpty() bool {
	fmt.Println(len(c.ErrorMap))
	return len(c.ErrorMap) == 0
}

func (c *customError) getMap() map[string]string {
	return c.ErrorMap
}

func (c *customError) wrap(code, message string) {
	if _, ok := c.ErrorMap[code]; ok {
		c.ErrorMap[code] = fmt.Sprintf("%s=>%s", c.ErrorMap[code], message)
		return
	}
	if _, ok := ErrMap[code]; ok {
		c.ErrorMap[code] = fmt.Sprintf("%s=>%s", ErrMap[code], message)
		return
	}
	c.addSet(UNDEFINED, message)
}

func getMessage(code string) string {
	if s, ok := ErrMap[code]; ok {
		return s
	}
	return ""
}

func (c customError) Error() string {
	var msg string
	initFlag := true
	for _, v := range c.ErrorMap {
		if initFlag {
			msg = v
			initFlag = false
			continue
		}
		msg = fmt.Sprintf("%s, %s", msg, v)
	}

	return msg
}
