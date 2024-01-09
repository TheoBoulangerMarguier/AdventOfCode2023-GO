/*
--- Day 21: Step Counter ---
You manage to catch the airship right as it's dropping someone else off on their all-expenses-paid trip to Desert Island! It even helpfully drops you off near the gardener and his massive farm.

"You got the sand flowing again! Great work! Now we just need to wait until we have enough sand to filter the water for Snow Island and we'll have snow again in no time."

While you wait, one of the Elves that works with the gardener heard how good you are at solving problems and would like your help. He needs to get his steps in for the day, and so he'd like to know which garden plots he can reach with exactly his remaining 64 steps.

He gives you an up-to-date map (your puzzle input) of his starting position (S), garden plots (.), and rocks (#). For example:

...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........

The Elf starts at the starting position (S) which also counts as a garden plot. Then, he can take one step north, south, east, or west, but only onto tiles that are garden plots. This would allow him to reach any of the tiles marked O:

...........
.....###.#.
.###.##..#.
..#.#...#..
....#O#....
.##.OS####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........
Then, he takes a second step. Since at this point he could be at either tile marked O, his second step would allow him to reach any garden plot that is one step north, south, east, or west of any tile that he could have reached after the first step:

...........
.....###.#.
.###.##..#.
..#.#O..#..
....#.#....
.##O.O####.
.##.O#...#.
.......##..
.##.#.####.
.##..##.##.
...........
After two steps, he could be at any of the tiles marked O above, including the starting position (either by going north-then-south or by going west-then-east).

A single third step leads to even more possibilities:

...........
.....###.#.
.###.##..#.
..#.#.O.#..
...O#O#....
.##.OS####.
.##O.#...#.
....O..##..
.##.#.####.
.##..##.##.
...........
He will continue like this until his steps for the day have been exhausted. After a total of 6 steps, he could reach any of the garden plots marked O:

...........
.....###.#.
.###.##.O#.
.O#O#O.O#..
O.O.#.#.O..
.##O.O####.
.##.O#O..#.
.O.O.O.##..
.##.#.####.
.##O.##.##.
...........
In this example, if the Elf's goal was to get exactly 6 more steps today, he could use them to reach any of 16 garden plots.

However, the Elf actually needs to get 64 steps today, and the map he's handed you is much larger than the example map.

Starting from the garden plot marked S on your map, how many garden plots could the Elf reach in exactly 64 steps?
*/

package Day21

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type Point struct {
	x, y int
}

type Float64Point struct {
	x, y float64
}

type Cell struct {
	pos, chunkPos, parent Point
	value                 rune
	step                  int
	visited               bool
}

func Day21() [2]int {
	return [2]int{
		d21p1(),
		d21p2(),
	}
}

