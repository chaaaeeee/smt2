package main

type User struct {
	Username string `json:"username"` 
	Password string `json:"password"`
}

type UserPasswordChange struct {
	Username string `json:"username"`
	Password string `json:"password"`
	NewPassword string `json:"newPassword"`
}

type Movies struct {

}
