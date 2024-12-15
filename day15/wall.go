package main

type Wall struct{}

func (_ *Wall) CharFor(_ int) string {
	return "#"
}

func (_ *Wall) CanMove(Map, rune) bool {
	return false
}

func (_ *Wall) Move(Map, rune) {}
