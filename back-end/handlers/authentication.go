package handlers

import (
	"fmt"
	"net/http"
	"taimekalender/back-end/helpers"
	"taimekalender/back-end/models"
	"taimekalender/back-end/repository"
	"taimekalender/back-end/sessions"
	"time"
)

func Authenticate(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("taimekalender-session")
	if err != nil {
		_ = helpers.ErrorJson(w, fmt.Errorf("cookie not found"), http.StatusUnauthorized)
		return
	}
	authenticated, err := repository.AuthenticateViaCookie(cookie.Value)
	if !authenticated || err != nil {
		_ = helpers.ErrorJson(w, fmt.Errorf("user not found"), http.StatusUnauthorized)
		return
	}

	_ = helpers.WriteJson(w, http.StatusOK, nil, nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	_ = helpers.ReadJson(w, r, &user)

	userID, authenticated, err := repository.AuthenticateViaPassword(user.Username, user.Password)
	if !authenticated || err != nil {
		_ = helpers.ErrorJson(w, fmt.Errorf("invalid credentials %v", err), http.StatusUnauthorized)
		return
	}

	uuid, err := sessions.Set(userID)
	if err != nil {
		_ = helpers.ErrorJson(w, fmt.Errorf("cannot create session: %v", err), http.StatusServiceUnavailable)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "taimekalender-session",
		Value:   uuid,
		Expires: time.Now().Add(240 * time.Hour),
	})

	_ = helpers.WriteJson(w, http.StatusOK, nil, nil)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("taimekalender-session")
	if err != nil {
		_ = helpers.ErrorJson(w, fmt.Errorf("cookie not found: %v", err), http.StatusUnauthorized)
		return
	}

	sessions.Remove(cookie.Value)
	cookie.Expires = time.Now().Add(-1 * time.Hour)
	http.SetCookie(w, cookie)

	_ = helpers.WriteJson(w, http.StatusOK, nil, nil)
}
