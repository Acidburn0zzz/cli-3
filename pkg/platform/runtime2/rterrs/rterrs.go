// Package rterrs contains runtime error types. This may be a temporary package
// and it may be useful into the future.
package rterrs

import "errors"

var NotImplemented = errors.New("not implemented")
