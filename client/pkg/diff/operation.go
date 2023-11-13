package diff

const (
	INSERTION = iota
	DELETION
	SUBSTITUTION
)

type Operation struct {
	Type      int  `json:"type"`
	Line      int  `json:"line"`
	Position  int  `json:"position"`
	Character byte `json:"character"`
}
