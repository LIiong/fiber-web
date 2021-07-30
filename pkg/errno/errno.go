package errno

import "fmt"

type Err struct {
	Code    int
	Message string
	Err     error
}

func (err *Err) Add(message string) error {
	err.Message += " " + message
	return err
}

func (err *Err) Addf(format string, args ...interface{}) error {
	err.Message += " " + fmt.Sprintf(format, args...)
	return err
}

func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Err)
}

func DecodeErr(err error) (int, string) {
	if err == nil {
		return Ok.Code, Ok.Message
	}

	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message
	default:
	}

	return InternalServerError.Code, err.Error()
}
