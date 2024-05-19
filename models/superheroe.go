package models

type (
	Superheroe struct {
		Nombre    string `json:"name"`
		Biography struct {
			FullName string `json:"fullName"`
		} `json:"biography"`
		Powerstats struct {
			Intelligence int `json:"intelligence"`
			Strength     int `json:"strength"`
			Speed        int `json:"speed"`
			Durability   int `json:"durability"`
			Power        int `json:"power"`
			Combat       int `json:"combat"`
		} `json:"powerstats"`
		Images struct {
			Xs string `json:"xs"`
			Sm string `json:"sm"`
			Md string `json:"md"`
			Lg string `json:"lg"`
		} `json:"images"`
	}
)
