package middleware

import (
	"errors"
	"net/http"
	"strconv"
)

func GetUserID(r *http.Request) (uint, error) {
	cookie, err := r.Cookie("user_id")
	if err != nil {
		return 0, errors.New("not Authenticated")
	}
	id, err := strconv.Atoi(cookie.Value)
	if err != nil {
		return 0, errors.New("Invalid user id")
	}
	return uint(id), nil
}