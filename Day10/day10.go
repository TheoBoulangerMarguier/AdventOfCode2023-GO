/*
--- Day 10: Pipe Maze ---
You use the hang glider to ride the hot air from Desert Island all the way up to
the floating metal island. This island is surprisingly cold and there definitely
 aren't any thermals to glide on, so you leave your hang glider behind.

You wander around for a while, but you don't find any people or animals.
However, you do occasionally find signposts labeled "Hot Springs" pointing in a
seemingly consistent direction; maybe you can find someone at the hot springs
and ask them where the desert-machine parts are made.

The landscape here is alien; even the flowers and trees are made of metal.
As you stop to admire some metal grass, you notice something metallic scurry
away in your peripheral vision and jump into a big pipe! It didn't look like any
animal you've ever seen; if you want a better look, you'll need to get ahead of it.

Scanning the area, you discover that the entire field you're standing on is densely
packed with pipes; it was hard to tell at first because they're the same metallic
silver color as the "ground". You make a quick sketch of all of the surface
pipes you can see (your puzzle input).

The pipes are arranged in a two-dimensional grid of tiles:

| is a vertical pipe connecting north and south.
- is a horizontal pipe connecting east and west.
L is a 90-degree bend connecting north and east.
J is a 90-degree bend connecting north and west.
7 is a 90-degree bend connecting south and west.
F is a 90-degree bend connecting south and east.
. is ground; there is no pipe in this tile.
S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.
Based on the acoustics of the animal's scurrying, you're confident the pipe that contains the animal is one large, continuous loop.

For example, here is a square loop of pipe:

.....
.F-7.
.|.|.
.L-J.
.....

If the animal had entered this loop in the northwest corner,
the sketch would instead look like this:

.....
.S-7.
.|.|.
.L-J.
.....

In the above diagram, the S tile is still a 90-degree F bend:
you can tell because of how the adjacent pipes connect to it.

Unfortunately, there are also many pipes that aren't connected to the loop!
This sketch shows the same loop as above:

-L|F7
7S-7|
L|7||
-L-J|
L|-JF

In the above diagram, you can still figure out which pipes form the main loop:
they're the ones connected to S, pipes those pipes connect to,
pipes those pipes connect to, and so on. Every pipe in the main loop connects to
its two neighbors (including S, which will have exactly two pipes connecting to it,
and which is assumed to connect back to those two pipes).

Here is a sketch that contains a slightly more complex main loop:

..F7.
.FJ|.
SJ.L7
|F--J
LJ...
Here's the same example sketch with the extra, non-main-loop pipe tiles also shown:

7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ
If you want to get out ahead of the animal, you should find the tile in the loop
that is farthest from the starting position.
Because the animal is in the pipe, it doesn't make sense to measure this by direct distance.
Instead, you need to find the tile that would take the longest number of steps along the loop
to reach from the starting point - regardless of which way around the loop the animal went.

In the first example with the square loop:

.....
.S-7.
.|.|.
.L-J.
.....
You can count the distance each tile in the loop is from the starting point like this:

.....
.012.
.1.3.
.234.
.....
In this example, the farthest point from the start is 4 steps away.

Here's the more complex loop again:

..F7.
.FJ|.
SJ.L7
|F--J
LJ...
Here are the distances for each tile on that loop:

..45.
.236.
01.78
14567
23...
Find the single giant loop starting at S. How many steps along the loop does it
take to get from the starting position to the point farthest from the starting position?

--- Part Two ---
You quickly reach the farthest point of the loop, but the animal never emerges. Maybe its nest is within the area enclosed by the loop?

To determine whether it's even worth taking the time to search for such a nest, you should calculate how many tiles are contained within the loop. For example:

...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........
The above loop encloses merely four tiles - the two pairs of . in the southwest and southeast (marked I below). The middle . tiles (marked O below) are not in the loop. Here is the same loop again with those regions marked:

...........
.S-------7.
.|F-----7|.
.||OOOOO||.
.||OOOOO||.
.|L-7OF-J|.
.|II|O|II|.
.L--JOL--J.
.....O.....
In fact, there doesn't even need to be a full tile path to the outside for tiles to count as outside the loop - squeezing between pipes is also allowed! Here, I is still within the loop and O is still outside the loop:

..........
.S------7.
.|F----7|.
.||OOOO||.
.||OOOO||.
.|L-7F-J|.
.|II||II|.
.L--JL--J.
..........
In both of the above examples, 4 tiles are enclosed by the loop.

Here's a larger example:

.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...
The above sketch has many random bits of ground, some of which are in the loop (I) and some of which are outside it (O):

OF----7F7F7F7F-7OOOO
O|F--7||||||||FJOOOO
O||OFJ||||||||L7OOOO
FJL7L7LJLJ||LJIL-7OO
L--JOL7IIILJS7F-7L7O
OOOOF-JIIF7FJ|L7L7L7
OOOOL7IF7||L7|IL7L7|
OOOOO|FJLJ|FJ|F7|OLJ
OOOOFJL-7O||O||||OOO
OOOOL---JOLJOLJLJOOO
In this larger example, 8 tiles are enclosed by the loop.

Any tile that isn't part of the main loop can count as being enclosed by the loop. Here's another example with many bits of junk pipe lying around that aren't connected to the main loop at all:

FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L
Here are just the tiles that are enclosed by the loop marked with I:

FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJIF7FJ-
L---JF-JLJIIIIFJLJJ7
|F|F-JF---7IIIL7L|7|
|FFJF7L7F-JF7IIL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L
In this last example, 10 tiles are enclosed by the loop.

Figure out whether you have time to search for the nest by calculating the area within the loop. How many tiles are enclosed by the loop?
*/

