package utility

import (
	"city_game/src/pb"
	"container/heap"
	"slices"
)

type Cord [2]int

func (c Cord) Up() Cord    { return Cord{c[0], c[1] - 1} }
func (c Cord) Down() Cord  { return Cord{c[0], c[1] + 1} }
func (c Cord) Left() Cord  { return Cord{c[0] - 1, c[1]} }
func (c Cord) Right() Cord { return Cord{c[0] + 1, c[1]} }
func (c Cord) InBounds(w int, h int) bool {
	return c[0] >= 0 && c[0] < w && c[1] >= 0 && c[1] < h
}

type WeightedCord struct {
	c Cord
	g int
	w int
}
type CordWeightedQueue []*WeightedCord

func (q CordWeightedQueue) Len() int { return len(q) }

func (q CordWeightedQueue) Less(i, j int) bool {
	return q[i].w < q[j].w
}

func (q CordWeightedQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *CordWeightedQueue) Push(c any) {
	cord := c.(*WeightedCord)
	*q = append(*q, cord)
}

func (q *CordWeightedQueue) Pop() any {
	old := *q
	n := len(old)
	node := old[n-1]
	old[n-1] = nil
	*q = old[0 : n-1]
	return node
}

func FindShortestPath(grid [][]int, start_v *pb.Coordinate, end_v *pb.Coordinate) []*pb.Coordinate {
	height := len(grid)
	width := len(grid[0])

	start := Cord{int(start_v.X), int(start_v.Y)}
	end := Cord{int(end_v.X), int(end_v.Y)}

	h := func(a Cord) int { return distance_pathfinder(a, end) }
	f := func(a Cord, g int) int { return h(a) + g }

	queue := &CordWeightedQueue{}
	from := map[Cord]Cord{}

	push := func(a Cord, g int) { heap.Push(queue, &WeightedCord{c: a, g: g, w: f(a, g)}) }
	pop := func() (Cord, int) {
		wc := heap.Pop(queue).(*WeightedCord)
		return wc.c, wc.g
	}

	push(start, 0)

	for queue.Len() != 0 {
		if _, found_goal := from[end]; found_goal {
			break
		}

		c, g := pop()
		for _, n := range []Cord{c.Up(), c.Down(), c.Right(), c.Left()} {
			if !n.InBounds(width, height) {
				continue
			}
			if n[0] == start[0] && n[1] == start[1] {
				continue
			}
			if n[0] == end[0] && n[1] == end[1] {
				from[n] = c
				break
			}
			if grid[n[1]][n[0]] != -1 {
				continue
			}
			if _, found_cord := from[n]; found_cord {
				continue
			}

			from[n] = c
			push(n, g+1)
		}
	}

	if _, found := from[end]; !found {
		return nil
	}

	path := []*pb.Coordinate{}
	n := end
	for {
		cord := &pb.Coordinate{X: uint32(n[0]), Y: uint32(n[1])}
		path = append(path, cord)
		nn, found := from[n]
		if !found {
			break
		}
		n = nn
	}

	slices.Reverse(path)

	return path
}

func distance_pathfinder(a Cord, b Cord) int {
	return AbsI(a[0]-b[0]) + AbsI(a[1]-b[1])
}
