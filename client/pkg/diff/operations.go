package diff

type Operation interface{}

type Insertion struct {
	Line      int  `json:"line"`
	Position  int  `json:"position"`
	Character byte `json:"character"`
}

type Deletion struct {
	Line     int `json:"line"`
	Position int `json:"position"`
}

type Substitution struct {
	Line      int  `json:"line"`
	Position  int  `json:"position"`
	Character byte `json:"character"`
}
