package errs

import "encoding/json"

const CannotOpenFileErr = "Can not open file"

const UnexpectedError = "Unexpected error"

var InvalidJson *json.UnmarshalTypeError
