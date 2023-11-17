package controller

var DataUser User = User{"nul"}
var Nom Name
var Diff Difficulty

type Name struct {
	Username string
}

type Difficulty struct {
	Difficulte string
}
type User struct {
	Word string
}
