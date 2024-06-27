package models


	import (
		"golang.org/x/crypto/bcrypt"
		"go.mongodb.org/mongo-driver/bson/primitive"
	)

	type User struct {
		ID primitive.ObjectID `bson:"_id,omitempty"`
		username string   `bson:"username"`
		password string  `bson:"password"`
	}

	func (u *User) HashPassword(password string) error{
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.password= string(hashedPassword)
	}