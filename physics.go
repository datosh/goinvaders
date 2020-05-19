package spaceinvaders

type Rect struct {
	x, y, w, h float64
}

func DoCollide(e1, e2 Rect) bool {
	// AABB collision test
	if e1.x < e2.x+e2.w &&
		e1.x+e1.w > e2.x &&
		e1.y < e2.y+e2.h &&
		e1.y+e1.h > e2.y {
		return true
	}
	return false
}

func In(inside, outside Rect) bool {
	return inside.x <= outside.x &&
		(inside.x+inside.w) <= (outside.x+outside.w) &&
		inside.y <= outside.y &&
		(inside.y+inside.h) <= (outside.y+outside.h)
}
