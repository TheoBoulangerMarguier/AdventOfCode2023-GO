/**/

package Day15

import (
	"bufio"
	"log"
	"os"
)

func Day15() [2]int {
	return [2]int{
		d15p1(),
		d15p2(),
	}
}

func d15p1() int {
	strs := loadData("./Day15/Ressources/day15_input.txt")
	sum := 0
	for _, str := range strs {
		sum += hash(str)
	}
	return sum
}

func d15p2() int {
	return 0
}

func loadData(path string) []string {
	output := []string{}

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

	currentBloc := []rune{}
	for scanner.Scan() {
		for _, r := range scanner.Text() {
			if r == ',' {
				output = append(output, string(currentBloc))
				currentBloc = []rune{}
			} else {
				currentBloc = append(currentBloc, r)
			}
		}
	}
	output = append(output, string(currentBloc))

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	return output
}

func hash(str string) int {
	currentValue := 0
	for _, c := range str {
		currentValue += int(c)
		currentValue *= 17
		currentValue %= 256
	}
	return currentValue
}
