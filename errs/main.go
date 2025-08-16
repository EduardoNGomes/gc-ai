package errs

import "encoding/json"

const CannotOpenFileErr = "Can not open file"

const UnexpectedError = "Unexpected error"

const ErrorOnConvertDataToByte = "Error on convert data to byte"

const ErrorOnWriteFileConfig = "Error on write file config"

var InvalidJson *json.UnmarshalTypeError
