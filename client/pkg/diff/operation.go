package diff

import (
	"os"
)

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

func (o Operation) Apply(path string) error {
	var err error

	if o.Type == INSERTION {
		err = insert(path, o)
	} else if o.Type == DELETION {
		// TODO: Implement
	} else if o.Type == SUBSTITUTION {
		// TODO: Implement
	}

	return err
}

// NOTE: Very inefficient for large files
func insert(path string, operation Operation) error {
	content, _ := os.ReadFile(path)

	newContent := make([]byte, len(content)+1)
	copy(newContent[:operation.Position], content[:operation.Position])
	newContent[operation.Position] = operation.Character
	copy(newContent[operation.Position+1:], content[operation.Position:])

	return os.WriteFile(path, newContent, os.ModeAppend)
}
