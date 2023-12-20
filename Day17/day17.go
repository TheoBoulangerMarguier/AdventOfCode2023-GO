/*
--- Day 17: Clumsy Crucible ---
The lava starts flowing rapidly once the Lava Production Facility is operational.
As you leave, the reindeer offers you a parachute,
allowing you to quickly reach Gear Island.

As you descend, your bird's-eye view of Gear Island reveals why you had trouble
finding anyone on your way up: half of Gear Island is empty, but the half below
you is a giant factory city!

You land near the gradually-filling pool of lava at the base of your new lavafall.
Lavaducts will eventually carry the lava throughout the city,
but to make use of it immediately, Elves are loading it into large crucibles on wheels.

The crucibles are top-heavy and pushed by hand. Unfortunately, the crucibles
become very difficult to steer at high speeds, and so it can be hard to go in a
straight line for very long.

To get Desert Island the machine parts it needs as soon as possible,
you'll need to find the best way to get the crucible from the lava pool
to the machine parts factory. To do this, you need to minimize
heat loss while choosing a route that doesn't require the crucible
to go in a straight line for too long.

Fortunately, the Elves here have a map (your puzzle input) that uses traffic
patterns, ambient temperature, and hundreds of other parameters to calculate
exactly how much heat loss can be expected for a crucible entering any particular
city block.

For example:

2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533

Each city block is marked by a single digit that represents the amount of heat
loss if the crucible enters that block. The starting point, the lava pool,
is the top-left city block; the destination, the machine parts factory,
is the bottom-right city block. (Because you already start in the top-left
block, you don't incur that block's heat loss unless you leave that
block and then return to it.)

Because it is difficult to keep the top-heavy crucible going in a
straight line for very long, it can move at most three blocks in a
single direction before it must turn 90 degrees left or right.
The crucible also can't reverse direction; after entering each city block,
it may only turn left, continue straight, or turn right.

One way to minimize heat loss is this path:

2>>34^>>>1323
32v>>>35v5623
32552456v>>54
3446585845v52
4546657867v>6
14385987984v4
44578769877v6
36378779796v>
465496798688v
456467998645v
12246868655<v
25465488877v5
43226746555v>

This path never moves more than three consecutive blocks in the same direction
and incurs a heat loss of only 102.

Directing the crucible from the lava pool to the machine parts factory,
but not moving more than three consecutive blocks in the same direction,
what is the least heat loss it can incur?
*/

package Day17

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
)

type Point struct {
	x, y int
}

type Node struct {
	pos, dir Point
	steps    int
}

type Grid struct {
	size  Point
	costs map[Point]int
}

// called by main do display the result of both parts
func Day17() [2]int {
	return [2]int{
		d17p1(),
		d17p2(),
	}
}

// get data from input file, Create a grid and fill it with Nodes
func loadData(path string) Grid {
	//open the file and check for error
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	//on closing the file check for error
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	//walk through the input and construct the default nodes
	scanner := bufio.NewScanner(file)
	grid := Grid{Point{0, 0}, map[Point]int{}}
	x, y := 0, 0
	for scanner.Scan() {
		x = 0
		for _, r := range scanner.Text() {
			v, err := strconv.Atoi(string(r))
			if err != nil {
				log.Fatal(err)
			} else {
				grid.costs[Point{x, y}] = v
			}
			x++
		}
		y++
	}
	grid.size = Point{x, y}

	//check if scanner encountered error during the scan
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	return grid
}

// part 1, find best path with constrain of max 3 steps
func d17p1() int {
	grid := loadData("./Day17/Ressources/day17_input.txt")
	start := Node{Point{0, 0}, Point{0, 0}, 0}
	goal := Node{Point{grid.size.x - 1, grid.size.y - 1}, Point{0, 0}, 0}
	path, _ := AStar(start, goal, grid, 1, 3)
	sum := 0
	for _, p := range path {
		sum += grid.costs[p.pos]
	}
	return sum - grid.costs[start.pos]
}

