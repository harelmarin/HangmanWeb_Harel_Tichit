package main

import (
	routeur "hangman/routeur"
	initTemplate "hangman/templates"
)

func main() {
	initTemplate.InitTemplate()
	routeur.InitServer()
}