package Day10

import (
	"bufio"
	"log"
	"os"
)

type cell struct {
	pos      [2]int
	canUp    bool
	canDown  bool
	canLeft  bool
	canRight bool
	dist     int
	char     rune
	visited  bool
	enclosed bool
}

var asciiTable = map[rune]cell{
	'|': {
		canUp:   true,
		canDown: true,
	},
	'-': {
		canLeft:  true,
		canRight: true,
	},
	'L': {
		canUp:    true,
		canRight: true,
	},
	'J': {
		canLeft: true,
		canUp:   true,
	},
	'7': {
		canLeft: true,
		canDown: true,
	},
	'F': {
		canDown:  true,
		canRight: true,
	},
	'.': {},
	'S': {},
}

func Day10() [2]int {
	return [2]int{
		d10p1(),
		d10p2(),
	}
}

/*
scan the maze and return:
- grid map[[2]int]*cell
- startPos [2]int
- maxDist int
- maxX
- maxY
*/
func ScanMazeForMainLoop() (map[[2]int]*cell, [2]int, int, int, int) {
	file, err := os.Open("./Day10/Ressources/day10_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	grid := map[[2]int]*cell{}
	startPos := [2]int{-1, -1}
	maxX := 0
	maxY := 0
	y := 0

	for scanner.Scan() {
		row := []rune(scanner.Text())
		for x := 0; x < len(row); x++ {
			if x > maxX {
				maxX = x
			}
			if y > maxY {
				maxY = y
			}

			//convert input char to cell struct
			newCell := cell{
				pos:      [2]int{x, y},
				canUp:    asciiTable[row[x]].canUp,
				canDown:  asciiTable[row[x]].canDown,
				canLeft:  asciiTable[row[x]].canLeft,
				canRight: asciiTable[row[x]].canRight,
				char:     row[x],
			}

			if newCell.char == 'S' {
				startPos = newCell.pos
			}
			grid[newCell.pos] = &newCell
		}
		y++
	}

	//checking left neighbour's right connection to see if we can go left from start
	if v, ok := grid[[2]int{startPos[0] - 1, startPos[1]}]; ok {
		if v.canRight {
			grid[startPos].canLeft = true
		}
	}

	//checking right neighbour's left connection to see if we can go right from start

	if v, ok := grid[[2]int{startPos[0] + 1, startPos[1]}]; ok {
		if v.canLeft {
			grid[startPos].canRight = true
		}
	}

	//checking up neighbour's down connection to see if we can go up from start
	if v, ok := grid[[2]int{startPos[0], startPos[1] - 1}]; ok {
		if v.canDown {
			grid[startPos].canUp = true
		}
	}

	//checking down neighbour's up connection to see if we can go down from start
	if v, ok := grid[[2]int{startPos[0], startPos[1] + 1}]; ok {
		if v.canUp {
			grid[startPos].canDown = true
		}
	}

	maxDist := 0
	currentPos := [][2]int{startPos}

	for i := 0; i < len(currentPos); i++ {
		grid[currentPos[i]].visited = true
		hasMoved := false

		//moving down
		if grid[currentPos[i]].canDown {
			newPos := [2]int{currentPos[i][0], currentPos[i][1] + 1}
			if !grid[newPos].visited {
				grid[newPos].dist = grid[currentPos[i]].dist + 1
				currentPos = append(currentPos, newPos)
				if grid[newPos].dist > maxDist {
					maxDist = grid[newPos].dist
				}
				hasMoved = true
			}
		}

		//moving up
		if grid[currentPos[i]].canUp {
			newPos := [2]int{currentPos[i][0], currentPos[i][1] - 1}
			if !grid[newPos].visited {
				grid[newPos].dist = grid[currentPos[i]].dist + 1
				currentPos = append(currentPos, newPos)
				if grid[newPos].dist > maxDist {
					maxDist = grid[newPos].dist
				}
				hasMoved = true
			}
		}

		//moving left
		if grid[currentPos[i]].canLeft {
			newPos := [2]int{currentPos[i][0] - 1, currentPos[i][1]}
			if !grid[newPos].visited {
				grid[newPos].dist = grid[currentPos[i]].dist + 1
				currentPos = append(currentPos, newPos)
				if grid[newPos].dist > maxDist {
					maxDist = grid[newPos].dist
				}
				hasMoved = true

			}
		}

		//moving right
		if grid[currentPos[i]].canRight {
			newPos := [2]int{currentPos[i][0] + 1, currentPos[i][1]}
			if !grid[newPos].visited {
				grid[newPos].dist = grid[currentPos[i]].dist + 1
				currentPos = append(currentPos, newPos)
				if grid[newPos].dist > maxDist {
					maxDist = grid[newPos].dist
				}
				hasMoved = true
			}
		}

		//remove positions that are not able to progress (next to validated or dead ends)
		if !hasMoved {
			currentPos[i], currentPos[len(currentPos)-1] = currentPos[len(currentPos)-1], currentPos[i]
			currentPos = currentPos[:len(currentPos)-1]
			i--
		}
	}

	return grid, startPos, maxDist, maxX, maxY
}

func d10p1() int {
	_, _, maxDist, _, _ := ScanMazeForMainLoop()
	return maxDist
}

func d10p2() int {
	//we first do the extact same thing as part 1 as we need to map the loop
	grid, _, _, maxX, maxY := ScanMazeForMainLoop()
	enclosedCount := 0

	//check all the "non-visited" (loop parts) cells
	//look to the left of them and count the number of parts going up
	//if odd then assume it's enclosed
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if grid[[2]int{x, y}].visited {
				continue
			}
			leftCounter := 0
			for x2 := 0; x2 < x; x2++ {
				if grid[[2]int{x2, y}].visited && grid[[2]int{x2, y}].canUp {
					leftCounter++
				}
			}
			if leftCounter%2 != 0 && leftCounter != 0 {
				grid[[2]int{x, y}].enclosed = true
				enclosedCount++
			}
		}
	}

	return enclosedCount
}
