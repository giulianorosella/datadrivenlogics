package models

type isTheorem int

const (
	Unset isTheorem = iota - 1
	False
	True
)

type Formula struct {
	Expression       string
	IsClassicTh      isTheorem
	IsIntuitionistTh isTheorem
}
