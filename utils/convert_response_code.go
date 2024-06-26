package utils

import (
	"errors"
	"net/http"

	"github.com/blueharvest-alterra/go-back-end/constant"
)

func ConvertResponseCode(err error) int {
	switch {
	case errors.Is(err, constant.ErrInsertDatabase):
		return http.StatusInternalServerError
	case errors.Is(err, constant.ErrEmptyInput):
		return http.StatusBadRequest
	case errors.Is(err, constant.ErrDuplicatedData):
		return http.StatusBadRequest
	case errors.Is(err, constant.ErrNotFound):
		return http.StatusNotFound
	case errors.Is(err, constant.ErrInvalidEmailOrPassword):
		return http.StatusUnauthorized
	case errors.Is(err, constant.ErrNotAuthorized):
		return http.StatusUnauthorized
	case errors.Is(err, constant.ErrFailedUpdate):
		return http.StatusInternalServerError
	case errors.Is(err, constant.ErrTokenNotFound):
		return http.StatusForbidden
	case errors.Is(err, constant.ErrTokenNotValid):
		return http.StatusForbidden
	case errors.Is(err, constant.ErrProductStatusValue):
		return http.StatusBadRequest
	case errors.Is(err, constant.ErrCustomerNotFound):
		return http.StatusNotFound
	case errors.Is(err, constant.ErrPromoNotFound):
		return http.StatusNotFound
	case errors.Is(err, constant.ErrCustomerAddressNotFound):
		return http.StatusNotFound
	case errors.Is(err, constant.ErrProductNotFound):
		return http.StatusNotFound
	case errors.Is(err, constant.ErrPromoUnavailable):
		return http.StatusBadRequest
	case errors.Is(err, constant.ErrProductUnavailable):
		return http.StatusBadRequest
	case errors.Is(err, constant.ErrMinumumAmount):
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
