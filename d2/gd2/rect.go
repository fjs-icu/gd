package gd2

type Rect struct {
	X, Y, W, H int
}

func (c Rect) Left() int {
	return c.X
}

func (c Rect) Top() int {
	return c.Y
}

func (c Rect) Right() int {
	return c.X + c.W - 1
}

func (c Rect) Bottom() int {
	return c.Y + c.H - 1
}

func (c *Rect) IsZero() bool {
	return c.X == 0 && c.Y == 0 && c.W == 0 && c.H == 0
}

type Size struct {
	Cx, Cy int32
}
