package utils

import (
	"net/http"
	"strconv"
)

func SetSession(w http.ResponseWriter, userID uint) {
	c := &http.Cookie{
		Name: "user_id",
		Value: strconv.Itoa(int(userID)),
		Path: "/",
		HttpOnly: true,
	}
	http.SetCookie(w,c)
}