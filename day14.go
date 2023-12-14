/*
--- Day 14: Parabolic Reflector Dish ---
You reach the place where all of the mirrors were pointing: a massive parabolic
reflector dish attached to the side of another large mountain.

The dish is made up of many small mirrors, but while the mirrors themselves are
roughly in the shape of a parabolic reflector dish, each individual mirror seems
to be pointing in slightly the wrong direction. If the dish is meant to focus
light, all it's doing right now is sending it in a vague direction.

This system must be what provides the energy for the lava! If you focus the
reflector dish, maybe you can go where it's pointing and use the light to fix
the lava production.

Upon closer inspection, the individual mirrors each appear to be connected via
an elaborate system of ropes and pulleys to a large metal platform below the dish.
The platform is covered in large rocks of various shapes.
Depending on their position, the weight of the rocks deforms the platform,
and the shape of the platform controls which ropes move and ultimately
the focus of the dish.

In short: if you move the rocks, you can focus the dish. The platform even has a
control panel on the side that lets you tilt it in one of four directions!
The rounded rocks (O) will roll when the platform is tilted,
while the cube-shaped rocks (#) will stay in place.
You note the positions of all of the empty spaces (.) and rocks (your puzzle input).
For example:

O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....

Start by tilting the lever so all of the rocks will slide north as far as they will go:

OOOO.#.O..
OO..#....#
OO..O##..O
O..#.OO...
........#.
..#....#.#
..O..#.O.O
..O.......
#....###..
#....#....

You notice that the support beams along the north side of the platform are damaged;
to ensure the platform doesn't collapse, you should calculate the total
load on the north support beams.

The amount of load caused by a single rounded rock (O) is equal to the number
of rows from the rock to the south edge of the platform,
including the row the rock is on. (Cube-shaped rocks (#) don't contribute to load.)
So, the amount of load caused by each rock in each row is as follows:

OOOO.#.O.. 10
OO..#....#  9
OO..O##..O  8
O..#.OO...  7
........#.  6
..#....#.#  5
..O..#.O.O  4
..O.......  3
#....###..  2
#....#....  1

The total load is the sum of the load caused by all of the rounded rocks.
In this example, the total load is 136.

Tilt the platform so that the rounded rocks all roll north.
Afterward, what is the total load on the north support beams?

--- Part Two ---
The parabolic reflector dish deforms, but not in a way that focuses the beam.
To do that, you'll need to move the rocks to the edges of the platform.
Fortunately, a button on the side of the control panel labeled "spin cycle"
attempts to do just that!

Each cycle tilts the platform four times so that the rounded rocks roll north,
then west, then south, then east. After each tilt, the rounded rocks roll as far
as they can before the platform tilts in the next direction. After one cycle,
the platform will have finished rolling the rounded rocks in those four
directions in that order.

Here's what happens in the example above after each of the first few cycles:

After 1 cycle:
.....#....
....#...O#
...OO##...
.OO#......
.....OOO#.
.O#...O#.#
....O#....
......OOOO
#...O###..
#..OO#....

After 2 cycles:
.....#....
....#...O#
.....##...
..O#......
.....OOO#.
.O#...O#.#
....O#...O
.......OOO
#..OO###..
#.OOO#...O

After 3 cycles:
.....#....
....#...O#
.....##...
..O#......
.....OOO#.
.O#...O#.#
....O#...O
.......OOO
#...O###.O
#.OOO#...O

This process should work if you leave it running long enough, but you're still
worried about the north support beams. To make sure they'll survive for a while,
you need to calculate the total load on the north support beams after 1000000000
cycles.

In the above example, after 1000000000 cycles, the total load on the north
support beams is 64.

Run the spin cycle for 1000000000 cycles. Afterward, what is the total load on the
north support beams?

*/

package main

import (
	"bufio"
	"log"
	"os"
)

func Day14() [2]int {
	return [2]int{
		d14p1(),
		d14p2(),
	}
}

type Rock struct {
	pos [2]int
}

type Cell struct {
	free bool
}

func loadData(path string) ([][]Cell, []Rock) {
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

	grid := [][]Cell{}
	rocks := []Rock{}
	y := 0
	for scanner.Scan() {
		x := 0
		grid = append(grid, []Cell{})

		for _, c := range scanner.Text() {

			if c == '.' {
				//new cell free
				newCell := Cell{
					free: true,
				}
				grid[y] = append(grid[y], newCell)
				x++
			} else if c == '#' {
				//new cell non free
				newCell := Cell{
					free: false,
				}
				grid[y] = append(grid[y], newCell)
				x++
			} else if c == 'O' {
				//new rock && new cell
				newCell := Cell{
					free: false,
				}
				grid[y] = append(grid[y], newCell)

				newRock := Rock{
					pos: [2]int{y, x},
				}
				rocks = append(rocks, newRock)
				x++
			}
		}
		y++
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	return grid, rocks
}

func d14p1() int {
	grid, rocks := loadData("./Ressources/day14_input.txt")
	gravity := [2]int{-1, 0}

	tilt(grid, rocks, gravity)

	sum := 0
	for i := 0; i < len(rocks); i++ {
		sum += len(grid) - rocks[i].pos[0]
	}
	return sum
}

func d14p2() int {
	grid, rocks := loadData("./Ressources/day14_input.txt")

	//big cycle that endup looping over the same value, no need to run all billion times
	//1000 is enough to get the result, 100 is not
	for i := 0; i < 1000; i++ {
		tilt(grid, rocks, [2]int{-1, 0}) //North
		tilt(grid, rocks, [2]int{0, -1}) //West
		tilt(grid, rocks, [2]int{1, 0})  //south
		tilt(grid, rocks, [2]int{0, 1})  //East
	}

	sum := 0
	for i := 0; i < len(rocks); i++ {
		sum += len(grid) - rocks[i].pos[0]
	}
	return sum
}

// check if a coordinate is in grid's bounds
func isInGrid(pos [2]int, grid [][]Cell) bool {
	if pos[0] < 0 || pos[0] >= len(grid) || pos[1] < 0 || pos[1] >= len(grid[0]) {
		return false
	}
	return true
}

// celular automata that move rocks if condition met
func tilt(grid [][]Cell, rocks []Rock, gravity [2]int) {
	hasAnyRockMoved := true

	for hasAnyRockMoved {
		hasAnyRockMoved = false

		for i := 0; i < len(rocks); i++ {
			currentPos := rocks[i].pos
			nextPos := [2]int{
				currentPos[0] + gravity[0],
				currentPos[1] + gravity[1]}

			if isInGrid(nextPos, grid) && grid[nextPos[0]][nextPos[1]].free {
				grid[nextPos[0]][nextPos[1]].free = false
				grid[currentPos[0]][currentPos[1]].free = true
				rocks[i].pos = nextPos
				hasAnyRockMoved = true
			}
		}
	}
}
