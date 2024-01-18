/*
--- Day 23: A Long Walk ---
The Elves resume water filtering operations! Clean water starts flowing over the
edge of Island Island.

They offer to help you go over the edge of Island Island, too! Just hold on tight
to one end of this impossibly long rope and they'll lower you down a safe distance
from the massive waterfall you just created.

As you finally reach Snow Island, you see that the water isn't really reaching
the ground: it's being absorbed by the air itself. It looks like you'll finally
have a little downtime while the moisture builds up to snow-producing levels.
Snow Island is pretty scenic, even without any snow; why not take a walk?

There's a map of nearby hiking trails (your puzzle input) that indicates paths (.),
forest (#), and steep slopes (^, >, v, and <).

For example:

#.#####################
#.......#########...###
#######.#########.#.###
###.....#.>.>.###.#.###
###v#####.#v#.###.#.###
###.>...#.#.#.....#...#
###v###.#.#.#########.#
###...#.#.#.......#...#
#####.#.#.#######.#.###
#.....#.#.#.......#...#
#.#####.#.#.#########v#
#.#...#...#...###...>.#
#.#.#v#######v###.###v#
#...#.>.#...>.>.#.###.#
#####v#.#.###v#.#.###.#
#.....#...#...#.#.#...#
#.#########.###.#.#.###
#...###...#...#...#.###
###.###.#.###v#####v###
#...#...#.#.>.>.#.>.###
#.###.###.#.###.#.#v###
#.....###...###...#...#
#####################.#

You're currently on the single path tile in the top row; your goal is to reach
the single path tile in the bottom row. Because of all the mist from the waterfall,
the slopes are probably quite icy; if you step onto a slope tile, your next step
must be downhill (in the direction the arrow is pointing).
To make sure you have the most scenic hike possible,
never step onto the same tile twice. What is the longest hike you can take?

In the example above, the longest hike you can take is marked with O, and your
starting position is marked S:

#S#####################
#OOOOOOO#########...###
#######O#########.#.###
###OOOOO#OOO>.###.#.###
###O#####O#O#.###.#.###
###OOOOO#O#O#.....#...#
###v###O#O#O#########.#
###...#O#O#OOOOOOO#...#
#####.#O#O#######O#.###
#.....#O#O#OOOOOOO#...#
#.#####O#O#O#########v#
#.#...#OOO#OOO###OOOOO#
#.#.#v#######O###O###O#
#...#.>.#...>OOO#O###O#
#####v#.#.###v#O#O###O#
#.....#...#...#O#O#OOO#
#.#########.###O#O#O###
#...###...#...#OOO#O###
###.###.#.###v#####O###
#...#...#.#.>.>.#.>O###
#.###.###.#.###.#.#O###
#.....###...###...#OOO#
#####################O#

This hike contains 94 steps.
(The other possible hikes you could have taken were 90, 86, 82, 82, and 74 steps long.)

Find the longest hike you can take through the hiking trails listed on your map.
How many steps long is the longest hike?

--- Part Two ---
As you reach the trailhead, you realize that the ground isn't as slippery as you expected; you'll have no problem climbing up the steep slopes.

Now, treat all slopes as if they were normal paths (.). You still want to make sure you have the most scenic hike possible, so continue to ensure that you never step onto the same tile twice. What is the longest hike you can take?

In the example above, this increases the longest hike to 154 steps:

#S#####################
#OOOOOOO#########OOO###
#######O#########O#O###
###OOOOO#.>OOO###O#O###
###O#####.#O#O###O#O###
###O>...#.#O#OOOOO#OOO#
###O###.#.#O#########O#
###OOO#.#.#OOOOOOO#OOO#
#####O#.#.#######O#O###
#OOOOO#.#.#OOOOOOO#OOO#
#O#####.#.#O#########O#
#O#OOO#...#OOO###...>O#
#O#O#O#######O###.###O#
#OOO#O>.#...>O>.#.###O#
#####O#.#.###O#.#.###O#
#OOOOO#...#OOO#.#.#OOO#
#O#########O###.#.#O###
#OOO###OOO#OOO#...#O###
###O###O#O###O#####O###
#OOO#OOO#O#OOO>.#.>O###
#O###O###O#O###.#.#O###
#OOOOO###OOO###...#OOO#
#####################O#
Find the longest hike you can take through the surprisingly dry hiking trails listed on your map. How many steps long is the longest hike?
*/

package Day23

import (
	utils "AdventOfCode/Utils"
	"bufio"
	"fmt"
	"os"
)

type Point2 struct {
	x, y int
}

type Cell struct {
	value   string
	step    int
	visited bool
	dir     Point2
	parent  Point2
}

