package main

func main() {

}

func hasCross(orders string) bool {
	pm := map[[2]int]struct{}{}

	current := [2]int{0, 0}
	pm[current] = struct{}{}
	for _, o := range orders {
		np := nextPos(current, string(o))
		if _, found := pm[np]; found {
			return true
		}
		pm[np] = struct{}{}
	}
	return false
}

func nextPos(current [2]int, order string) [2]int {
	switch order {
	case "E":
		current[0] += 1
	case "S":
		current[1] -= 1
	case "W":
		current[0] -= 1
	case "N":
		current[1] += 1
	}
	return current
}

func hasCross2(orders string) bool {
	pm := map[pos]struct{}{}
	p := &pos{0, 0}
	pm[*p] = struct{}{}
	for _, o := range orders {
		p.Move(string(o))
		if _, found := pm[*p]; found {
			return true
		}
		pm[*p] = struct{}{}
	}
	return false
}

type pos struct {
	x int
	y int
}

func (p *pos) Equals(p2 *pos) bool {
	return p.x == p2.x && p.y == p2.y
}

func (p *pos) Move(order string) {
	switch order {
	case "E":
		p.x += 1
	case "S":
		p.y -= 1
	case "W":
		p.x -= 1
	case "N":
		p.y += 1
	}
}
