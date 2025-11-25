package errs

import (
	"encoding/json"
	"errors"
)

const CannotOpenFileErr = "Can not open file"

const UnexpectedError = "Unexpected error"

const ErrorOnConvertDataToByte = "Error on convert data to byte"

const ErrorOnWriteFileConfig = "Error on write file config"

var EmptyKeyError = errors.New("Cannot make commit without AI Key")

var EmptyDiffError = errors.New("Cannot make commit without any diff")

var InvalidAgentSelected = errors.New("Invalid Agent selected")

var InvalidJson *json.UnmarshalTypeError