// part 2 find best path with steps between 4 and 10
func d17p2() int {
	grid := loadData("./Day17/Ressources/day17_input.txt")
	start := Node{Point{0, 0}, Point{0, 0}, 0}
	goal := Node{Point{grid.size.x - 1, grid.size.y - 1}, Point{0, 0}, 0}
	path, _ := AStar(start, goal, grid, 4, 10)
	sum := 0
	for _, p := range path {
		sum += grid.costs[p.pos]
	}
	return sum - grid.costs[start.pos]
}

func AStar(start, goal Node, grid Grid, minStep int, maxStep int) ([]Node, bool) {

	//intitialization
	openSet := []Node{start}
	cameFrom := map[Node]Node{}
	gScore, fScore := map[Node]int{}, map[Node]int{}
	gScore[start] = 0
	fScore[start] = manhattanDistance(start.pos, goal.pos)

	//check all the open node until none are left
	for len(openSet) != 0 {

		//select you current node based on lwoest F score
		i, current := lowestFScore(openSet, gScore)
		if pointEqual(current.pos, goal.pos) {
			return reconstructPath(cameFrom, current), true
		}
		openSet = append(openSet[:i], openSet[i+1:]...)

		//check all neighbours of the current node
		for _, offset := range []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			neighborV := Node{Point{current.pos.x + offset.x, current.pos.y + offset.y}, offset, 0}
			invertOffset := Point{offset.x * -1, offset.y * -1}

			//confirm that this neighbor is within constrains
			validPos := isInGrid(neighborV.pos, grid)
			aboveMaxStep := current.dir == offset && current.steps == maxStep
			belowMinStep := current.dir != offset && current.steps < minStep && !pointEqual(current.pos, start.pos)

			if !validPos || aboveMaxStep || belowMinStep {
				continue
			}

			//update the step count taken based on direction from current to neighbor
			if current.dir == offset {
				neighborV.steps = current.steps + 1
			} else if current.dir == invertOffset {
				continue
			} else {
				neighborV.steps = 1
			}

			//check if neighbor is a good match to be added to open list
			tentativeGScore := gScore[current] + grid.costs[neighborV.pos]
			_, initialized := gScore[neighborV]
			if !initialized {
				gScore[neighborV] = math.MaxInt
			}
			if tentativeGScore < gScore[neighborV] {
				cameFrom[neighborV] = current
				gScore[neighborV] = tentativeGScore
				fScore[neighborV] = tentativeGScore + manhattanDistance(neighborV.pos, goal.pos)
				if !contain(openSet, neighborV) {
					openSet = append(openSet, neighborV)
				}
			}
		}
	}

	return []Node{}, false
}

func reconstructPath(cameFrom map[Node]Node, current Node) []Node {
	totalPath := []Node{current}

	ended := false
	for !ended {
		v, ok := cameFrom[current]
		if ok {
			current = v
			totalPath = append([]Node{current}, totalPath...)
		} else {
			ended = true
		}
	}

	return totalPath
}

func lowestFScore(slice []Node, gScore map[Node]int) (int, Node) {
	min := math.MaxInt
	minID := -1
	var pos Node

	for i, p := range slice {
		s, ok := gScore[p]
		if ok {
			if s < min {
				pos = p
				min = s
				minID = i
			}
		}
	}

	return minID, pos
}

func pointEqual(a Point, b Point) bool {
	return (a.x == b.x && a.y == b.y)
}

// get the manhattan distance between 2 points of the grid
func manhattanDistance(p1, p2 Point) int {
	return int(math.Abs(float64(p1.x-p2.x)) + math.Abs(float64(p1.y-p2.y)))
}

// check a Point item is in the specified slice
func contain(slice []Node, item Node) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i].dir.x == item.dir.x && slice[i].dir.y == item.dir.y &&
			slice[i].pos.x == item.pos.x && slice[i].pos.y == item.pos.y &&
			slice[i].steps == item.steps {
			return true
		}
	}
	return false
}

func isInGrid(pos Point, grid Grid) bool {
	if pos.x < 0 || pos.y < 0 || pos.x > grid.size.x-1 || pos.y > grid.size.y-1 {

		return false
	}
	return true
}
