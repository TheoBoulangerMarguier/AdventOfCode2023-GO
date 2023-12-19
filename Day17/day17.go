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
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type Point struct {
	x, y int
}

type Node struct {
	weight  int
	g, h, f int
	parent  Point
	dir     Point
	steps   int
}

type Grid struct {
	size  Point
	nodes map[Point]Node
	open  []Point
	close []Point
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
	grid := Grid{Point{0, 0}, map[Point]Node{}, []Point{}, []Point{}}
	x, y := 0, 0
	for scanner.Scan() {
		x = 0
		for _, r := range scanner.Text() {
			v, err := strconv.Atoi(string(r))
			if err != nil {
				log.Fatal(err)
			} else {
				node := Node{
					v,
					0, 0, 0,
					Point{-1, -1},
					Point{0, 0},
					0,
				}
				grid.nodes[Point{x, y}] = node
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

func d17p1() int {
	//initialization
	grid := loadData("./Day17/Ressources/day17_input-mini.txt")
	start := Point{0, 0}
	end := Point{grid.size.x - 1, grid.size.y - 1}
	grid.open = append(grid.open, start)
	completed := false

	//Core logic of A*
	for !completed {
		//get the lowest F cost Node from the open list
		i, currentPos := getLowestFCostPos(grid)
		currentNode := grid.nodes[currentPos]

		//move this node from open to close list
		grid.open = append(grid.open[:i], grid.open[i+1:]...)
		grid.close = append(grid.close, currentPos)

		//early break if we find the end node
		if currentPos == end {
			completed = true
			break
		}

		//get all direct neighbours		 {left}	 {right}   {up}    {down}
		for _, offset := range []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			neighbourPos := Point{currentPos.x + offset.x, currentPos.y + offset.y}

			//check if the position is valid in the grid
			if _, inGrid := grid.nodes[neighbourPos]; !inGrid {
				continue
			}

			//check if the position is not in the grid or if we can still move in this direction
			if contain(grid.close, neighbourPos) || currentNode.steps == 3 && offset == currentNode.dir {
				continue
			}

			//avoid backtrack
			invert := Point{currentNode.dir.x * -1, currentNode.dir.y * -1}
			if offset == invert {
				continue
			}

			//calculate neighbour GHF costs
			neighbourNode := grid.nodes[neighbourPos]
			neighbourG := neighbourNode.weight + currentNode.g
			neighbourH := manhattanDistance(neighbourPos, end)
			neighbourF := neighbourG + neighbourH

			//check if not in open or path is shorter than existing
			inOpenList := contain(grid.open, neighbourPos)
			if !inOpenList || neighbourG < neighbourNode.g {

				//update neighbour node data in grid
				neighbourNode.g = neighbourG
				neighbourNode.h = neighbourH
				neighbourNode.f = neighbourF
				neighbourNode.parent = currentPos
				neighbourNode.dir = offset

				//update directional counter either increase of set to 1
				neighbourNode.steps = getUpdatedStepCount(currentNode, offset)

				grid.nodes[neighbourPos] = neighbourNode

				//register to the open list for next iteration
				if !inOpenList {
					grid.open = append(grid.open, neighbourPos)
				}
			}
		}
	}

	//get resulting path and the sum of the weights travelled
	path, sum := getPath(start, end, grid)
	printPath(grid, path)

	return sum
}

func d17p2() int {
	return 0
}

// look into the open nodes and will get the one withe the lowest F cost
// in case there is a tie, get the one that as the lowest H cost
func getLowestFCostPos(grid Grid) (int, Point) {
	empty := Point{-1, -1}
	minPos := empty
	minID := -1
	for i, nodePos := range grid.open {
		if minPos == empty {
			minPos, minID = nodePos, i
			continue
		}

		node := grid.nodes[nodePos]
		minNode := grid.nodes[minPos]

		if node.f < minNode.f {
			minPos, minID = nodePos, i
		} else if node.g == minNode.g && node.h < minNode.h {
			minPos, minID = nodePos, i
		}
	}
	return minID, minPos
}

// get the manhattan distance between 2 points of the grid
func manhattanDistance(p1, p2 Point) int {
	return int(math.Abs(float64(p1.x-p2.x)) + math.Abs(float64(p1.y-p2.y)))
}

// check a Point item is in the specified slice
func contain(slice []Point, item Point) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == item {
			return true
		}
	}
	return false
}

// backtrack the parent node from end to start to get the shortest path
func getPath(start Point, end Point, grid Grid) ([]Point, int) {
	path := []Point{}
	current := end
	sum := 0
	for current != start {
		sum += grid.nodes[current].weight
		path = append(path, current)
		current = grid.nodes[current].parent
	}
	return path, sum
}

// helper to print the final state of the grid once path is found, will show weight
func printPath(grid Grid, path []Point) {
	sum := 0
	for y := 0; y < grid.size.y; y++ {
		for x := 0; x < grid.size.x; x++ {
			if contain(path, Point{x, y}) {
				fmt.Print(grid.nodes[Point{x, y}].weight)
				sum += grid.nodes[Point{x, y}].weight
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println("path sum weights:", sum)
}

// increase the step count by 1 if the direction is teh same otherwise will return 1
func getUpdatedStepCount(currentNode Node, direction Point) int {
	if currentNode.dir == direction {
		return currentNode.steps + 1
	} else {
		return 1
	}
}
