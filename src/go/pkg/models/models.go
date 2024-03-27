package models

type isTheorem int

const (
	Unset isTheorem = iota
	True
	False
)

type Formula struct {
	Expression       string
	IsClassicTh      isTheorem
	IsIntuitionistTh isTheorem
}
