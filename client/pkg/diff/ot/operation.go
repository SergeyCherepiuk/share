package ot

const (
	INSERTION = iota
	DELETION
	SUBSTITUTION
)

type Operation struct {
	Type      int  `json:"type"`
	Position  int  `json:"position"`
	Character byte `json:"character"`
}
