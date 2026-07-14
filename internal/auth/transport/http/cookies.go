package http

import (
    "net/http"
    "time"
)

const refreshCookieName = "refresh_token"

func setRefreshCookie(w http.ResponseWriter,  r *http.Request, token string, expiresAt time.Time) {
	secure := r.TLS != nil
    http.SetCookie(w, &http.Cookie{
        Name:     refreshCookieName,
        Value:    token,
        Path:     "/",
        HttpOnly: true,
        Secure:   secure, 
        SameSite: http.SameSiteLaxMode,
        Expires:  expiresAt,
    })
}

func clearRefreshCookie(w http.ResponseWriter,  r *http.Request) {
	secure := r.TLS != nil
    http.SetCookie(w, &http.Cookie{
        Name:     refreshCookieName,
        Value:    "",
        Path:     "/",
        HttpOnly: true,
        Secure:   secure,
        SameSite: http.SameSiteLaxMode,
        MaxAge:   -1,
    })
}
