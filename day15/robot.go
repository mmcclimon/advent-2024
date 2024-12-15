package main

type Robot struct {
	x, y int
}

func (r *Robot) xy() xy { return xy{r.x, r.y} }

func (_ *Robot) CharFor(_ int) string {
	return "@"
}

func (r *Robot) CanMove(grid Map, dir rune) bool {
	return canMoveSingle(grid, r, dir)
}

func (r *Robot) Move(grid Map, dir rune) {
	if !r.CanMove(grid, dir) {
		return
	}

	next := moveSingle(grid, r, dir)
	r.x = next.x
	r.y = next.y
}
