/*
--- Day 16: The Floor Will Be Lava ---
With the beam of light completely focused somewhere, the reindeer leads you
deeper still into the Lava Production Facility. At some point, you realize
that the steel facility walls have been replaced with cave, and the doorways
are just cave, and the floor is cave, and you're pretty sure this is actually
just a giant cave.

Finally, as you approach what must be the heart of the mountain, you see a
bright light in a cavern up ahead. There, you discover that the beam of light
you so carefully focused is emerging from the cavern wall closest to the facility
and pouring all of its energy into a contraption on the opposite side.

Upon closer inspection, the contraption appears to be a flat, two-dimensional
square grid containing empty space (.), mirrors (/ and \), and splitters (| and -).

The contraption is aligned so that most of the beam bounces around the grid,
but each tile on the grid converts some of the beam's light into heat
to melt the rock in the cavern.

You note the layout of the contraption (your puzzle input). For example:

.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....

The beam enters in the top-left corner from the left and heading to the right.
Then, its behavior depends on what it encounters as it moves:

If the beam encounters empty space (.), it continues in the same direction.
If the beam encounters a mirror (/ or \), the beam is reflected 90 degrees
depending on the angle of the mirror. For instance, a rightward-moving beam
that encounters a / mirror would continue upward in the mirror's column, while
a rightward-moving beam that encounters a \ mirror would continue downward
from the mirror's column.
If the beam encounters the pointy end of a splitter (| or -), the beam passes
through the splitter as if the splitter were empty space. For instance, a
rightward-moving beam that encounters a - splitter would continue in the same
direction.
If the beam encounters the flat side of a splitter (| or -), the beam is split
into two beams going in each of the two directions the splitter's pointy ends
are pointing. For instance, a rightward-moving beam that encounters a | splitter
would split into two beams: one that continues upward from the splitter's column
and one that continues downward from the splitter's column.
Beams do not interact with other beams; a tile can have many beams passing
through it at the same time. A tile is energized if that tile has at least one
beam pass through it, reflect in it, or split in it.

In the above example, here is how the beam of light bounces around the contraption:

>|<<<\....
|v-.\^....
.v...|->>>
.v...v^.|.
.v...v^...
.v...v^..\
.v../2\\..
<->-/vv|..
.|<<<2-|.\
.v//.|.v..

Beams are only shown on empty tiles; arrows indicate the direction of the beams.
If a tile contains beams moving in multiple directions, the number of distinct
directions is shown instead. Here is the same diagram but instead only showing
whether a tile is energized (#) or not (.):

######....
.#...#....
.#...#####
.#...##...
.#...##...
.#...##...
.#..####..
########..
.#######..
.#...#.#..

Ultimately, in this example, 46 tiles become energized.

The light isn't energizing enough tiles to produce lava;
to debug the contraption, you need to start by analyzing the current
situation. With the beam starting in the top-left heading right,
how many tiles end up being energized?

--- Part Two ---
As you try to work out what might be wrong, the reindeer tugs on your shirt and
leads you to a nearby control panel. There, a collection of buttons lets you
align the contraption so that the beam enters from any edge tile and heading
away from that edge. (You can choose either of two directions for the beam if
it starts on a corner; for instance, if the beam starts in the bottom-right corner,
it can start heading either left or upward.)

So, the beam could start on any tile in the top row (heading downward),
any tile in the bottom row (heading upward), any tile in the leftmost column
(heading right), or any tile in the rightmost column (heading left).
To produce lava, you need to find the configuration that energizes as many
tiles as possible.

In the above example, this can be achieved by starting the beam in the fourth
tile from the left in the top row:

.|<2<\....
|v-v\^....
.v.v.|->>>
.v.v.v^.|.
.v.v.v^...
.v.v.v^..\
.v.v/2\\..
<-2-/vv|..
.|<<<2-|.\
.v//.|.v..

Using this configuration, 51 tiles are energized:

.#####....
.#.#.#....
.#.#.#####
.#.#.##...
.#.#.##...
.#.#.##...
.#.#####..
########..
.#######..
.#...#.#..

Find the initial beam configuration that energizes the largest number of tiles;
how many tiles are energized in that configuration?
*/

