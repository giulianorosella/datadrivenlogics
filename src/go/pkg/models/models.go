package models

type isTheorem int

const (
	Unset isTheorem = iota - 1
	True
	False
)

type Formula struct {
	Expression       string
	IsClassicTh      isTheorem
	IsIntuitionistTh isTheorem
}
