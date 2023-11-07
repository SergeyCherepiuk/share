package ot

type operation interface{}

type insertion struct {
	Position  int  `json:"position"`
	Character byte `json:"character"`
}

type deletion struct {
	Position int `json:"position"`
}

type substitution struct {
	Position  int  `json:"position"`
	Character byte `json:"character"`
}
