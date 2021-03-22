package models

type Link struct {
	ID    string `json:"_id"`
	Alias string `json:"alias"`
	URL   string `json:"url"`
}
