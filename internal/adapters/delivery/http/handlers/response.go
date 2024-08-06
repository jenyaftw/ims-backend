package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jenyaftw/scaffold-go/internal/core/domain"
)

var errorStatusMap = map[error]int{
	domain.ErrInternal:        http.StatusInternalServerError,
	domain.ErrDataNotFound:    http.StatusNotFound,
	domain.ErrDataConflict:    http.StatusConflict,
	domain.ErrUnauthorized:    http.StatusUnauthorized,
	domain.ErrForbidden:       http.StatusForbidden,
	domain.ErrInvalidPassword: http.StatusUnauthorized,

	domain.ErrMissingAuthHeader: http.StatusUnauthorized,
	domain.ErrInvalidAuthToken:  http.StatusUnauthorized,
	domain.ErrInvalidTokenType:  http.StatusUnauthorized,

	domain.ErrUserNotFound: http.StatusNotFound,
}

func validationError(w http.ResponseWriter, err error) {
	errMsgs := parseError(err)
	errRsp := newErrorResponse(errMsgs)

	res, err := json.Marshal(errRsp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Error(w, string(res), http.StatusBadRequest)
}

func HandleError(w http.ResponseWriter, err error) {
	statusCode, ok := errorStatusMap[err]
	if !ok {
		statusCode = http.StatusInternalServerError
	}

	errMsg := parseError(err)
	errRsp := newErrorResponse(errMsg)

	res, err := json.Marshal(errRsp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Error(w, string(res), statusCode)
}

func parseError(err error) []string {
	var errMsgs []string

	if errors.As(err, &validator.ValidationErrors{}) {
		for _, err := range err.(validator.ValidationErrors) {
			errMsgs = append(errMsgs, err.Error())
		}
	} else {
		errMsgs = append(errMsgs, err.Error())
	}

	return errMsgs
}

type errorResponse struct {
	Success  bool     `json:"success"`
	Messages []string `json:"messages"`
}

func newErrorResponse(errMsgs []string) errorResponse {
	return errorResponse{
		Success:  false,
		Messages: errMsgs,
	}
}

type userResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func newUserResponse(user domain.User) userResponse {
	return userResponse{
		Name:  user.Name,
		Email: user.Email,
	}
}

type tokenResponse struct {
	AccessToken string `json:"accessToken"`
	ExpiresAt   int64  `json:"expiresAt"`
}

func newTokenResponse(token domain.Token) tokenResponse {
	return tokenResponse{
		AccessToken: token.Text,
		ExpiresAt:   token.ExpiresAt,
	}
}
