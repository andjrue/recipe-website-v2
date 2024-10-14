package main

type User struct {
	Email    string `bson:"email"`
	Username string `bson:"username"`
	Password string `bson:"password"`
	// Recipes []Recipe -- Will add this later
}

func newUser(email, username, password string) *User {
	return &User{
		Email:    email,
		Username: username,
		Password: password,
	}
}
