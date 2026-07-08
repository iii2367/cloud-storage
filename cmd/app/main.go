package main

import (
	"cloud-storage/internal/auth"
	"cloud-storage/internal/auth/postgres"
	"cloud-storage/internal/config"
	"cloud-storage/internal/database"
	"context"
	"time"

	//"github.com/jackc/pgx/v5/pgxpool"
	"fmt"
	//"context"
)

func main() {
	cfg := config.MustLoad("MAIN")

	fmt.Print(cfg.Server.Host + ":" + cfg.Server.Port + "\n" + cfg.Database.Host + ":" + cfg.Database.Port + "\n")

	/*db, err := database.NewPostgres(&cfg.Database)	
	if err != nil {
		return
	}
	db.Exec(
		context.Background(),
		`
		INSERT INTO users (name, email, password_hash) VALUES ($1, $2, $3)`,
		"Test", 
		"Text@",
		"bjhdkjnkdjnkedkbegbjbvghvg",
		)*/

	/*mgjwt := jwt.New(cfg.JWT)
	jt, err := mgjwt.Generate(123)
	if err != nil {
		return
	}
	fmt.Println(jt)
	sl, err := mgjwt.Parse(jt)
	if err != nil {
		return
	}
	fmt.Println(sl.ID)*/

	var user auth.User
	user.Name = "IiI"
	user.Email = "tesst@gmail.com"
	user.PasswordHash = "1234567890"
	db, err := database.New(&cfg.Database)
	if err != nil {
		return
	}
	repo := postgres.New(db)	
	erro := repo.Create(context.Background(), &user)
	if erro != nil {
		return
	}
	usr, erro := repo.FindByEmail(context.Background(), "test@gmail.com")
	if erro != nil {
		return
	}
	fmt.Printf("ID: %d\nName: %s\nEmail: %s\nPasswordHash: %s\nCreateAt: %s\n", usr.ID, usr.Name, usr.Email, usr.PasswordHash, usr.CreatedAt.Format(time.RFC3339))
	usrw, errow := repo.FindByID(context.Background(), 37)
	if errow != nil {
		return
	}
	fmt.Printf("ID: %d\nName: %s\nEmail: %s\nPasswordHash: %s\nCreateAt: %s\n", usrw.ID, usrw.Name, usrw.Email, usrw.PasswordHash, usrw.CreatedAt.Format(time.RFC3339))

		
}