package Day16

import (
	"bufio"
	"log"
	"os"
	"sync"
)

func Day16() [2]int {
	return [2]int{
		d16p1(),
		d16p2(),
	}
}

const (
	empty     = 0
	mirror1   = 1
	mirror2   = 2
	splitter1 = 3
	splitter2 = 4
)

type Vector2 struct {
	X int
	Y int
}

type BeamHead struct {
	direction Vector2
	position  Vector2
}

func loadData(path string) (int, int, map[Vector2]int) {
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
	gridTypes := map[Vector2]int{}
	y := 0
	x := 0
	for scanner.Scan() {
		x = 0
		for _, r := range scanner.Text() {
			p := Vector2{
				X: x,
				Y: y,
			}

			newCell := -1

			if r == '.' {
				newCell = empty
			} else if r == '/' {
				newCell = mirror1
			} else if r == '\\' {
				newCell = mirror2
			} else if r == '-' {
				newCell = splitter1
			} else if r == '|' {
				newCell = splitter2
			}
			gridTypes[p] = newCell
			x++
		}
		y++
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	return x, y, gridTypes
}

func d16p1() int {
	_, _, gridTypes := loadData("./Day16/Ressources/day16_input.txt")
	gridState := map[Vector2]bool{}
	copyGrid := map[Vector2]int{}
	for k, v := range gridTypes {
		copyGrid[k] = v
		gridState[k] = false
	}

	return SendBeam(copyGrid, gridState, Vector2{0, 0}, Vector2{1, 0})
}

func d16p2() int {
	xMax, yMax, gridTypes := loadData("./Day16/Ressources/day16_input.txt")
	results := []int{}
	// Limit the number of concurrent goroutines using a semaphore
	maxConcurrent := 10 // Set the maximum number of concurrent goroutines
	semaphore := make(chan struct{}, maxConcurrent)

	// columns from 0 to xmax (inclusive)
	var wg1 sync.WaitGroup
	wg1.Add(xMax)
	for x := 0; x < xMax; x++ {
		semaphore <- struct{}{} // Acquire semaphore
		go func(x int) {
			defer func() { <-semaphore }() // Release semaphore
			defer wg1.Done()

			gridState := map[Vector2]bool{}
			copyGrid := map[Vector2]int{}
			for k, v := range gridTypes {
				copyGrid[k] = v
				gridState[k] = false
			}
			results = append(results, SendBeam(copyGrid, gridState, Vector2{x, 0}, Vector2{0, 1}))

			gridState = map[Vector2]bool{}
			copyGrid = map[Vector2]int{}
			for k, v := range gridTypes {
				copyGrid[k] = v
				gridState[k] = false
			}
			results = append(results, SendBeam(copyGrid, gridState, Vector2{x, yMax - 1}, Vector2{0, -1}))
		}(x)
	}
	wg1.Wait()

	// rows from 0 to ymax (inclusive)
	var wg2 sync.WaitGroup
	wg2.Add(yMax)
	for y := 0; y < yMax; y++ {
		semaphore <- struct{}{} // Acquire semaphore
		go func(y int) {
			defer func() { <-semaphore }() // Release semaphore
			defer wg2.Done()

			gridState := map[Vector2]bool{}
			copyGrid := map[Vector2]int{}
			for k, v := range gridTypes {
				copyGrid[k] = v
				gridState[k] = false
			}

			results = append(results, SendBeam(copyGrid, gridState, Vector2{0, y}, Vector2{1, 0}))

			gridState = map[Vector2]bool{}
			copyGrid = map[Vector2]int{}
			for k, v := range gridTypes {
				copyGrid[k] = v
				gridState[k] = false
			}

			results = append(results, SendBeam(copyGrid, gridState, Vector2{xMax - 1, y}, Vector2{-1, 0}))
		}(y)
	}
	wg2.Wait()

	max := 0
	for _, r := range results {
		if r > max {
			max = r
		}
	}

	return max
}

func updateBeamHead(beamHead *BeamHead, dir Vector2, pos Vector2) {
	beamHead.direction = dir
	beamHead.position = pos
}

func updateCell(state *map[Vector2]bool, pos Vector2) int {
	if !(*state)[pos] {
		(*state)[pos] = true
		return 1
	}
	return 0
}

// check if a coordinate is in grid's bounds
func isInGrid(grid map[Vector2]int, pos Vector2) bool {
	_, ok := grid[pos]
	return ok
}

func SendBeam(gridTypes map[Vector2]int, gridState map[Vector2]bool, startPos Vector2, startDir Vector2) int {
	beamHeads := []BeamHead{{startDir, startPos}}
	updateCell(&gridState, startPos)

	cumul := 0
	lastCount := -1
	currentCount := 1

	for cumul != 10 {
		lastCount = currentCount
		for i, b := range beamHeads {

			nextPos := Vector2{
				X: b.position.X + b.direction.X,
				Y: b.position.Y + b.direction.Y,
			}

			if isInGrid(gridTypes, nextPos) {
				if gridTypes[nextPos] == empty {
					updateBeamHead(&beamHeads[i], beamHeads[i].direction, nextPos)
					currentCount += updateCell(&gridState, nextPos)
				} else if gridTypes[nextPos] == mirror1 {
					if b.direction.X == 1 {
						updateBeamHead(&beamHeads[i], Vector2{0, -1}, nextPos)
					} else if b.direction.X == -1 {
						updateBeamHead(&beamHeads[i], Vector2{0, 1}, nextPos)
					} else if b.direction.Y == 1 {
						updateBeamHead(&beamHeads[i], Vector2{-1, 0}, nextPos)
					} else if b.direction.Y == -1 {
						updateBeamHead(&beamHeads[i], Vector2{1, 0}, nextPos)
					}
					currentCount += updateCell(&gridState, nextPos)
				} else if gridTypes[nextPos] == mirror2 {
					if b.direction.X == 1 {
						updateBeamHead(&beamHeads[i], Vector2{0, 1}, nextPos)
					} else if b.direction.X == -1 {
						updateBeamHead(&beamHeads[i], Vector2{0, -1}, nextPos)
					} else if b.direction.Y == 1 {
						updateBeamHead(&beamHeads[i], Vector2{1, 0}, nextPos)
					} else if b.direction.Y == -1 {
						updateBeamHead(&beamHeads[i], Vector2{-1, 0}, nextPos)
					}
					currentCount += updateCell(&gridState, nextPos)
				} else if gridTypes[nextPos] == splitter1 {
					if b.direction.X != 0 {
						updateBeamHead(&beamHeads[i], beamHeads[i].direction, nextPos)
					} else {
						updateBeamHead(&beamHeads[i], Vector2{-1, 0}, nextPos)
						newBeamHead := BeamHead{Vector2{1, 0}, nextPos}
						beamHeads = append(beamHeads, newBeamHead)
					}
					currentCount += updateCell(&gridState, nextPos)
				} else if gridTypes[nextPos] == splitter2 {
					if b.direction.Y != 0 {
						updateBeamHead(&beamHeads[i], beamHeads[i].direction, nextPos)
					} else {
						updateBeamHead(&beamHeads[i], Vector2{0, -1}, nextPos)
						newBeamHead := BeamHead{Vector2{0, 1}, nextPos}
						beamHeads = append(beamHeads, newBeamHead)
					}
					currentCount += updateCell(&gridState, nextPos)
				}
			}
		}
		if lastCount == currentCount {
			cumul++
		} else {
			cumul = 0
		}
	}
	return currentCount
}
