package bps2

import (
	"fmt"
	"math/rand"
	"time"
)

type Rect struct {
	x, y, width, height int
}

type Node struct {
	rect      Rect
	left      *Node
	right     *Node
	room      *Rect
	corridors []*Rect
}

func (r Rect) Center() (int, int) {
	return r.x + r.width/2, r.y + r.height/2
}

func (r Rect) Intersect(other Rect) bool {
	return r.x < other.x+other.width && r.x+r.width > other.x && r.y < other.y+other.height && r.y+r.height > other.y
}

const (
	MinRoomSize = 1
	MaxRoomSize = 2
	MinLeafSize = 10
	Dimension   = 100
)

func (n *Node) Split() bool {
	if n.left != nil || n.right != nil {
		return false
	}

	horizontal := rand.Float32() < 0.5
	if n.rect.width > n.rect.height && float32(n.rect.width)/float32(n.rect.height) >= 1.25 {
		horizontal = false
	} else if n.rect.height > n.rect.width && float32(n.rect.height)/float32(n.rect.width) >= 1.25 {
		horizontal = true
	}

	max := 0
	if horizontal {
		max = n.rect.height - MinLeafSize
	} else {
		max = n.rect.width - MinLeafSize
	}

	if max <= MinLeafSize {
		return false
	}

	split := rand.Intn(max-MinLeafSize) + MinLeafSize

	if horizontal {
		n.left = &Node{rect: Rect{x: n.rect.x, y: n.rect.y, width: n.rect.width, height: split}}
		n.right = &Node{rect: Rect{x: n.rect.x, y: n.rect.y + split, width: n.rect.width, height: n.rect.height - split}}
	} else {
		n.left = &Node{rect: Rect{x: n.rect.x, y: n.rect.y, width: split, height: n.rect.height}}
		n.right = &Node{rect: Rect{x: n.rect.x + split, y: n.rect.y, width: n.rect.width - split, height: n.rect.height}}
	}

	return true
}

func (n *Node) GenerateRooms() {
	if n.left != nil || n.right != nil {
		if n.left != nil {
			n.left.GenerateRooms()
		}
		if n.right != nil {
			n.right.GenerateRooms()
		}
		if n.left != nil && n.right != nil {
			createCorridor(n.left.getRoom(), n.right.getRoom(), n)
		}
	} else {
		w := rand.Intn(MaxRoomSize-MinRoomSize) + MinRoomSize
		h := rand.Intn(MaxRoomSize-MinRoomSize) + MinRoomSize
		x := rand.Intn(n.rect.width-w) + n.rect.x
		y := rand.Intn(n.rect.height-h) + n.rect.y
		n.room = &Rect{x: x, y: y, width: w, height: h}
	}
}

func (n *Node) getRoom() *Rect {
	if n.room != nil {
		return n.room
	}
	if n.left != nil {
		lRoom := n.left.getRoom()
		if lRoom != nil {
			return lRoom
		}
	}
	if n.right != nil {
		rRoom := n.right.getRoom()
		if rRoom != nil {
			return rRoom
		}
	}
	return nil
}

func createCorridor(r1, r2 *Rect, parent *Node) {
	x1, y1 := r1.Center()
	x2, y2 := r2.Center()

	if rand.Float32() < 0.5 {
		// horizontal then vertical
		if x1 < x2 {
			parent.corridors = append(parent.corridors, &Rect{x: x1, y: y1, width: x2 - x1, height: 1})
			if y1 < y2 {
				parent.corridors = append(parent.corridors, &Rect{x: x2, y: y1, width: 1, height: y2 - y1})
			} else {
				parent.corridors = append(parent.corridors, &Rect{x: x2, y: y2, width: 1, height: y1 - y2})
			}
		} else {
			parent.corridors = append(parent.corridors, &Rect{x: x2, y: y2, width: x1 - x2, height: 1})
			if y1 < y2 {
				parent.corridors = append(parent.corridors, &Rect{x: x1, y: y1, width: 1, height: y2 - y1})
			} else {
				parent.corridors = append(parent.corridors, &Rect{x: x1, y: y2, width: 1, height: y1 - y2})
			}
		}
	} else {
		// vertical then horizontal
		if y1 < y2 {
			parent.corridors = append(parent.corridors, &Rect{x: x1, y: y1, width: 1, height: y2 - y1})
			if x1 < x2 {
				parent.corridors = append(parent.corridors, &Rect{x: x1, y: y2, width: x2 - x1, height: 1})
			} else {
				parent.corridors = append(parent.corridors, &Rect{x: x2, y: y2, width: x1 - x2, height: 1})
			}
		} else {
			parent.corridors = append(parent.corridors, &Rect{x: x2, y: y2, width: 1, height: y1 - y2})
			if x1 < x2 {
				parent.corridors = append(parent.corridors, &Rect{x: x1, y: y1, width: x2 - x1, height: 1})
			} else {
				parent.corridors = append(parent.corridors, &Rect{x: x2, y: y1, width: x1 - x2, height: 1})
			}
		}
	}
}

