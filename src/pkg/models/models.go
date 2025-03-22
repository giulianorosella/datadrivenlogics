package models

type IsTheorem int

const (
	Unset IsTheorem = iota - 1
	False
	True
)

type Formula struct {
	Expression       string
	IsClassicTh      IsTheorem
	IsIntuitionistTh IsTheorem
}
