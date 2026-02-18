package handlers

import (
	"net/http"

	"github.com/f18charles/expense-tracker/internal/auth"
	"github.com/f18charles/expense-tracker/internal/models"
	"github.com/f18charles/expense-tracker/internal/utils"
	"gorm.io/gorm"
)

func SignUp(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			utils.RenderTemplate(w, "signup.html", nil)
			return
		}

		name := r.FormValue("name")
		email := r.FormValue("email")
		password := r.FormValue("password")

		hash, err := auth.HashPass(password)
		if err != nil {
			http.Error(w, "Failed to hash password", http.StatusInternalServerError)
			return 
		}

		user := models.User {
			Name: name,
			Email: email,
			PassHash: string(hash),
		}

		if err := db.Create(&user).Error; err != nil {
			utils.ErrorHandler(w,r,http.StatusBadRequest,"Email already Exists")
			return
		}

		utils.SetSession(w, user.ID)
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	}
}

func SignIn(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			utils.RenderTemplate(w,"signin.html",nil)
			return
		}

		email := r.FormValue("email")
		password := r.FormValue("password")

		var user models.User

		if err := db.Where("email = ?", email).First(&user).Error; err != nil {
			http.Error(w, "Wrong email", http.StatusUnauthorized)
		}

		if !auth.CheckPass(password, user.PassHash) {
			http.Error(w, "Wrong password", http.StatusUnauthorized)
			return
		}

		utils.SetSession(w, user.ID)
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	}
}

func SignOut() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := &http.Cookie {
			Name: "user_id",
			Value: "",
			Path: "/",
			MaxAge: -1,
		}
		http.SetCookie(w, c)
		http.Redirect(w, r, "/signin", http.StatusSeeOther)
	}
}