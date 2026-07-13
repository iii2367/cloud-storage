package http

import (
    "net/http"
    "time"
)

const refreshCookieName = "refresh_token"

func setRefreshCookie(w http.ResponseWriter, token string, expiresAt time.Time) {
    http.SetCookie(w, &http.Cookie{
        Name:     refreshCookieName,
        Value:    token,
        Path:     "/",
        HttpOnly: true,
        Secure:   false, // true in production
        SameSite: http.SameSiteLaxMode,
        Expires:  expiresAt,
    })
}

func clearRefreshCookie(w http.ResponseWriter) {
    http.SetCookie(w, &http.Cookie{
        Name:     refreshCookieName,
        Value:    "",
        Path:     "/",
        HttpOnly: true,
        Secure:   false,
        SameSite: http.SameSiteLaxMode,
        MaxAge:   -1,
    })
}
