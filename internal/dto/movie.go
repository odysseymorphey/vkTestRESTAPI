package dto

type Movie struct {
	MovieID     string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Release     string  `json:"release"`
	Rating      int     `json:"rating"`
	Actors      []Actor `json:"actors"`
}
