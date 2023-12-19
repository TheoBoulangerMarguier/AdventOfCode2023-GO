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

func Day17() [2]int {
	return [2]int{
		d17p1(),
		d17p2(),
	}
}

type Vector2 struct {
	x, y int
}

type Node struct {
	weight   int
	g, h, f  int
	parent   Vector2
	lastDir  Vector2
	dirCount int
}

type Grid struct {
	size  Vector2
	nodes map[Vector2]Node
	open  []Vector2
	close []Vector2
}

func loadData(path string) Grid {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	grid := Grid{Vector2{0, 0}, map[Vector2]Node{}, []Vector2{}, []Vector2{}}
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
					Vector2{-1, -1},
					Vector2{0, 0},
					0,
				}
				grid.nodes[Vector2{x, y}] = node
			}
			x++
		}
		y++
	}
	grid.size = Vector2{x, y}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	return grid
}

func d17p1() int {
	/* INPUT : "./Day17/Ressources/day17_input-mini.txt"

	11199
	12199
	99199
	99131
	99111

	*/
	grid := loadData("./Day17/Ressources/day17_input-mini.txt")

	start := Vector2{0, 0}
	end := Vector2{grid.size.x - 1, grid.size.y - 1}

	grid.open = append(grid.open, start)

	completed := false
	for !completed {
		//get the lowest F cost Node from the open list
		i, current := getLowestFCostPos(grid)

		//move this node from open to close list
		grid.open = append(grid.open[:i], grid.open[i+1:]...)
		grid.close = append(grid.close, current)

		//early break if we find the end node
		if current == end {
			completed = true
			break
		}

		//get all direct neighbours		 {left}	 {right}   {up}    {down}
		for _, offset := range []Vector2{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			neighbourPos := Vector2{current.x + offset.x, current.y + offset.y}

			//check if the position is valid in the grid
			_, inGrid := grid.nodes[neighbourPos]
			if !inGrid {
				continue
			}

			//check if the position is not in the grid or if we can still move in this direction
			inCloseList := contain(grid.close, neighbourPos)
			reachedMaxStep := grid.nodes[current].dirCount == 3 &&
				offset == grid.nodes[current].lastDir

			if inCloseList || reachedMaxStep {
				continue
			}

			//calculate neighbour GHF costs
			neighbourG := grid.nodes[neighbourPos].weight + grid.nodes[current].g
			neighbourH := manhattanDistance(neighbourPos, end)
			neighbourF := neighbourG + neighbourH

			//check if not in open or path is shorter than existing
			inOpenList := contain(grid.open, neighbourPos)
			if !inOpenList || neighbourG < grid.nodes[neighbourPos].g {

				//update neighbour node data in grid
				nodeUpdate := grid.nodes[neighbourPos]
				nodeUpdate.g = neighbourG
				nodeUpdate.h = neighbourH
				nodeUpdate.f = neighbourF
				nodeUpdate.parent = current
				nodeUpdate.lastDir = offset

				//update directional counter
				if grid.nodes[current].lastDir == offset {
					nodeUpdate.dirCount = grid.nodes[current].dirCount + 1
				} else {
					nodeUpdate.dirCount = 1
				}

				grid.nodes[neighbourPos] = nodeUpdate

				//register to the open list for next iteration
				if !inOpenList {
					grid.open = append(grid.open, neighbourPos)
				}
			}
		}
	}

	//build path
	path := []Vector2{}
	current := end
	for current != start {
		path = append(path, current)
		current = grid.nodes[current].parent
	}
	fmt.Println()

	//print path
	sum := 0
	for y := 0; y < grid.size.y; y++ {
		for x := 0; x < grid.size.x; x++ {
			if contain(path, Vector2{x, y}) {
				fmt.Print(grid.nodes[Vector2{x, y}].weight)
				sum += grid.nodes[Vector2{x, y}].weight
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println("path sum weights:", sum)

	/*ISSUE, output :

	.11..
	..1..
	..1..
	..131
	....1

	SUM : 10
	EXPECTED: 9

	*/

	return 0
}

func d17p2() int {
	return 0
}

func getLowestFCostPos(grid Grid) (int, Vector2) {
	empty := Vector2{-1, -1}
	minPos := Vector2{-1, -1}
	minID := -1
	for i, nodePos := range grid.open {
		if minPos == empty {
			minPos = nodePos
			minID = i
			continue
		}

		if grid.nodes[nodePos].f < grid.nodes[minPos].f {
			minPos = nodePos
			minID = i
		} else if grid.nodes[nodePos].g == grid.nodes[minPos].g &&
			grid.nodes[nodePos].h < grid.nodes[minPos].h {
			minID = i
			minPos = nodePos
		}
	}
	return minID, minPos
}

func manhattanDistance(p1, p2 Vector2) int {
	return int(math.Abs(float64(p1.x-p2.x)) + math.Abs(float64(p1.y-p2.y)))
}

func contain(slice []Vector2, item Vector2) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == item {
			return true
		}
	}
	return false
}