func ParseInput() (map[Point]Cell, Point, Point) {
	file, err := os.Open("./Day21/Ressources/day21_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	grid := map[Point]Cell{}
	start := Point{-1, -1}
	bounds := Point{-1, -1}
	y := 0
	for scanner.Scan() {
		chars := []rune(scanner.Text())
		for x := 0; x < len(chars); x++ {
			pos := Point{x, y}

			cell := Cell{
				pos:      pos,
				chunkPos: Point{0, 0},
				parent:   Point{-1, -1},
				value:    chars[x],
				step:     0,
			}
			if chars[x] == 'S' {
				start = pos
				cell.visited = true
			}
			grid[pos] = cell
		}

		y++
		bounds.x = len(chars)
		bounds.y = y
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	file.Close()
	return grid, start, bounds
}

func d21p1() int {
	grid, start, _ := ParseInput()
	maxStep, count := 64, 0
	even := maxStep%2 == 0
	if even {
		count++
	}
	toVisit := []Point{start}
	for i := 0; i < len(toVisit); i++ {
		if grid[toVisit[i]].step > maxStep-1 {
			break
		}
		//find neighbours that are not visited and
		neighbors := []Point{
			{
				x: toVisit[i].x,
				y: toVisit[i].y - 1,
			},
			{
				x: toVisit[i].x,
				y: toVisit[i].y + 1,
			},
			{
				x: toVisit[i].x - 1,
				y: toVisit[i].y,
			},
			{
				x: toVisit[i].x + 1,
				y: toVisit[i].y,
			},
		}

		for _, nPos := range neighbors {
			cell, ok := grid[nPos]
			if ok && cell.value != '#' && !cell.visited {
				cell.step = grid[toVisit[i]].step + 1
				cell.visited = true
				grid[nPos] = cell
				toVisit = append(toVisit, nPos)
				vEven := cell.step%2 == 0
				if (even && vEven) || (!even && !vEven) {
					count++
				}
			}
		}
	}
	return count
}

func d21p2() int {
	/*
		searching for searchDelta(65,196,65+131*x) will give us the d2 when we only increase from 65 by 131*x wpaced by x = 0,1,2
		with input known (puzzle input):
			65+131*x = input
			131*x = input - 65
			x = (input-65)/131

		delta of delta when jumping by 65 with periodicity of 131
			=> d2 = (p2.y - p1.y) / (p2.x - p1.x) - (p1.y - p0.y) / (p1.x - p0.x)
			=> n = d2/2

		we will solve p2(input, ???)	where ??? is the answer to the puzzle
		the dDelta result of the input will be:
			=> n*x = d2

		let's replace the original d2 equation with the new data for the prediction:
			d2 is n*x so we can express:
				(??? - p1.y) / (input - p1.x) - (p1.y - p0.y) / (p1.x - p0.x) = (n*x)
				(??? - p1.y) / (input - p1.x) = (n*x) + (p1.y - p0.y) / (p1.x - p0.x)
				(??? - p1.y) = ((n*x) + (p1.y - p0.y) / (p1.x - p0.x))*(input - p1.x)
				??? = (((n*x) + (p1.y - p0.y) / (p1.x - p0.x))*(input - p1.x))+p1.y
	*/
	input := float64(26501365)
	d2, p0, p1, _ := searchDelta(65, 65+131, 65+131*2)
	n := d2 / 2
	x := (input - 65) / 131
	y := (((n * x) + (p1.y-p0.y)/(p1.x-p0.x)) * (input - p1.x)) + p1.y

	return int(math.Ceil(y))
}

func searchDelta(a, b, c int) (float64, Float64Point, Float64Point, Float64Point) {
	p0 := Float64Point{float64(a), float64(getCountInExpandingGrid(a))}
	p1 := Float64Point{float64(b), float64(getCountInExpandingGrid(b))}
	p2 := Float64Point{float64(c), float64(getCountInExpandingGrid(c))}

	d0 := (p1.y - p0.y) / (p1.x - p0.x)
	d1 := (p2.y - p1.y) / (p2.x - p1.x)

	d2 := d1 - d0

	fmt.Println("x values for y=", a, b, c, ":", p0.y, p1.y, p2.y, " delta:", d2)
	return d2, p0, p1, p2
}

func getCountInExpandingGrid(steps int) int {
	templateGrid, start, bounds := ParseInput()
	grid := deepMapCopy(templateGrid)

	maxStep, count := steps, 0
	even := maxStep%2 == 0
	if even {
		count++
	}
	toVisit := []Point{start}
	for i := 0; i < len(toVisit); i++ {
		if grid[toVisit[i]].step > maxStep-1 {
			break
		}
		//find neighbours that are not visited and
		neighbors := []Point{
			{
				x: toVisit[i].x,
				y: toVisit[i].y - 1,
			},
			{
				x: toVisit[i].x,
				y: toVisit[i].y + 1,
			},
			{
				x: toVisit[i].x - 1,
				y: toVisit[i].y,
			},
			{
				x: toVisit[i].x + 1,
				y: toVisit[i].y,
			},
		}

		for _, nPos := range neighbors {
			cell, ok := grid[nPos]

			if !ok {
				direction := Point{
					x: nPos.x - toVisit[i].x,
					y: nPos.y - toVisit[i].y,
				}
				expandGrid(grid, templateGrid, direction, grid[toVisit[i]].chunkPos, bounds)
				cell, ok = grid[nPos]
			}

			if ok && cell.value != '#' && !cell.visited {
				cell.step = grid[toVisit[i]].step + 1
				cell.visited = true
				grid[nPos] = cell
				toVisit = append(toVisit, nPos)
				vEven := cell.step%2 == 0
				if (even && vEven) || (!even && !vEven) {
					count++
				}
			}
		}
	}
	return count
}

func expandGrid(gridTobeExpanded, template map[Point]Cell, direction, origin, bounds Point) {
	left := Point{-1, 0}
	right := Point{1, 0}
	up := Point{0, -1}
	down := Point{0, 1}

	offsetX, offsetY := 0, 0

	if direction == left {
		offsetX = -bounds.x
	} else if direction == right {
		offsetX = bounds.x
	} else if direction == up {
		offsetY = -bounds.y
	} else if direction == down {
		offsetY = bounds.y
	} else {
		log.Fatal("[expandGrid] wrong direction:", direction)
	}

	for y := 0; y < bounds.y; y++ {
		for x := 0; x < bounds.x; x++ {
			pos := Point{x: origin.x + x + offsetX, y: origin.y + y + offsetY}
			newCell := Cell{
				pos:      pos,
				chunkPos: Point{x: origin.x + offsetX, y: origin.y + offsetY},
				value:    template[Point{x, y}].value,
			}
			gridTobeExpanded[pos] = newCell
		}
	}
}

func deepMapCopy(grid map[Point]Cell) map[Point]Cell {
	newGrid := map[Point]Cell{}
	for k, v := range grid {
		newCell := Cell{
			pos:      Point{v.pos.x, v.pos.y},
			chunkPos: Point{v.chunkPos.x, v.chunkPos.y},
			parent:   Point{v.parent.x, v.parent.y},
			value:    v.value,
			step:     v.step,
			visited:  v.visited,
		}
		newGrid[k] = newCell
	}
	return newGrid
}
