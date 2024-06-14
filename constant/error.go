package constant

import "errors"

var ErrInsertDatabase error = errors.New("invalid Add Data in Database")
var ErrInvalidRequest error = errors.New("invalid Request")
var ErrEmptyInput error = errors.New("input cannot be empty")
var ErrDuplicatedData error = errors.New("duplicated data")
var ErrNotFound error = errors.New("not found")
var ErrInvalidEmailOrPassword error = errors.New("invalid email or password")
var ErrNotAuthorized error = errors.New("not authorized")
var ErrFailedUpdate error = errors.New("failed to update the data")
var ErrTokenNotFound error = errors.New("token not found")
var ErrTokenNotValid error = errors.New("token not valid")
var ErrPaymentGateway error = errors.New("payment gateway error")
var ErrProductStatusValue = errors.New("product status value available or unavailable")
var ErrOpenAICallAPI = errors.New("open ai call api error")
