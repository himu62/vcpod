package object

import "time"

type Type string

const (
	Directory Type = "directory"
	Book           = "book"
	Video          = "video"
	File           = "file"
)

type (
	Object struct {
		ID   uint
		Path string
		Type Type

		Hash     string
		HashType string

		Registered time.Time
		LastUsed   time.Time
		Used       []time.Time
		Score      int
		FileSize   uint64
	}
	Objects []*Object
)
