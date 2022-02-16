package domain

type GameLog interface {
	Start(*Game) string
	Attack(*Game) string
	Finish(*Game) string
}