func createAdditionalCorridor(r1, r2 *Rect, parent *Node) {
	x1, y1 := r1.Center()
	x2, y2 := r2.Center()

	if rand.Float32() < 0.5 {
		// horizontal then vertical
		if x1 < x2 {
			parent.corridors = append(parent.corridors, &Rect{x: x1, y: y1, width: x2 - x1, height: 1})
			if y1 < y2 {
				parent.corridors = append(parent.corridors, &Rect{x: x2, y: y1, width: 1, height: y2 - y1})
			} else {
				parent.corridors = append(parent.corridors, &Rect{x: x2, y: y2, width: 1, height: y1 - y2})
			}
		} else {
			parent.corridors = append(parent.corridors, &Rect{x: x2, y: y2, width: x1 - x2, height: 1})
			if y1 < y2 {
				parent.corridors = append(parent.corridors, &Rect{x: x1, y: y1, width: 1, height: y2 - y1})
			} else {
				parent.corridors = append(parent.corridors, &Rect{x: x1, y: y2, width: 1, height: y1 - y2})
			}
		}
	} else {
		// vertical then horizontal
		if y1 < y2 {
			parent.corridors = append(parent.corridors, &Rect{x: x1, y: y1, width: 1, height: y2 - y1})
			if x1 < x2 {
				parent.corridors = append(parent.corridors, &Rect{x: x1, y: y2, width: x2 - x1, height: 1})
			} else {
				parent.corridors = append(parent.corridors, &Rect{x: x2, y: y2, width: x1 - x2, height: 1})
			}
		} else {
			parent.corridors = append(parent.corridors, &Rect{x: x2, y: y2, width: 1, height: y1 - y2})
			if x1 < x2 {
				parent.corridors = append(parent.corridors, &Rect{x: x1, y: y1, width: x2 - x1, height: 1})
			} else {
				parent.corridors = append(parent.corridors, &Rect{x: x2, y: y1, width: x1 - x2, height: 1})
			}
		}
	}
}
func GenerateMap() [][]int {
	rand.Seed(time.Now().UnixNano())

	root := &Node{rect: Rect{x: 0, y: 0, width: Dimension, height: Dimension}}

	nodes := []*Node{root}
	leaves := []*Node{}
	for len(nodes) > 0 {
		node := nodes[0]
		nodes = nodes[1:]
		if node.Split() {
			nodes = append(nodes, node.left, node.right)
		} else {
			leaves = append(leaves, node)
		}
	}

	root.GenerateRooms()

	leaves = append(leaves, &Node{
		rect:      Rect{},
		left:      nil,
		right:     nil,
		room:      &Rect{x: 1, y: 3, width: 2, height: 2},
		corridors: nil,
	})
	// Create additional corridors to ensure every room is connected by at least two corridors
	for _, leaf := range leaves {
		if leaf.room != nil {
			// Connect to its sibling
			sibling := findSibling(root, leaf)
			if sibling != nil && sibling.room != nil {
				createAdditionalCorridor(leaf.room, sibling.room, root)
			}

			// Connect to another random room
			if len(leaves) > 1 {
				var other *Node
				for {
					other = leaves[rand.Intn(len(leaves))]
					if other != leaf && other.room != nil {
						break
					}
				}
				createAdditionalCorridor(leaf.room, other.room, root)
			}
		}
	}

	// Create a grid to represent the map
	grid := make([][]int, Dimension)
	for i := range grid {
		grid[i] = make([]int, Dimension)
		for j := range grid[i] {
			grid[i][j] = 0
		}
	}

	// Fill the grid with rooms
	for _, leaf := range leaves {
		if leaf.room != nil {
			for y := leaf.room.y; y < leaf.room.y+leaf.room.height; y++ {
				for x := leaf.room.x; x < leaf.room.x+leaf.room.width; x++ {
					grid[y][x] = 1
				}
			}
		}
	}

	// Fill the grid with corridors
	printCorridors(root, grid)

	// Print the grid
	for _, row := range grid {
		for _, cell := range row {
			fmt.Print(fmt.Sprintf("%d", cell))
		}
		fmt.Println()
	}

	return grid

}

func printCorridors(node *Node, grid [][]int) {
	if node == nil {
		return
	}
	for _, corridor := range node.corridors {
		for y := corridor.y; y < corridor.y+corridor.height; y++ {
			for x := corridor.x; x < corridor.x+corridor.width; x++ {
				grid[y][x] = 1
			}
		}
	}
	printCorridors(node.right, grid)
	printCorridors(node.left, grid)
}

func findSibling(parent, child *Node) *Node {
	if parent.left == child {
		return parent.right
	} else if parent.right == child {
		return parent.left
	} else {
		if parent.left != nil {
			sibling := findSibling(parent.left, child)
			if sibling != nil {
				return sibling
			}
		}
		if parent.right != nil {
			sibling := findSibling(parent.right, child)
			if sibling != nil {
				return sibling
			}
		}
	}
	return nil
}
