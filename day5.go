/*
--- Day 5: If You Give A Seed A Fertilizer ---
You take the boat and find the gardener right where you were told he would be:
managing a giant "garden" that looks more to you like a farm.

"A water source? Island Island is the water source!" You point out that
Snow Island isn't receiving any water.

"Oh, we had to stop the water because we ran out of sand to filter it with!
Can't make snow with dirty water. Don't worry, I'm sure we'll get more sand soon;
 we only turned off the water a few days... weeks... oh no."
 His face sinks into a look of horrified realization.

"I've been so busy making sure everyone here has food that I completely forgot
to check why we stopped getting more sand! There's a ferry leaving soon that is
headed over in that direction - it's much faster than your boat.
Could you please go check it out?"

You barely have time to agree to this request when he brings up another.
"While you wait for the ferry, maybe you can help us with our food
production problem. The latest Island Island Almanac just arrived and
we're having trouble making sense of it."

The almanac (your puzzle input) lists all of the seeds that need to be planted.
It also lists what type of soil to use with each kind of seed,
what type of fertilizer to use with each kind of soil, what type of water
to use with each kind of fertilizer, and so on. Every type of seed, soil,
fertilizer and so on is identified with a number, but numbers are reused
by each category - that is, soil 123 and fertilizer 123 aren't necessarily
related to each other.

For example:

seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
The almanac starts by listing which seeds need to be planted:
seeds 79, 14, 55, and 13.

The rest of the almanac contains a list of maps which describe how to
convert numbers from a source category into numbers in a destination category.
That is, the section that starts with seed-to-soil map: describes how to
convert a seed number (the source) to a soil number (the destination).
This lets the gardener and his team know which soil to use with which seeds,
which water to use with which fertilizer, and so on.

Rather than list every source number and its corresponding destination
number one by one, the maps describe entire ranges of numbers that
can be converted. Each line within a map contains three numbers:
the destination range start, the source range start, and the range length.

Consider again the example seed-to-soil map:

50 98 2
52 50 48
The first line has a destination range start of 50, a source range start of 98,
and a range length of 2. This line means that the source range starts at 98 and
contains two values: 98 and 99. The destination range is the same length,
but it starts at 50, so its two values are 50 and 51. With this information,
you know that seed number 98 corresponds to soil number 50 and that
seed number 99 corresponds to soil number 51.

The second line means that the source range starts at 50 and contains 48 values:
50, 51, ..., 96, 97. This corresponds to a destination range starting at 52
and also containing 48 values: 52, 53, ..., 98, 99. So,
seed number 53 corresponds to soil number 55.

Any source numbers that aren't mapped correspond to the same destination number.
So, seed number 10 corresponds to soil number 10.

So, the entire list of seed numbers and their corresponding soil numbers
looks like this:

seed  soil
0     0
1     1
...   ...
48    48
49    49
50    52
51    53
...   ...
96    98
97    99
98    50
99    51
With this map, you can look up the soil number required for each initial seed number:

Seed number 79 corresponds to soil number 81.
Seed number 14 corresponds to soil number 14.
Seed number 55 corresponds to soil number 57.
Seed number 13 corresponds to soil number 13.
The gardener and his team want to get started as soon as possible,
so they'd like to know the closest location that needs a seed. Using these maps,
find the lowest location number that corresponds to any of the initial seeds.
To do this, you'll need to convert each seed number through other categories
until you can find its corresponding location number. In this example,
the corresponding types are:

Seed 79, soil 81, fertilizer 81, water 81, light 74, temperature 78, humidity 78, location 82.
Seed 14, soil 14, fertilizer 53, water 49, light 42, temperature 42, humidity 43, location 43.
Seed 55, soil 57, fertilizer 57, water 53, light 46, temperature 82, humidity 82, location 86.
Seed 13, soil 13, fertilizer 52, water 41, light 34, temperature 34, humidity 35, location 35.
So, the lowest location number in this example is 35.

What is the lowest location number that corresponds to any of the initial seed numbers?


--- Part Two ---
Everyone will starve if you only plant such a small number of seeds.
Re-reading the almanac, it looks like the seeds:
line actually describes ranges of seed numbers.

The values on the initial seeds: line come in pairs. Within each pair,
the first value is the start of the range and the second value is the length of
the range. So, in the first line of the example above:

seeds: 79 14 55 13
This line describes two ranges of seed numbers to be planted in the garden.
The first range starts with seed number 79 and contains 14 values:
79, 80, ..., 91, 92. The second range starts with seed number
55 and contains 13 values: 55, 56, ..., 66, 67.

Now, rather than considering four seed numbers, you need to consider a total
of 27 seed numbers.

In the above example, the lowest location number can be obtained from seed
number 82, which corresponds to soil 84, fertilizer 84, water 84, light 77,
temperature 45, humidity 46, and location 46. So, the lowest location number is 46.

Consider all of the initial seed numbers listed in the ranges on the first
line of the almanac. What is the lowest location number that corresponds
to any of the initial seed numbers?
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

// main representation of the input from the exercise (seeds)
type block struct {
	start         int
	end           int
	slicedAlready bool
}

// main representation of the input categories filters (soil, water, etc..)
type blockFilter struct {
	start   int
	end     int
	newBase int
}

func Day5() {
	d5p1()
	d5p2()
}

// extract the txt input into a slice fo seeds :[]int ;and a slice for each category
// the category slices contains a slice of filters : [][]blockFilter
func Extract() ([][]blockFilter, []int) {
	file, err := os.Open("./Ressources/day5_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	seedsList := make([]int, 0)

	categoryMap := [][]blockFilter{
		{},
		{},
		{},
		{},
		{},
		{},
		{},
	}

	lastLine := ""
	currentCategory := -1
	lineID := 1

	for scanner.Scan() {
		if lineID == 1 {
			//get seed list
			split := strings.Split(strings.Split(scanner.Text(), "seeds: ")[1], " ")
			for i := 0; i < len(split); i++ {
				seedId, err := strconv.Atoi(split[i])
				if err != nil {
					log.Fatal(err)
				}
				seedsList = append(seedsList, seedId)
			}
		} else {
			//check if we enter a new category:
			if scanner.Text() != "" {
				if lastLine == "" {
					currentCategory++
					lastLine = scanner.Text()
					continue
				}

				line := strings.Split(scanner.Text(), " ")
				if len(line) != 3 {
					log.Fatal("no 3 numbers found in line:\"" + scanner.Text() + "\"")
				}

				dst, err1 := strconv.Atoi(line[0])
				src, err2 := strconv.Atoi(line[1])
				size, err3 := strconv.Atoi(line[2])

				if err1 != nil {
					log.Fatal(err1)
				}
				if err2 != nil {
					log.Fatal(err2)
				}
				if err3 != nil {
					log.Fatal(err3)
				}

				filter := blockFilter{
					start:   src,
					end:     src + size - 1,
					newBase: dst,
				}

				categoryMap[currentCategory] = append(categoryMap[currentCategory], filter)
			}
		}

		lineID++
		lastLine = scanner.Text()
	}

	return categoryMap, seedsList
}

// main logic for the conversion used in part 1 of the exercise
// it work on the list of seed as if each element is a seed
// unlike part2 that has seedID+range
func Part1Converter(filters []blockFilter, key int) int {
	result := -1

	for filterID := 0; filterID < len(filters); filterID++ {
		if key >= filters[filterID].start && key < filters[filterID].end {
			//       destination range start +  position - source range start
			result = filters[filterID].newBase + key - filters[filterID].start
			break
		}
	}

	if result == -1 {
		return key
	} else {
		return result
	}
}

// core logic of part1 return the result in print
func d5p1() {

	category, seeds := Extract()

	min := math.MaxInt
	for i := 0; i < len(seeds); i++ {
		s := Part1Converter(category[0], seeds[i])
		s = Part1Converter(category[1], s)
		s = Part1Converter(category[2], s)
		s = Part1Converter(category[3], s)
		s = Part1Converter(category[4], s)
		s = Part1Converter(category[5], s)
		s = Part1Converter(category[6], s)

		if s < min {
			min = s
		}
	}

	fmt.Printf("Result Day5 Part1: %d\n", min)

}

// detect overlap between a range A and a range B using their start/end as coordinates
func checkIntersect(aStart int, aEnd int, bStart int, bEnd int) (int, string) {
	if aEnd < bStart || bEnd < aStart {
		if aEnd == bStart-1 {
			return 7, "A glued to B by the left but no intersect"
		} else if bEnd == aStart-1 {
			return 8, "A glued to B by the right but no intersect"
		} else {
			return 0, "no overlap and no glued data"
		}
	} else if aStart >= bStart && aEnd <= bEnd {
		return 1, "B fully overlap A"
	} else if aStart <= bStart && aEnd <= bEnd {
		return 2, "single point right"
	} else if aEnd >= bEnd && aStart >= bEnd {
		return 3, "single point left"
	} else if aStart < bStart && aEnd <= bEnd {
		return 4, "A overlap B from the left"
	} else if aStart >= bStart && aEnd > bEnd {
		return 5, "A overlap B from the right"
	} else if bStart > aStart && bEnd < aEnd {
		return 6, "A fully overlap B"
	}
	return -1, "ERROR case not handled"
}

// take 2 block and merge them so the ranges expressed are the smallest possible
// Consider {-1,-1} as a void part
func blockMerger(a block, b block) [2]block {
	output := [2]block{
		{
			start: -1,
			end:   -1,
		},
		{
			start: -1,
			end:   -1,
		},
	}

	switch code, logMsg := checkIntersect(a.start, a.end, b.start, b.end); code {
	case 0:
		output[0] = a
		output[1] = b
	case 1:
		output[0] = b
	case 2:
		output[0].start = a.start
		output[0].end = b.end
	case 3:
		output[0].start = b.start
		output[0].end = a.end
	case 4:
		output[0].start = a.start
		output[0].end = b.end
	case 5:
		output[0].start = b.start
		output[0].end = a.end
	case 6:
		output[0] = a
	case 7:
		output[0].start = a.start
		output[0].end = b.end
	case 8:
		output[0].start = b.start
		output[0].end = a.end
	case -1:
		log.Fatal(logMsg)
	}

	return output
}

// takes in 1 block and one filter and cut out the part of the block that fit the filter
// will return all bits of the original bloc after the cut, expressed by their new
// coordinates / range.
// Consider {-1,-1} as a void part
func filterCutBlock(a block, b blockFilter) []block {
	output := []block{
		{
			start: -1,
			end:   -1,
		},
		{
			start: -1,
			end:   -1,
		},
		{
			start: -1,
			end:   -1,
		},
	}

	switch code, logMsg := checkIntersect(a.start, a.end, b.start, b.end); code {
	case 0:
		output[1].start = a.start
		output[1].end = a.end
	case 1:
		output[0].start = a.start
		output[0].end = a.end
	case 2:
		output[0].start = a.end
		output[0].end = a.end
		output[1].start = a.start
		output[1].end = output[0].start - 1
	case 3:
		output[0].start = a.start
		output[0].end = a.start
		output[1].start = output[0].end - 1
		output[1].end = a.end
	case 4:
		output[0].start = b.start
		output[0].end = a.end
		output[1].start = a.start
		output[1].end = output[0].start - 1
	case 5:
		output[0].start = a.start
		output[0].end = b.end
		output[1].start = output[0].end + 1
		output[1].end = a.end
	case 6:
		output[0].start = b.start
		output[0].end = b.end
		output[1].start = a.start
		output[1].end = output[0].start - 1
		output[2].start = output[0].end + 1
		output[2].end = a.end
	case -1:
		log.Fatal(logMsg)
	}

	return output
}

// convert an input range START END to a filter map START END NEWBASE
func filterConversion(inputValue int, filterSrc int, filterDst int) int {
	return inputValue - filterSrc + filterDst
}

// takes in a slice and will compare each elements to merge and remove overlaps
// return the same slice with hopefully less elements
func checkForMerge(slice []block) []block {
	for aBlockID := 0; aBlockID < len(slice); aBlockID++ {
		for bBlockID := aBlockID + 1; bBlockID < len(slice); bBlockID++ {
			merge := blockMerger(slice[aBlockID], slice[bBlockID])
			slice[aBlockID].start = merge[0].start
			slice[aBlockID].end = merge[0].end
			if merge[1].start != -1 {
				slice[bBlockID].start = merge[1].start
				slice[bBlockID].end = merge[1].end
			} else {
				//swaping unwanted ID with last and shrink the slice size by 1
				slice[bBlockID], slice[len(slice)-1] = slice[len(slice)-1], slice[bBlockID]
				slice = slice[:len(slice)-1]
				bBlockID--
			}
		}
	}
	return slice
}

// core logic of part 2, will return the result as print
func d5p2() {
	filters, seeds := Extract()

	if len(seeds)%2 != 0 {
		log.Fatal("ERROR : SEED/RANGE broken, needs to be a pair amount of numbers")
	}

	blocks := []block{}

	//convert seed list to blocks
	for seedID := 0; seedID < len(seeds); seedID += 2 {
		newBlock := block{
			start: seeds[seedID],
			end:   seeds[seedID] + seeds[seedID+1] - 1,
		}
		blocks = append(blocks, newBlock)
	}

	//multi threading the scan
	var wg sync.WaitGroup
	multithreadLenght := len(blocks)
	wg.Add(len(blocks))
	packetOutput := []int{}

	for blockID := 0; blockID < multithreadLenght; blockID++ {
		go func(blockID int) {
			defer wg.Done()

			fmt.Println("starting blockID:", blockID, blocks[blockID])

			//creating a packet for each bloc that will contain the splits of this block
			packet := []block{
				blocks[blockID],
			}

			for layerID := 0; layerID < len(filters); layerID++ {
				for filterID := 0; filterID < len(filters[layerID]); filterID++ {
					for extractID := 0; extractID < len(packet); extractID++ {
						if packet[extractID].slicedAlready {
							continue
						}
						//in each layer
						cut := filterCutBlock(packet[extractID], filters[layerID][filterID])

						if cut[0].start != -1 {

							packet[extractID].start = filterConversion(
								cut[0].start,
								filters[layerID][filterID].start,
								filters[layerID][filterID].newBase)

							packet[extractID].end = filterConversion(
								cut[0].end,
								filters[layerID][filterID].start,
								filters[layerID][filterID].newBase)

							packet[extractID].slicedAlready = true

							if cut[1].start != -1 {
								packet = append(packet, cut[1])
							}

							if cut[2].start != -1 {
								packet = append(packet, cut[2])
							}

						}
					}
				}

				//regroup data to avoid slices getting to big layer after layer
				packet = checkForMerge(packet)
				for packetID := 0; packetID < len(packet); packetID++ {
					packet[packetID].slicedAlready = false
				}
			}

			// check for min for this final packet and push it to the output
			packetMin := math.MaxInt
			for packetID := 0; packetID < len(packet); packetID++ {
				if packet[packetID].start < packetMin {
					packetMin = packet[packetID].start
				}
			}
			packetOutput = append(packetOutput, packetMin)
			fmt.Println("ending blockID:", blockID, "min value of this block: ", packetMin)

		}(blockID)
	}

	wg.Wait()

	//check for real min after
	min := math.MaxInt
	for _, output := range packetOutput {
		if output < min {
			min = output
		}
	}
	fmt.Printf("Result Day5 Part2: %d\n", min)
}