type Grid struct {
	smallGrid, fullGrid map[Point2]Cell
	start, end, bounds  Point2
}

func Day23() [2]int {
	return [2]int{
		d23p1(),
		d23p2(),
	}
}

const (
	part1 int = 1
	part2 int = 2
)

func d23p1() int {
	return FindLongestPathlenghtP1(DfsP1(loadData("./Day23/Ressources/day23_input-mini2.txt")))
}

func d23p2() int {
	gridData := loadData("./Day23/Ressources/day23_input-mini2.txt")
	paths := BrutForceDfsP2(gridData)
	size, bestPath := FindLongestPathlenghtP2(paths)
	printgridPath(gridData, bestPath)
	return size
}

func loadData(path string) Grid {
	file, errFile := os.Open(path)
	if errFile != nil {
		panic(errFile)
	}
	defer func() {
		if errClose := file.Close(); errClose != nil {
			panic(errClose)
		}
	}()
	scanner := bufio.NewScanner(file)

	gridData := Grid{
		smallGrid: map[Point2]Cell{},
		fullGrid:  map[Point2]Cell{},
		start:     Point2{-1, -1},
		end:       Point2{-1, -1},
		bounds:    Point2{0, 0},
	}

	lines := [][]rune{}
	for scanner.Scan() {
		text := []rune(scanner.Text())
		lines = append(lines, text)

		//Create cells for each characters
		for x := 0; x < len(text); x++ {
			pos := Point2{x, gridData.bounds.y}
			cell := Cell{
				value:  string(text[x]),
				step:   0,
				dir:    Point2{0, 0},
				parent: Point2{-1, -1},
			}
			if cell.value == "#" {
				gridData.fullGrid[pos] = cell
			} else {
				gridData.smallGrid[pos] = cell
				gridData.fullGrid[pos] = cell
			}
		}
		gridData.bounds.y++

		//update the bounds of the grid
		if len(text) > gridData.bounds.x {
			gridData.bounds.x = len(text)
		}
	}

	//find start and end pos in first and last line
	foundStart, foundEnd := false, false
	for i := 0; i < gridData.bounds.x; i++ {
		if lines[0][i] == '.' {
			gridData.start = Point2{i, 0}
			foundStart = true
		}
		if lines[len(lines)-1][i] == '.' {
			gridData.end = Point2{i, len(lines) - 1}
			foundEnd = true
		}
		if foundStart && foundEnd {
			break
		}
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	return gridData
}

func DfsP1(gridData Grid) Grid {
	//init start cell
	startCell := gridData.smallGrid[gridData.start]
	startCell.step = 1
	startCell.dir = Point2{0, 1}
	gridData.smallGrid[gridData.start] = startCell

	toVisit := []Point2{gridData.start}
	for i := 0; i < len(toVisit); i++ {
		currentPos := toVisit[i]
		currentCell := gridData.smallGrid[currentPos]
		currentCell.visited = true

		for _, dir := range []Point2{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			nextPos := Point2{currentPos.x + dir.x, currentPos.y + dir.y}
			nextCell, exist := gridData.smallGrid[nextPos]

			//ignore out of bounds positions
			if !exist {
				continue
			}

			//prevent from backtracking
			if dir.x == currentCell.dir.x*-1 && dir.y == currentCell.dir.y*-1 {
				continue
			}

			//prevent from climbing slopes
			if (dir.x == -1 && nextCell.value == ">") || (dir.y == -1 && nextCell.value == "v") {
				continue
			}

			//update the neighbor cell and add to the toVisit stack
			if (nextCell.visited && nextCell.step < currentCell.step+1) || !nextCell.visited {
				nextCell.step = currentCell.step + 1
				nextCell.dir = dir
				nextCell.parent = currentPos
				gridData.smallGrid[nextPos] = nextCell
				toVisit = append(toVisit, nextPos)
			}
		}

		gridData.smallGrid[currentPos] = currentCell
		gridData.fullGrid[currentPos] = currentCell
	}
	printgrid(gridData)
	return gridData
}

func FindLongestPathlenghtP1(gridData Grid) int {
	path := []Point2{gridData.end}
	for i := 0; i < len(path); i++ {
		pos := path[i]

		if pos.x == gridData.start.x && pos.y == gridData.start.y {
			break
		}

		maxStep := 0
		maxPos := Point2{-1, -1}
		for _, dir := range []Point2{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			nextPos := Point2{pos.x + dir.x, pos.y + dir.y}
			nextCell, exist := gridData.smallGrid[nextPos]
			inPath, err := utils.SliceContains(path, nextPos)
			canRight := dir.x != 1 || (dir.x == 1 && nextCell.value != ">")
			canDown := dir.y != 1 || (dir.y == 1 && nextCell.value != "v")

			if err != nil {
				panic(err)
			}

			if exist && !inPath && canRight && canDown {
				if nextCell.step >= maxStep {
					maxStep = nextCell.step
					maxPos = nextPos
				}
			}
		}
		if maxPos.x != -1 && maxPos.y != -1 {
			path = append(path, maxPos)
		}
	}
	return len(path) - 1
}

func BrutForceDfsP2(gridData Grid) [][]Point2 {
	paths := [][]Point2{}
	finalPaths := [][]Point2{}

	paths = append(paths, []Point2{gridData.start})

	for i := 0; i < len(paths); i++ {
		fmt.Println()
		fmt.Println("path i:", len(paths[i]))
		for j := len(paths[i]) - 1; j < len(paths[i]); j++ {
			currentPos := paths[i][j]
			fmt.Println("checking ", currentPos)

			if currentPos.x == gridData.end.x && currentPos.y == gridData.end.y {
				fmt.Println("found end point")
				finalPaths = append(finalPaths, paths[i])
				break
			}

			allNextPos := []Point2{}
			for _, dir := range []Point2{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				nextPos := Point2{currentPos.x + dir.x, currentPos.y + dir.y}
				fmt.Print("checking pos:", nextPos, " in dir:", dir, "result:")
				//ignore out of bounds positions
				if _, exist := gridData.smallGrid[nextPos]; !exist {
					fmt.Print("FAIL doesn't exist \n")
					continue
				}

				//prevent from going twice on the same cell
				if inpath, err := utils.SliceContains(paths[i], nextPos); err != nil {
					panic(err)
				} else if inpath {
					fmt.Print("FAIL already in path explored, not allowed\n")
					continue
				}
				fmt.Print("PASS")
				allNextPos = append(allNextPos, nextPos)
			}

			count := len(allNextPos)
			fmt.Println("we found", count, "possible poses")
			if count == 0 {
				fmt.Println("end of the path")
				continue
			} else if count == 1 {
				fmt.Println("adding", allNextPos[0], " to the current path:", i)
				paths[i] = append(paths[i], allNextPos[0])
			} else {
				for k := 1; k < count; k++ {
					fmt.Println("creating new branch with added pos", allNextPos[k])
					paths = append(paths, append(paths[i], allNextPos[k]))
				}
				fmt.Println("and adding", allNextPos[0], " to the current path:", i)
				paths[i] = append(paths[i], allNextPos[0])
			}
		}
	}
	fmt.Println("END found", len(finalPaths), "final paths")
	return finalPaths
}

func FindLongestPathlenghtP2(paths [][]Point2) (int, []Point2) {
	max := 0
	bestPathID := -1

	if len(paths) == 0 {
		return max, []Point2{}
	}

	for i := 0; i < len(paths); i++ {
		if max < len(paths[i]) {
			max = len(paths[i])
			bestPathID = i
		}
	}
	fmt.Println(paths[bestPathID])
	return max, paths[bestPathID]
}

func printgrid(grid Grid) {
	fmt.Println()
	for y := 0; y < grid.bounds.y; y++ {
		for x := 0; x < grid.bounds.x; x++ {
			//SHOW DIR
			if grid.fullGrid[Point2{x, y}].value == "#" {
				fmt.Print(grid.fullGrid[Point2{x, y}].value)
			} else {
				if grid.fullGrid[Point2{x, y}].dir.x == 1 {
					fmt.Print("\033[32m>\033[0m")
				} else if grid.fullGrid[Point2{x, y}].dir.x == -1 {
					fmt.Print("\033[32m<\033[0m")
				} else if grid.fullGrid[Point2{x, y}].dir.y == 1 {
					fmt.Print("\033[32mv\033[0m")
				} else if grid.fullGrid[Point2{x, y}].dir.y == -1 {
					fmt.Print("\033[32m^\033[0m")
				} else {
					fmt.Print(grid.fullGrid[Point2{x, y}].value)
				}
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func printgridPath(grid Grid, path []Point2) {
	fmt.Println()
	for y := 0; y < grid.bounds.y; y++ {
		for x := 0; x < grid.bounds.x; x++ {
			//SHOW DIR
			if grid.fullGrid[Point2{x, y}].value == "#" {
				fmt.Print(grid.fullGrid[Point2{x, y}].value)
			} else {
				inpath, err := utils.SliceContains(path, Point2{x, y})
				if err != nil {
					panic(err)
				}
				if inpath {
					fmt.Print("\033[32m", grid.fullGrid[Point2{x, y}].value, "\033[0m")
				} else {
					fmt.Print(grid.fullGrid[Point2{x, y}].value)
				}
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
