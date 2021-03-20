package error

import "fmt"

type PodError struct {
	Code     int
	Err      string
	Message  string
	Location string
}

func (c *PodError) Error() string {
	return fmt.Sprintf("%v:%v,%v", c.Location, c.Err, c.Message)
}
