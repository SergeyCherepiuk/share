package ot

type Operation interface{}

type Insertion struct {
	Position  int
	Character byte
}

type Deletion struct {
	Position int
}

type Substitution struct {
	Position  int
	Character byte
}
