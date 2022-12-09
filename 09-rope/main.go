package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	visited := make(map[position]struct{})
	rope := make([]position, 10)

	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		var (
			moveTo   direction
			distance int
		)

		_, err := fmt.Sscanf(s.Text(), "%c %d", &moveTo, &distance)
		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i < distance; i++ {
			rope = moveRope(rope, moveTo)
			visited[rope[len(rope)-1]] = struct{}{}
		}
	}

	fmt.Println(len(visited))
}

func moveRope(rope []position, moveTo direction) []position {
	newRope := make([]position, len(rope))
	copy(newRope, rope)

	newRope[0] = newRope[0].move(moveTo)

	for i := 1; i < len(newRope); i++ {
		desiredPosition := tailPosition(newRope[i-1], newRope[i])
		if desiredPosition == newRope[i] {
			break
		}

		newRope[i] = desiredPosition
	}

	return newRope
}

type position struct {
	x, y int
}

type direction rune

const (
	Up    direction = 'U'
	Down  direction = 'D'
	Left  direction = 'L'
	Right direction = 'R'
)

func tailPosition(head, tail position) position {
	if head.touches(tail) {
		return tail
	}

	switch head {
	case tail.top().top():
		return tail.top()
	case tail.top().top().left():
		return tail.top().left()
	case tail.top().top().left().left():
		return tail.top().left()
	case tail.top().top().right():
		return tail.top().right()
	case tail.top().top().right().right():
		return tail.top().right()
	case tail.top().right().right():
		return tail.top().right()
	case tail.right().right():
		return tail.right()
	case tail.right().right().bottom():
		return tail.right().bottom()
	case tail.right().right().bottom().bottom():
		return tail.right().bottom()
	case tail.right().bottom().bottom():
		return tail.right().bottom()
	case tail.bottom().bottom():
		return tail.bottom()
	case tail.left().bottom().bottom():
		return tail.left().bottom()
	case tail.left().left().bottom().bottom():
		return tail.left().bottom()
	case tail.left().left().bottom():
		return tail.left().bottom()
	case tail.left().left():
		return tail.left()
	case tail.left().left().top():
		return tail.left().top()
	}
	panic(fmt.Sprintf("h: %+v t: %+v", head, tail))
}

func (p position) touches(other position) bool {
	if p == other {
		return true
	} else if p.top() == other {
		return true
	} else if p.top().right() == other {
		return true
	} else if p.right() == other {
		return true
	} else if p.right().bottom() == other {
		return true
	} else if p.bottom() == other {
		return true
	} else if p.bottom().left() == other {
		return true
	} else if p.left() == other {
		return true
	} else if p.left().top() == other {
		return true
	}
	return false
}

func (p position) top() position {
	return position{p.x, p.y - 1}
}

func (p position) bottom() position {
	return position{p.x, p.y + 1}
}

func (p position) left() position {
	return position{p.x - 1, p.y}
}

func (p position) right() position {
	return position{p.x + 1, p.y}
}

func (p position) move(to direction) position {
	switch to {
	case Left:
		return p.left()
	case Right:
		return p.right()
	case Up:
		return p.top()
	case Down:
		return p.bottom()
	}
	panic("unknown direction")
}
