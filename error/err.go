package error

import "fmt"

type CodeError struct {
	Code     int
	Err      string
	Message  string
	Location string
}

func (c *CodeError) Error() string {
	return fmt.Sprintf("%v:%v,%v", c.Location, c.Err, c.Message)
}
