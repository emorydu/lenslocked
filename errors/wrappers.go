package errors

import "errors"

// These variables are used to give us access to existing
// functions in the std lib errors package. We can also
// wrap them in custom functionality as needed if we want,
// or mock them during testing.
var (
	Is = errors.Is
	As = errors.As
)

//func As(err error, target any) bool {
//	return errors.As(err, target)
//}
