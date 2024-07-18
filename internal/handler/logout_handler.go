package handler

import (
	"net/http"
	"time"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{Name: "forum_session", Value: "", Expires: time.Now()}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
