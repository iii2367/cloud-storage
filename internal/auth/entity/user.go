package entity

import ("time")

type User struct {
    ID           uint
    Name         string
    Email        string
    PasswordHash string
	CreatedAt    time.Time
}
