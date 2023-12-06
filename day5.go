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
)

const SEED_SOIL_MAP = "seed-to-soil map:"
const SOIL_FERTILIZER_MAP = "soil-to-fertilizer map:"
const FERTILIZER_WATER_MAP = "fertilizer-to-water map:"
const WATER_LIGHT_MAP = "water-to-light map:"
const LIGHT_TEMPERATURE = "light-to-temperature map:"
const TEMPERATURE_HUMIDITY = "temperature-to-humidity map:"
const HUMIDITY_LOCATION = "humidity-to-location map:"

func Day5() {
	d5p1()
	d5p2()
}

type data struct {
	category map[string][][3]int
	seeds    []int
}

func Extract() data {
	file, err := os.Open("./Ressources/day5_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	seedsList := make([]int, 0)

	categoryMap := map[string][][3]int{
		SEED_SOIL_MAP:        [][3]int{},
		SOIL_FERTILIZER_MAP:  [][3]int{},
		FERTILIZER_WATER_MAP: [][3]int{},
		WATER_LIGHT_MAP:      [][3]int{},
		LIGHT_TEMPERATURE:    [][3]int{},
		TEMPERATURE_HUMIDITY: [][3]int{},
		HUMIDITY_LOCATION:    [][3]int{},
	}

	lastLine := ""
	currentCategory := ""
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

			if lastLine == "" {
				//changed category
				currentCategory = scanner.Text()

			} else if scanner.Text() != "" {
				//fill current category with values
				dataSplit := strings.Split(scanner.Text(), " ")

				//check if data has correctly 3 pieces
				if len(dataSplit) != 3 {
					log.Fatal("data splitting went wrong, check data in row: \n" + scanner.Text())
				}

				destinationRangeStart, err1 := strconv.Atoi(dataSplit[0])
				sourceRangeStart, err2 := strconv.Atoi(dataSplit[1])
				rangeLength, err3 := strconv.Atoi(dataSplit[2])

				//error if any of the data can't be converted to int
				if err1 != nil || err2 != nil || err3 != nil {
					log.Fatal("ERROR destinationRangeStart, or sourceRangeStart or rangeLength couldn't parse to int")
				}

				//populate the category over the range provided
				data := [3]int{destinationRangeStart, sourceRangeStart, rangeLength}
				categoryMap[currentCategory] = append(categoryMap[currentCategory], data)
			}
		}
		lineID++
		lastLine = scanner.Text()
	}

	data := data{
		category: categoryMap,
		seeds:    seedsList,
	}

	return data
}

func Converter(slice [][3]int, key int) int {
	result := -1

	for i := 0; i < len(slice); i++ {
		if key >= slice[i][1] && key < slice[i][1]+slice[i][2] {
			//       destination range start +  position - source range start
			result = slice[i][0] + key - slice[i][1]
			break
		}
	}

	if result == -1 {
		return key
	} else {
		return result
	}
}

func d5p1() {

	data := Extract()

	min := math.MaxInt
	for i := 0; i < len(data.seeds); i++ {
		s := Converter(data.category[SEED_SOIL_MAP], data.seeds[i])
		s = Converter(data.category[SOIL_FERTILIZER_MAP], s)
		s = Converter(data.category[FERTILIZER_WATER_MAP], s)
		s = Converter(data.category[WATER_LIGHT_MAP], s)
		s = Converter(data.category[LIGHT_TEMPERATURE], s)
		s = Converter(data.category[TEMPERATURE_HUMIDITY], s)
		s = Converter(data.category[HUMIDITY_LOCATION], s)

		if s < min {
			min = s
		}
	}

	fmt.Printf("Result Day5 Part1: %d\n", min)

}

/* input: a [2]int{position, range} and b [2]int{position, range}
 * output: [][2]int{{intersect pos, range},{non-intersect1 pos, range}, {non-intersect2 pos, range}}
 * check for: overlap bewteen a and b and output the different portions
 *** output[0] is intersecting part and range
 *** output[1] and  output[2] are non-intersecting part and range
 *** if there are no intersection or no non-intersection the representaion will be {-1,-1}
 */
func rangeCutter(a pair, b pair) []pair {
	output := []pair{}

	intersectStart := -1
	intersectEnd := -1
	nonIntersectingStart1 := -1
	nonIntersectingStart2 := -1
	nonIntersectingEnd1 := -1
	nonIntersectingEnd2 := -1

	if a.end < b.start || b.end < a.start {
		//fmt.Println("no overlap")
		nonIntersectingStart1 = a.start
		nonIntersectingEnd1 = a.end
	} else if a.start >= b.start && a.end <= b.end {
		//fmt.Println("full overlap")
		intersectStart = a.start
		intersectEnd = a.end
	} else if a.start <= b.start && a.end <= b.end {
		//fmt.Println("single point right")
		intersectStart = a.end
		intersectEnd = a.end
		nonIntersectingStart1 = a.start
		nonIntersectingEnd1 = intersectStart - 1
	} else if a.end >= b.end && a.start >= b.end {
		//fmt.Println("single point left")
		intersectStart = a.start
		intersectEnd = a.start
		nonIntersectingStart1 = intersectEnd - 1
		nonIntersectingEnd1 = a.end
	} else if a.start < b.start && a.end <= b.end {
		//fmt.Println("A overlap B from the left")
		intersectStart = b.start
		intersectEnd = a.end
		nonIntersectingStart1 = a.start
		nonIntersectingEnd1 = intersectStart - 1
	} else if a.start >= b.start && a.end > b.end {
		//fmt.Println("A overlap B from the right")
		intersectStart = a.start
		intersectEnd = b.end
		nonIntersectingStart1 = intersectEnd + 1
		nonIntersectingEnd1 = a.end
	} else if b.start > a.start && b.end < a.end {
		//fmt.Println("B cut A in 3")
		intersectStart = b.start
		intersectEnd = b.end
		nonIntersectingStart1 = a.start
		nonIntersectingEnd1 = intersectStart - 1
		nonIntersectingStart2 = intersectEnd + 1
		nonIntersectingEnd2 = a.end
	}

	if intersectStart == -1 {
		output = append(output, pair{
			start: -1,
			end:   -1,
			size:  -1,
		})
	} else {
		output = append(output, pair{
			start: intersectStart,
			end:   intersectEnd,
			size:  intersectEnd - intersectStart + 1,
		})
	}

	if nonIntersectingStart1 == -1 {
		output = append(output, pair{
			start: -1,
			end:   -1,
			size:  -1,
		})
	} else {
		output = append(output, pair{
			start: nonIntersectingStart1,
			end:   nonIntersectingEnd1,
			size:  nonIntersectingEnd1 - nonIntersectingStart1 + 1,
		})
	}

	if nonIntersectingStart2 == -1 {
		output = append(output, pair{
			start: -1,
			end:   -1,
			size:  -1,
		})
	} else {
		output = append(output, pair{
			start: nonIntersectingStart2,
			end:   nonIntersectingEnd2,
			size:  nonIntersectingEnd2 - nonIntersectingStart2 + 1,
		})
	}
	//fmt.Println(output)
	return output
}

type pair struct {
	start   int
	end     int
	size    int
	checked bool
}

func d5p2() {
	data := Extract()

	if len(data.seeds)%2 != 0 {
		log.Fatal("ERROR : SEED/RANGE broken, needs to be a pair amount of numbers")
	}
	min := math.MaxInt

	//todo

	fmt.Printf("Result Day5 Part2: %d\n", min)

}
