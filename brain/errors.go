package brain

import "errors"

// ErrNotFound means that tried to find the key but no result found
var ErrNotFound = errors.New("not found")
