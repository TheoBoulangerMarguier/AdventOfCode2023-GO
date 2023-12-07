/*
--- Day 7: Camel Cards ---
Your all-expenses-paid trip turns out to be a one-way, five-minute ride in an airship.
(At least it's a cool airship!) It drops you off at the edge of a vast
desert and descends back to Island Island.

"Did you bring the parts?"

You turn around to see an Elf completely covered in white clothing, wearing goggles,
and riding a large camel.

"Did you bring the parts?" she asks again, louder this time. You aren't sure what
parts she's looking for; you're here to figure out why the sand stopped.

"The parts! For the sand, yes! Come with me; I will show you." She beckons
you onto the camel.

After riding a bit across the sands of Desert Island, you can see what look like
very large rocks covering half of the horizon. The Elf explains that the rocks
are all along the part of Desert Island that is directly above Island Island,
making it hard to even get there. Normally, they use big machines to move the rocks
 and filter the sand, but the machines have broken down because Desert Island
 recently stopped receiving the parts they need to fix the machines.

You've already assumed it'll be your job to figure out why the parts stopped when
 she asks if you can help. You agree automatically.

Because the journey will take a few days, she offers to teach you the
game of Camel Cards. Camel Cards is sort of similar to poker except
it's designed to be easier to play while riding a camel.

In Camel Cards, you get a list of hands, and your goal is to order them based
on the strength of each hand. A hand consists of five cards labeled one of
A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2.
The relative strength of each card follows this order,
where A is the highest and 2 is the lowest.

Every hand is exactly one type. From strongest to weakest, they are:

Five of a kind, where all five cards have the same label: AAAAA
Four of a kind, where four cards have the same label and one card has a different label: AA8AA
Full house, where three cards have the same label, and the remaining two cards share a different label: 23332
Three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98
Two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432
One pair, where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
High card, where all cards' labels are distinct: 23456
Hands are primarily ordered based on type; for example, every full house is stronger than any three of a kind.

If two hands have the same type, a second ordering rule takes effect.
Start by comparing the first card in each hand. If these cards are different,
the hand with the stronger first card is considered stronger.
If the first card in each hand have the same label, however,
then move on to considering the second card in each hand.
If they differ, the hand with the higher second card wins;
otherwise, continue with the third card in each hand, then the fourth, then the fifth.

So, 33332 and 2AAAA are both four of a kind hands, but 33332 is stronger because
its first card is stronger. Similarly, 77888 and 77788 are both a full house,
but 77888 is stronger because its third card is stronger
(and both hands have the same first and second card).

To play Camel Cards, you are given a list of hands and their
corresponding bid (your puzzle input). For example:

32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483

This example shows five hands; each hand is followed by its bid amount.
Each hand wins an amount equal to its bid multiplied by its rank,
where the weakest hand gets rank 1, the second-weakest hand gets rank 2,
and so on up to the strongest hand. Because there are five hands in this example,
the strongest hand will have rank 5 and its bid will be multiplied by 5.

So, the first step is to put the hands in order of strength:

32T3K is the only one pair and the other hands are all a stronger type,
so it gets rank 1.
KK677 and KTJJT are both two pair. Their first cards both have the same label,
but the second card of KK677 is stronger (K vs T), so KTJJT gets rank 2 and KK677 gets rank 3.
T55J5 and QQQJA are both three of a kind. QQQJA has a stronger first card,
so it gets rank 5 and T55J5 gets rank 4.
Now, you can determine the total winnings of this set of hands by adding up the
result of multiplying each hand's bid with its rank
(765 * 1 + 220 * 2 + 28 * 3 + 684 * 4 + 483 * 5).
So the total winnings in this example are 6440.

Find the rank of every hand in your set. What are the total winnings?


--- Part Two ---
To make things a little more interesting, the Elf introduces one additional rule.
Now, J cards are jokers - wildcards that can act like whatever
card would make the hand the strongest type possible.

To balance this, J cards are now the weakest individual cards, weaker even than 2.
 The other cards stay in the same order: A, K, Q, T, 9, 8, 7, 6, 5, 4, 3, 2, J.

J cards can pretend to be whatever card is best for the purpose of determining hand type;
for example, QJJQ2 is now considered four of a kind.
However, for the purpose of breaking ties between two hands of the same type,
J is always treated as J, not the card it's pretending to be:
JKKK2 is weaker than QQQQ2 because J is weaker than Q.

Now, the above example goes very differently:

32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483

32T3K is still the only one pair; it doesn't contain any jokers,
	so its strength doesn't increase.
KK677 is now the only two pair, making it the second-weakest hand.
T55J5, KTJJT, and QQQJA are now all four of a kind! T55J5 gets rank 3,
	QQQJA gets rank 4, and KTJJT gets rank 5.
With the new joker rule, the total winnings in this example are 5905.

Using the new joker rule, find the rank of every hand in your set.
What are the new total winnings?
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type hand struct {
	cards     string
	bid       int
	frequency map[rune]int
	score     int
}

var cardScore1 map[rune]int = map[rune]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

var cardScore2 map[rune]int = map[rune]int{
	'J': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'Q': 12,
	'K': 13,
	'A': 14,
}

const FIVE_OF_A_KIND string = "Five of a kind"
const FOUR_OF_A_KIND string = "Four of a kind"
const FULL_HOUSE string = "Full house"
const THREE_OF_A_KIND string = "Three of a kind"
const TWO_PAIR string = "Two pair"
const ONE_PAIR string = "One pair"
const HIGH_CARD string = "High card"

var comboScore map[string]int = map[string]int{
	FIVE_OF_A_KIND:  7,
	FOUR_OF_A_KIND:  6,
	FULL_HOUSE:      5,
	THREE_OF_A_KIND: 4,
	TWO_PAIR:        3,
	ONE_PAIR:        2,
	HIGH_CARD:       1,
}

func Day7() {
	d7p1()
	d7p2()
}

// walk through the cards and check for repetition,
// frequency of a card is equal to number of repetition of such card
func getCharFrequency(cards string) map[rune]int {
	var output map[rune]int = map[rune]int{}

	for _, r := range cards {
		_, ok := output[r]
		if ok {
			output[r]++
		} else {
			output[r] = 1
		}
	}

	return output
}

// calculate the combo score of a set of card based on the card frequency
func getComboScoreFromHand(frequency map[rune]int, useWildCard bool) int {
	jFreq, ok := frequency['J']
	wildCardInitialized := false

	//in case of wild card we convert it to the higest value cart type
	if useWildCard && ok {

		var rHighestFreq rune
		max := 0

		if jFreq == 5 {
			return comboScore[FIVE_OF_A_KIND]
		}

		for r, i := range frequency {
			if r == 'J' {
				continue
			}
			if i > max {
				max = i
				rHighestFreq = r
			}
		}

		frequency[rHighestFreq] += jFreq
		delete(frequency, 'J')
		wildCardInitialized = true
	}

	//once wildcard check are passed we calculate score normally
	if !useWildCard || !ok || wildCardInitialized {
		tempCombo := 0
		for _, f := range frequency {
			if f == 5 {
				return comboScore[FIVE_OF_A_KIND]
			} else if f == 4 {
				return comboScore[FOUR_OF_A_KIND]
			} else if f == 3 {
				if tempCombo == 0 {
					tempCombo = comboScore[THREE_OF_A_KIND]
				} else if tempCombo == comboScore[ONE_PAIR] {
					return comboScore[FULL_HOUSE]
				}
			} else if f == 2 {
				if tempCombo == 0 {
					tempCombo = comboScore[ONE_PAIR]
				} else if tempCombo == comboScore[THREE_OF_A_KIND] {
					return comboScore[FULL_HOUSE]
				} else if tempCombo == comboScore[ONE_PAIR] {
					return comboScore[TWO_PAIR]
				}
			}
		}
		if tempCombo == 0 {
			return comboScore[HIGH_CARD]
		} else {
			return tempCombo
		}
	}
	return -1
}

// compare 2 hands and see which one has the best score (-1 : a > b) (0 : a==B) (1 a < b)
func compareHands(a hand, b hand, useWildCard bool) int {
	//TODO compare a to be to see who is higher
	scoreTable := cardScore1

	if useWildCard {
		scoreTable = cardScore2

	}

	if a.score < b.score {
		return 1
	} else if a.score > b.score {
		return -1
	} else {
		bRunes := []rune(b.cards)
		for i, aRune := range a.cards {
			if aRune == bRunes[i] {
				continue
			}

			if scoreTable[aRune] < scoreTable[bRunes[i]] {
				return 1
			} else {
				return -1
			}
		}
		return 0
	}
}

// buble sort the array of hand based on their score
func bubbleSort(arr []hand, useWildeCard bool) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if compareHands(arr[j], arr[j+1], useWildeCard) == -1 {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func d7p1() {

	//step 1 load the file and scan it to extract data
	file, err := os.Open("./Ressources/day7_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	data := []hand{}

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")

		b, err := strconv.Atoi(line[1])

		if err != nil {
			log.Fatal(err)
		}

		//get all the necessay info from the extracted hand of card
		f := getCharFrequency(line[0])
		s := getComboScoreFromHand(f, false)

		h := hand{
			cards:     line[0],
			bid:       b,
			frequency: f,
			score:     s,
		}

		data = append(data, h)
	}

	//step 2 sort the array based on score
	bubbleSort(data, false)

	//step 3 calculate result
	sumProd := 0
	for i := 1; i <= len(data); i++ {
		sumProd += i * data[i-1].bid
	}

	fmt.Printf("Result Day7 Part1: %d\n", sumProd)

}

func d7p2() {

	//step 1 load the file and scan it to extract data
	file, err := os.Open("./Ressources/day7_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	data := []hand{}

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")

		b, err := strconv.Atoi(line[1])

		if err != nil {
			log.Fatal(err)
		}

		//get all the necessay info from the extracted hand of card
		f := getCharFrequency(line[0])
		s := getComboScoreFromHand(f, true)

		h := hand{
			cards:     line[0],
			bid:       b,
			frequency: f,
			score:     s,
		}

		data = append(data, h)
	}

	//step 2 sort the array based on score
	bubbleSort(data, true)

	//step 3 calculate result
	sumProd := 0
	for i := 1; i <= len(data); i++ {
		sumProd += i * data[i-1].bid
	}

	fmt.Printf("Result Day7 Part2: %d\n", sumProd)
}
