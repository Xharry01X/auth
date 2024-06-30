package database

import (
	 "database/sql"
    "github.com/harshit-1245/auth/backend/models"
)

func CreateUser(db *sql.DB,user *models.User) error {
	_, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)",user.Username,user.Password)
	return err
}

func FindUserByUsername(db *sql.DB,username string)(*models.User, error){
	var user models.User
    err := db.QueryRow("SELECT username, password FROM users WHERE username = ?", username).Scan(&user.Username, &user.Password)
    if err != nil {
        return nil, err
    }
    return &user, nil
}