package utils

import (
	"net/http"
	"strconv"
)

// SetSession sets a simple cookie-based session containing the numeric user ID.
// Used by server-rendered flows; API clients typically use JWT instead.
func SetSession(w http.ResponseWriter, userID uint) {
	c := &http.Cookie{
		Name:     "user_id",
		Value:    strconv.Itoa(int(userID)),
		Path:     "/",
		HttpOnly: true,
		MaxAge:   86400,
	}
	http.SetCookie(w, c)
}
