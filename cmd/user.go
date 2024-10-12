package main

type User struct {
    
    Email string
    Username string
    Password string
    // Recipes []Recipe -- Will add this later
}

func newUser(email, username, password string) *User {
    return &User{
        Email: email,
        Username: username,
        Password: password,
    }
}


