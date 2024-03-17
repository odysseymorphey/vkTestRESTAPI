package dto

type Movie struct {
	MovieID     string
	title       string
	description string
	release     string
	rating      int
	actors      []Actor
}
