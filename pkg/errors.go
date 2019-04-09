package pkg

// Define your Error struct
type Error struct {
	msg string
}
// Create a function Error() string and associate it to the struct.
func (err *Error) Error() string {
	return err.msg
}
// Now you can construct an error object using MyError struct.
func Throw(err string) error {
	return &Error{err}
}

