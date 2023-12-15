package platform

type Platform struct {
	matrix [][]Space
}

func (p *Platform) Copy() *Platform {
	var matrix [][]Space
	for _, row := range p.matrix {
		var nrow []Space
		for _, space := range row {
			nrow = append(nrow, space)
		}
		matrix = append(matrix, nrow)
	}
	return &Platform{matrix: matrix}
}

func (p *Platform) Equal(p2 *Platform) bool {
	for y, row := range p.matrix {
		for x, space := range row {
			if space != p2.matrix[y][x] {
				return false
			}
		}
	}
	return true
}

func (p *Platform) String() string {
	var str string
	for _, row := range p.matrix {
		for _, space := range row {
			str += string(space)
		}
		str += "\n"
	}
	return str
}

func (p *Platform) Get(x, y int) Space {
	return p.matrix[y][x]
}

func (p *Platform) AddRow(row []Space) {
	p.matrix = append(p.matrix, row)
}

func (p *Platform) MoveUp(x, y int) {
	space := p.Get(x, y)
	if space == RoundedRock {
		for i := y - 1; i >= 0; i-- {
			mspace := p.Get(x, i)
			if mspace == Empty {
				p.matrix[i][x] = RoundedRock
				p.matrix[i+1][x] = Empty
			} else {
				break
			}
		}
	}
}

func (p *Platform) MoveDown(x, y int) {
	space := p.Get(x, y)
	if space == RoundedRock {
		for i := y + 1; i < len(p.matrix); i++ {
			mspace := p.Get(x, i)
			if mspace == Empty {
				p.matrix[i][x] = RoundedRock
				p.matrix[i-1][x] = Empty
			} else {
				break
			}
		}
	}
}

func (p *Platform) MoveLeft(x, y int) {
	space := p.Get(x, y)
	if space == RoundedRock {
		for i := x - 1; i >= 0; i-- {
			mspace := p.Get(i, y)
			if mspace == Empty {
				p.matrix[y][i] = RoundedRock
				p.matrix[y][i+1] = Empty
			} else {
				break
			}
		}
	}
}

func (p *Platform) MoveRight(x, y int) {
	space := p.Get(x, y)
	if space == RoundedRock {
		for i := x + 1; i < len(p.matrix[y]); i++ {
			mspace := p.Get(i, y)
			if mspace == Empty {
				p.matrix[y][i] = RoundedRock
				p.matrix[y][i-1] = Empty
			} else {
				break
			}
		}
	}
}

func (p *Platform) TiltUp() {
	for y := 0; y < len(p.matrix); y++ {
		for x := 0; x < len(p.matrix[y]); x++ {
			p.MoveUp(x, y)
		}
	}
}

func (p *Platform) TiltDown() {
	for y := len(p.matrix) - 1; y >= 0; y-- {
		for x := 0; x < len(p.matrix[y]); x++ {
			p.MoveDown(x, y)
		}
	}
}

func (p *Platform) TiltLeft() {
	for x := 0; x < len(p.matrix[0]); x++ {
		for y := 0; y < len(p.matrix); y++ {
			p.MoveLeft(x, y)
		}
	}
}

func (p *Platform) TiltRight() {
	for x := len(p.matrix[0]) - 1; x >= 0; x-- {
		for y := 0; y < len(p.matrix); y++ {
			p.MoveRight(x, y)
		}
	}
}

func (p *Platform) SpinCycle() {
	p.TiltUp()
	p.TiltLeft()
	p.TiltDown()
	p.TiltRight()
}

func (p *Platform) TotalLoad() int {
	var total int
	for i, row := range p.matrix {
		for _, space := range row {
			if space == RoundedRock {
				total += len(p.matrix) - i
			}
		}
	}
	return total
}
