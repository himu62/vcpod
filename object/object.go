package object

const (
	TypeDirectory = iota
	TypeBook
	TypeVideo
	TypeFile
)

type Object struct {
	ID       uint
	Path     string
	Type     int
	Hash     string
	HashType string
}
